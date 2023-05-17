package statemanager

import (
	"context"
	"fmt"
	"time"

	"github.com/samber/lo"

	"github.com/iotaledger/hive.go/logger"
	consGR "github.com/iotaledger/wasp/packages/chain/cons/cons_gr"
	"github.com/iotaledger/wasp/packages/chain/statemanager/sm_gpa"
	"github.com/iotaledger/wasp/packages/chain/statemanager/sm_gpa/sm_gpa_utils"
	"github.com/iotaledger/wasp/packages/chain/statemanager/sm_gpa/sm_inputs"
	"github.com/iotaledger/wasp/packages/chain/statemanager/sm_utils"
	"github.com/iotaledger/wasp/packages/cryptolib"
	"github.com/iotaledger/wasp/packages/gpa"
	"github.com/iotaledger/wasp/packages/isc"
	"github.com/iotaledger/wasp/packages/metrics"
	"github.com/iotaledger/wasp/packages/peering"
	"github.com/iotaledger/wasp/packages/shutdown"
	"github.com/iotaledger/wasp/packages/state"
	"github.com/iotaledger/wasp/packages/util"
	"github.com/iotaledger/wasp/packages/util/pipe"
)

type StateMgr interface {
	consGR.StateMgr
	// The StateMgr has to find a common ancestor for the prevAO and nextAO, then return
	// the state for Next ao and reject blocks in range (commonAO, prevAO]. The StateMgr
	// can determine relative positions of the corresponding blocks based on their state
	// indexes.
	ChainFetchStateDiff(
		ctx context.Context,
		prevAO, nextAO *isc.AliasOutputWithID,
	) <-chan *sm_inputs.ChainFetchStateDiffResults
	// Invoked by the chain when a set of server (access⁻¹) nodes has changed.
	// These nodes should be used to perform block replication.
	ChainNodesUpdated(serverNodes, accessNodes, committeeNodes []*cryptolib.PublicKey)
	// This is called to save a prelim block, received from other nodes.
	// That should happen on the access nodes to receive the active state faster.
	// This function should save the block (in the WAL) synchronously.
	PreliminaryBlock(block state.Block) error
}

type reqChainNodesUpdated struct {
	serverNodes    []*cryptolib.PublicKey
	accessNodes    []*cryptolib.PublicKey
	committeeNodes []*cryptolib.PublicKey
}

func (r *reqChainNodesUpdated) String() string {
	short := func(pkList []*cryptolib.PublicKey) string {
		return lo.Reduce(pkList, func(acc string, item *cryptolib.PublicKey, _ int) string {
			return acc + " " + gpa.NodeIDFromPublicKey(item).ShortString()
		}, "")
	}
	return fmt.Sprintf("{reqChainNodesUpdated, serverNodes=%s, accessNodes=%s, committeeNodes=%s",
		short(r.serverNodes),
		short(r.accessNodes),
		short(r.committeeNodes),
	)
}

type reqPreliminaryBlock struct {
	block state.Block
	reply chan error
}

func (r *reqPreliminaryBlock) String() string {
	return fmt.Sprintf("{reqPreliminaryBlock, block.L1Commitment=%v", r.block.L1Commitment())
}

func (r *reqPreliminaryBlock) Respond(err error) {
	r.reply <- err
}

type stateManager struct {
	log                  *logger.Logger
	chainID              isc.ChainID
	stateManagerGPA      gpa.GPA
	nodeRandomiser       sm_utils.NodeRandomiser
	nodeIDToPubKey       map[gpa.NodeID]*cryptolib.PublicKey
	inputPipe            pipe.Pipe[gpa.Input]
	messagePipe          pipe.Pipe[*peering.PeerMessageIn]
	nodePubKeysPipe      pipe.Pipe[*reqChainNodesUpdated]
	preliminaryBlockPipe pipe.Pipe[*reqPreliminaryBlock]
	wal                  sm_gpa_utils.BlockWAL
	net                  peering.NetworkProvider
	netPeeringID         peering.PeeringID
	timers               sm_gpa.StateManagerTimers
	ctx                  context.Context
	cleanupFun           func()
	shutdownCoordinator  *shutdown.Coordinator
}

var (
	_ StateMgr        = &stateManager{}
	_ consGR.StateMgr = &stateManager{}
)

const (
	constMsgTypeStm           byte = iota
	constDefaultTimerTickTime      = 1 * time.Second
	constStatusTimerTime           = 10 * time.Second
)

func New(
	ctx context.Context,
	chainID isc.ChainID,
	me *cryptolib.PublicKey,
	peerPubKeys []*cryptolib.PublicKey,
	net peering.NetworkProvider,
	wal sm_gpa_utils.BlockWAL,
	store state.Store,
	shutdownCoordinator *shutdown.Coordinator,
	metrics metrics.IChainStateManagerMetrics,
	pipeMetrics metrics.IChainPipeMetrics,
	log *logger.Logger,
	timersOpt ...sm_gpa.StateManagerTimers,
) (StateMgr, error) {
	nr := sm_utils.NewNodeRandomiserNoInit(gpa.NodeIDFromPublicKey(me), log)
	var timers sm_gpa.StateManagerTimers
	if len(timersOpt) > 0 {
		timers = timersOpt[0]
	} else {
		timers = sm_gpa.NewStateManagerTimers()
	}

	stateManagerGPA, err := sm_gpa.New(chainID, nr, wal, store, metrics, log, timers)
	if err != nil {
		log.Errorf("failed to create state manager GPA: %w", err)
		return nil, err
	}
	result := &stateManager{
		log:                  log,
		chainID:              chainID,
		stateManagerGPA:      stateManagerGPA,
		nodeRandomiser:       nr,
		inputPipe:            pipe.NewInfinitePipe[gpa.Input](),
		messagePipe:          pipe.NewInfinitePipe[*peering.PeerMessageIn](),
		nodePubKeysPipe:      pipe.NewInfinitePipe[*reqChainNodesUpdated](),
		preliminaryBlockPipe: pipe.NewInfinitePipe[*reqPreliminaryBlock](),
		wal:                  wal,
		net:                  net,
		netPeeringID:         peering.HashPeeringIDFromBytes(chainID.Bytes(), []byte("StateManager")), // ChainID × StateManager
		timers:               timers,
		ctx:                  ctx,
		shutdownCoordinator:  shutdownCoordinator,
	}

	pipeMetrics.TrackPipeLen("sm-inputPipe", result.inputPipe.Len)
	pipeMetrics.TrackPipeLen("sm-messagePipe", result.messagePipe.Len)
	pipeMetrics.TrackPipeLen("sm-nodePubKeysPipe", result.nodePubKeysPipe.Len)
	pipeMetrics.TrackPipeLen("sm-preliminaryBlockPipe", result.preliminaryBlockPipe.Len)

	result.handleNodePublicKeys(&reqChainNodesUpdated{
		serverNodes:    peerPubKeys,
		accessNodes:    []*cryptolib.PublicKey{},
		committeeNodes: []*cryptolib.PublicKey{},
	})

	unhook := result.net.Attach(&result.netPeeringID, peering.PeerMessageReceiverStateManager, func(recv *peering.PeerMessageIn) {
		if recv.MsgType != constMsgTypeStm {
			result.log.Warnf("Unexpected message, type=%v", recv.MsgType)
			return
		}
		result.messagePipe.In() <- recv
	})

	result.cleanupFun = func() {
		// result.inputPipe.Close() // TODO: Uncomment it.
		// result.messagePipe.Close() // TODO: Uncomment it.
		util.ExecuteIfNotNil(unhook)
	}

	go result.run()
	return result, nil
}

// -------------------------------------
// Implementations for chain package
// -------------------------------------

func (smT *stateManager) ChainFetchStateDiff(ctx context.Context, prevAO, nextAO *isc.AliasOutputWithID) <-chan *sm_inputs.ChainFetchStateDiffResults {
	input, resultCh := sm_inputs.NewChainFetchStateDiff(ctx, prevAO, nextAO)
	smT.addInput(input)
	return resultCh
}

func (smT *stateManager) ChainNodesUpdated(serverNodes, accessNodes, committeeNodes []*cryptolib.PublicKey) {
	smT.nodePubKeysPipe.In() <- &reqChainNodesUpdated{
		serverNodes:    serverNodes,
		accessNodes:    accessNodes,
		committeeNodes: committeeNodes,
	}
}

func (smT *stateManager) PreliminaryBlock(block state.Block) error {
	reply := make(chan error, 1)
	smT.preliminaryBlockPipe.In() <- &reqPreliminaryBlock{
		block: block,
		reply: reply,
	}
	return <-reply
}

// -------------------------------------
// Implementations of consGR.StateMgr
// -------------------------------------

// ConsensusStateProposal asks State manager to ensure that all the blocks for aliasOutput are available.
// `nil` is sent via the returned channel upon successful retrieval of every block for aliasOutput.
func (smT *stateManager) ConsensusStateProposal(ctx context.Context, aliasOutput *isc.AliasOutputWithID) <-chan interface{} {
	input, resultCh := sm_inputs.NewConsensusStateProposal(ctx, aliasOutput)
	smT.addInput(input)
	return resultCh
}

// ConsensusDecidedState asks State manager to return a virtual state with stateCommitment as its state commitment
func (smT *stateManager) ConsensusDecidedState(ctx context.Context, aliasOutput *isc.AliasOutputWithID) <-chan state.State {
	input, resultCh := sm_inputs.NewConsensusDecidedState(ctx, aliasOutput)
	smT.addInput(input)
	return resultCh
}

func (smT *stateManager) ConsensusProducedBlock(ctx context.Context, stateDraft state.StateDraft) <-chan state.Block {
	input, resultCh := sm_inputs.NewConsensusBlockProduced(ctx, stateDraft)
	smT.addInput(input)
	return resultCh
}

// -------------------------------------
// Internal functions
// -------------------------------------

func (smT *stateManager) addInput(input gpa.Input) {
	smT.inputPipe.In() <- input
}

func (smT *stateManager) run() { //nolint:gocyclo
	defer smT.cleanupFun()
	inputPipeCh := smT.inputPipe.Out()
	messagePipeCh := smT.messagePipe.Out()
	nodePubKeysPipeCh := smT.nodePubKeysPipe.Out()
	preliminaryBlockPipeCh := smT.preliminaryBlockPipe.Out()
	timerTickCh := smT.timers.TimeProvider.After(smT.timers.StateManagerTimerTickPeriod)
	statusTimerCh := smT.timers.TimeProvider.After(constStatusTimerTime)
	for {
		if smT.ctx.Err() != nil {
			if smT.shutdownCoordinator == nil {
				return
			}
			// TODO what should the statemgr wait for?
			smT.shutdownCoordinator.WaitNestedWithLogging(1 * time.Second)
			smT.log.Debugf("Stopping state manager, because context was closed")
			smT.shutdownCoordinator.Done()
			return
		}
		select {
		case input, ok := <-inputPipeCh:
			if ok {
				smT.handleInput(input)
			} else {
				inputPipeCh = nil
			}
		case msg, ok := <-messagePipeCh:
			if ok {
				smT.handleMessage(msg)
			} else {
				messagePipeCh = nil
			}
		case msg, ok := <-nodePubKeysPipeCh:
			if ok {
				smT.handleNodePublicKeys(msg)
			} else {
				nodePubKeysPipeCh = nil
			}
		case msg, ok := <-preliminaryBlockPipeCh:
			if ok {
				smT.handlePreliminaryBlock(msg)
			} else {
				preliminaryBlockPipeCh = nil
			}
		case now, ok := <-timerTickCh:
			if ok {
				smT.handleTimerTick(now)
				timerTickCh = smT.timers.TimeProvider.After(smT.timers.StateManagerTimerTickPeriod)
			} else {
				timerTickCh = nil
			}
		case <-statusTimerCh:
			statusTimerCh = smT.timers.TimeProvider.After(constStatusTimerTime)
			smT.log.Debugf("State manager loop iteration; there are %v inputs, %v messages, %v public key changes waiting to be handled",
				smT.inputPipe.Len(), smT.messagePipe.Len(), smT.nodePubKeysPipe.Len())
		case <-smT.ctx.Done():
			continue
		}
	}
}

func (smT *stateManager) handleInput(input gpa.Input) {
	outMsgs := smT.stateManagerGPA.Input(input)
	smT.sendMessages(outMsgs)
}

func (smT *stateManager) handleMessage(peerMsg *peering.PeerMessageIn) {
	msg, err := smT.stateManagerGPA.UnmarshalMessage(peerMsg.MsgData)
	if err != nil {
		smT.log.Warnf("Parsing message failed: %v", err)
		return
	}
	msg.SetSender(gpa.NodeIDFromPublicKey(peerMsg.SenderPubKey))
	outMsgs := smT.stateManagerGPA.Message(msg)
	smT.sendMessages(outMsgs)
}

func (smT *stateManager) handleNodePublicKeys(req *reqChainNodesUpdated) {
	smT.log.Debugf("handleNodePublicKeys: %v", req)
	smT.nodeIDToPubKey = map[gpa.NodeID]*cryptolib.PublicKey{}
	peerNodeIDs := []gpa.NodeID{}
	for _, pubKey := range req.serverNodes {
		nodeID := gpa.NodeIDFromPublicKey(pubKey)
		if _, ok := smT.nodeIDToPubKey[nodeID]; !ok {
			smT.nodeIDToPubKey[nodeID] = pubKey
			peerNodeIDs = append(peerNodeIDs, nodeID)
		}
	}
	for _, pubKey := range req.accessNodes {
		nodeID := gpa.NodeIDFromPublicKey(pubKey)
		if _, ok := smT.nodeIDToPubKey[nodeID]; !ok {
			smT.nodeIDToPubKey[nodeID] = pubKey
			// Don't use access nodes for queries.
		}
	}
	for _, pubKey := range req.committeeNodes {
		nodeID := gpa.NodeIDFromPublicKey(pubKey)
		if _, ok := smT.nodeIDToPubKey[nodeID]; !ok {
			smT.nodeIDToPubKey[nodeID] = pubKey
			peerNodeIDs = append(peerNodeIDs, nodeID)
		}
	}

	smT.log.Infof("Updating list of nodeIDs: [%v]",
		lo.Reduce(peerNodeIDs, func(acc string, item gpa.NodeID, _ int) string {
			return acc + " " + item.ShortString()
		}, ""),
	)
	smT.nodeRandomiser.UpdateNodeIDs(peerNodeIDs)
}

func (smT *stateManager) handlePreliminaryBlock(msg *reqPreliminaryBlock) {
	if !smT.wal.Contains(msg.block.Hash()) {
		if err := smT.wal.Write(msg.block); err != nil {
			smT.log.Warnf("Preliminary block %v cannot be saved to the WAL: %v", msg.block.L1Commitment(), err)
			msg.Respond(err)
			return
		}
		smT.log.Warnf("Preliminary block %v saved to the WAL.", msg.block.L1Commitment())
		msg.Respond(nil)
		return
	}
	smT.log.Warnf("Preliminary block %v already exist in the WAL.", msg.block.L1Commitment())
	msg.Respond(nil)
}

func (smT *stateManager) handleTimerTick(now time.Time) {
	smT.handleInput(sm_inputs.NewStateManagerTimerTick(now))
}

func (smT *stateManager) sendMessages(outMsgs gpa.OutMessages) {
	if outMsgs == nil {
		return
	}
	outMsgs.MustIterate(func(msg gpa.Message) {
		msgData, err := msg.MarshalBinary()
		if err != nil {
			smT.log.Warnf("Failed to marshal message for sending: %v", err)
			return
		}
		pm := &peering.PeerMessageData{
			PeeringID:   smT.netPeeringID,
			MsgReceiver: peering.PeerMessageReceiverStateManager,
			MsgType:     constMsgTypeStm,
			MsgData:     msgData,
		}
		recipientPubKey, ok := smT.nodeIDToPubKey[msg.Recipient()]
		if !ok {
			smT.log.Debugf("Dropping outgoing message, because NodeID=%s it is not in the NodeList.", msg.Recipient().ShortString())
			return
		}
		smT.net.SendMsgByPubKey(recipientPubKey, pm)
	})
}
