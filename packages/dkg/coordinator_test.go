package dkg_test

// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

import (
	"fmt"
	"testing"
	"time"

	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address"
	"github.com/iotaledger/wasp/packages/dkg"
	"github.com/iotaledger/wasp/packages/testutil"
	"github.com/iotaledger/wasp/plugins/peering"
	"github.com/stretchr/testify/require"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/group/edwards25519"
	"go.dedis.ch/kyber/v3/pairing"
	"go.dedis.ch/kyber/v3/util/key"
)

func TestEd25519(t *testing.T) {
	t.SkipNow()
	//
	// Create a fake network and keys for the tests.
	var timeout = 10 * time.Second
	var treshold uint32 = 2
	var peerCount = 3
	var peerLocs []string = make([]string, peerCount)
	var peerPubs []kyber.Point = make([]kyber.Point, len(peerLocs))
	var peerSecs []kyber.Scalar = make([]kyber.Scalar, len(peerLocs))
	var suite = edwards25519.NewBlakeSHA256Ed25519()
	for i := range peerLocs {
		peerLocs[i] = fmt.Sprintf("P%06d", i)
		peerSecs[i] = suite.Scalar().Pick(suite.RandomStream())
		peerPubs[i] = suite.Point().Mul(peerSecs[i], nil)
	}
	var peeringNetwork *testutil.PeeringNetwork = testutil.NewPeeringNetwork(peerLocs, peerPubs, peerSecs, 10000)
	var networkProviders []peering.NetworkProvider = peeringNetwork.NetworkProviders()
	//
	// Initialize the DKG subsystem in each node.
	var dkgNodes []dkg.CoordNodeProvider = make([]dkg.CoordNodeProvider, len(peerLocs))
	for i := range peerLocs {
		registry := testutil.NewDkgRegistryProvider(suite)
		dkgNodes[i] = dkg.InitNode(peerSecs[i], peerPubs[i], suite, networkProviders[i], registry)
	}
	//
	// Initiate the key generation from some client node.
	var coordKey = suite.Scalar().Pick(suite.RandomStream())
	var coordPub = suite.Point().Mul(coordKey, nil)
	var coordNodeProvider dkg.CoordNodeProvider = testutil.NewDkgCoordNodeProvider(
		dkgNodes,
		timeout, // Single call timeout.
	)
	sharedAddr, sharedPub, err := dkg.GenerateDistributedKey(
		coordKey, coordPub,
		peerLocs, peerPubs,
		treshold, address.VersionED25519,
		timeout, suite, coordNodeProvider,
	)
	require.Nil(t, err)
	require.NotNil(t, sharedAddr)
	require.NotNil(t, sharedPub)
}

func TestBn256(t *testing.T) {
	var timeout = 10 * time.Second
	var treshold uint32 = 2
	var peerCount = 3
	var peerLocs []string = make([]string, peerCount)
	var peerPubs []kyber.Point = make([]kyber.Point, len(peerLocs))
	var peerSecs []kyber.Scalar = make([]kyber.Scalar, len(peerLocs))
	var suite = pairing.NewSuiteBn256() // NOTE: That's from the Pairing Adapter.
	for i := range peerLocs {
		peerPair := key.NewKeyPair(suite)
		peerLocs[i] = fmt.Sprintf("P%06d", i)
		peerSecs[i] = peerPair.Private
		peerPubs[i] = peerPair.Public
	}
	var peeringNetwork *testutil.PeeringNetwork = testutil.NewPeeringNetwork(peerLocs, peerPubs, peerSecs, 10000)
	var networkProviders []peering.NetworkProvider = peeringNetwork.NetworkProviders()
	//
	// Initialize the DKG subsystem in each node.
	var dkgNodes []dkg.CoordNodeProvider = make([]dkg.CoordNodeProvider, len(peerLocs))
	for i := range peerLocs {
		registry := testutil.NewDkgRegistryProvider(suite)
		dkgNodes[i] = dkg.InitNode(peerSecs[i], peerPubs[i], suite, networkProviders[i], registry)
	}
	//
	// Initiate the key generation from some client node.
	var coordPair = key.NewKeyPair(suite)
	var coordNodeProvider dkg.CoordNodeProvider = testutil.NewDkgCoordNodeProvider(
		dkgNodes,
		timeout, // Single call timeout.
	)
	sharedAddr, sharedPub, err := dkg.GenerateDistributedKey(
		coordPair.Private, coordPair.Public,
		peerLocs, peerPubs,
		treshold, address.VersionBLS,
		timeout, suite, coordNodeProvider,
	)
	require.Nil(t, err)
	require.NotNil(t, sharedAddr)
	require.NotNil(t, sharedPub)
}
