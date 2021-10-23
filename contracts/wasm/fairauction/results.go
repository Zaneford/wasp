// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package fairauction

import "github.com/iotaledger/wasp/packages/vm/wasmlib"

type ImmutableGetInfoResults struct {
	id int32
}

func (s ImmutableGetInfoResults) Bidders() wasmlib.ScImmutableInt32 {
	return wasmlib.NewScImmutableInt32(s.id, idxMap[IdxResultBidders])
}

func (s ImmutableGetInfoResults) Color() wasmlib.ScImmutableColor {
	return wasmlib.NewScImmutableColor(s.id, idxMap[IdxResultColor])
}

func (s ImmutableGetInfoResults) Creator() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, idxMap[IdxResultCreator])
}

func (s ImmutableGetInfoResults) Deposit() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxResultDeposit])
}

func (s ImmutableGetInfoResults) Description() wasmlib.ScImmutableString {
	return wasmlib.NewScImmutableString(s.id, idxMap[IdxResultDescription])
}

func (s ImmutableGetInfoResults) Duration() wasmlib.ScImmutableInt32 {
	return wasmlib.NewScImmutableInt32(s.id, idxMap[IdxResultDuration])
}

func (s ImmutableGetInfoResults) HighestBid() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxResultHighestBid])
}

func (s ImmutableGetInfoResults) HighestBidder() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, idxMap[IdxResultHighestBidder])
}

func (s ImmutableGetInfoResults) MinimumBid() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxResultMinimumBid])
}

func (s ImmutableGetInfoResults) NumTokens() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxResultNumTokens])
}

func (s ImmutableGetInfoResults) OwnerMargin() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxResultOwnerMargin])
}

func (s ImmutableGetInfoResults) WhenStarted() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxResultWhenStarted])
}

type MutableGetInfoResults struct {
	id int32
}

func (s MutableGetInfoResults) Bidders() wasmlib.ScMutableInt32 {
	return wasmlib.NewScMutableInt32(s.id, idxMap[IdxResultBidders])
}

func (s MutableGetInfoResults) Color() wasmlib.ScMutableColor {
	return wasmlib.NewScMutableColor(s.id, idxMap[IdxResultColor])
}

func (s MutableGetInfoResults) Creator() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, idxMap[IdxResultCreator])
}

func (s MutableGetInfoResults) Deposit() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxResultDeposit])
}

func (s MutableGetInfoResults) Description() wasmlib.ScMutableString {
	return wasmlib.NewScMutableString(s.id, idxMap[IdxResultDescription])
}

func (s MutableGetInfoResults) Duration() wasmlib.ScMutableInt32 {
	return wasmlib.NewScMutableInt32(s.id, idxMap[IdxResultDuration])
}

func (s MutableGetInfoResults) HighestBid() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxResultHighestBid])
}

func (s MutableGetInfoResults) HighestBidder() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, idxMap[IdxResultHighestBidder])
}

func (s MutableGetInfoResults) MinimumBid() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxResultMinimumBid])
}

func (s MutableGetInfoResults) NumTokens() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxResultNumTokens])
}

func (s MutableGetInfoResults) OwnerMargin() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxResultOwnerMargin])
}

func (s MutableGetInfoResults) WhenStarted() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxResultWhenStarted])
}