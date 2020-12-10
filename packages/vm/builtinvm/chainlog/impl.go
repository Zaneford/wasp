package log

import (
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/kv/datatypes"
	"github.com/iotaledger/wasp/packages/kv/dict"
	"github.com/iotaledger/wasp/packages/vm/vmtypes"
)

func initialize(ctx vmtypes.Sandbox) (dict.Dict, error) {
	ctx.Eventf("logsc.initialize.begin")
	ctx.Eventf("logsc.initialize.success hname = %s", Interface.Hname().String())
	return nil, nil
}

func storeLog(ctx vmtypes.Sandbox) (dict.Dict, error) {
	ctx.Eventf("logsc.storeLog.begin")
	logData, err := ctx.Params().Get(ParamLog)
	if err != nil {
		ctx.Log().Panicf("%v", err)
	}
	state := ctx.State()

	log := datatypes.NewMustTimestampedLog(state, VarLogName)

	log.Append(ctx.GetTimestamp(), logData)

	ctx.Eventf("---------------------------")

	return nil, nil
}

func getLogInfo(ctx vmtypes.SandboxView) (dict.Dict, error) {

	state := ctx.State()
	log := datatypes.NewMustTimestampedLog(state, VarLogName)
	ret := dict.New()
	ret.Set(VarLogName, codec.EncodeInt64(int64(log.Len())))

	return ret, nil
}

func getLasts(ctx vmtypes.SandboxView) (dict.Dict, error) {

	state := ctx.State()
	l, ok, err := codec.DecodeInt64(ctx.Params().MustGet(ParamLog))
	if err != nil {
		return nil, err
	}
	if !ok {
		l = 0
	}
	log, err := datatypes.NewTimestampedLog(state, VarLogName)

	if err != nil || log.Len() < uint32(l) {
		return nil, err
	}

	tts, _ := log.TakeTimeSlice(log.Earliest(), log.Latest())
	_, last := tts.FromToIndices()
	total := tts.NumPoints()
	data, erraw := log.LoadRecordsRaw(total-uint32(l), last, false)
	//fmt.Println("RAW DATA: ", data)

	if erraw != nil {
		return nil, err
	}

	ret := dict.New()

	a, err := datatypes.NewArray(ret, VarLogName)
	if err != nil {
		return nil, err
	}
	for _, s := range data {
		a.Push(s)
	}

	return ret, nil
}

// Gets logs between timestamp interval and last N number of records
//
// Parameters:
//  - ParamFromTs From interval
//  - ParamToTs To Interval
//  - ParamLastsRecords Amount of records that you want to return
func getLogsBetweenTs(ctx vmtypes.SandboxView) (dict.Dict, error) {

	state := ctx.State()
	fromTs, ok, err := codec.DecodeInt64(ctx.Params().MustGet(ParamFromTs))
	if err != nil {
		return nil, err
	}
	if !ok {
		fromTs = 0
	}
	toTs, ok, err := codec.DecodeInt64(ctx.Params().MustGet(ParamToTs))

	if err != nil {
		return nil, err
	}
	if !ok {
		toTs = ctx.GetTimestamp()
	}

	l, ok, err := codec.DecodeInt64(ctx.Params().MustGet(ParamLastsRecords))

	if err != nil {
		return nil, err
	}
	if !ok {
		l = 0 // 0 means all
	}

	log, err := datatypes.NewTimestampedLog(state, VarLogName)

	if err != nil {
		return nil, err
	}

	tts, err := log.TakeTimeSlice(fromTs, toTs) // returns nil if empty
	if err != nil {
		return nil, err
	}
	if tts.IsEmpty() {
		// empty time slice
		return nil, nil
	}
	first, last := tts.FromToIndices()
	from := first
	if l != 0 && tts.NumPoints() > uint32(l) {
		from = last - uint32(l)
	}
	data, err := log.LoadRecordsRaw(from, last, false)

	if err != nil {
		return nil, err
	}

	ret := dict.New()
	a, err := datatypes.NewArray(ret, VarLogName)
	if err != nil {
		return nil, err
	}
	for _, s := range data {
		a.Push(s)
	}

	return ret, nil
}
