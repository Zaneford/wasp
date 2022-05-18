// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package coreaccounts

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type ImmutableFoundryCreateNewParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableFoundryCreateNewParams) TokenScheme() wasmtypes.ScImmutableBytes {
	return wasmtypes.NewScImmutableBytes(s.proxy.Root(ParamTokenScheme))
}

type MutableFoundryCreateNewParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableFoundryCreateNewParams) TokenScheme() wasmtypes.ScMutableBytes {
	return wasmtypes.NewScMutableBytes(s.proxy.Root(ParamTokenScheme))
}

type ImmutableFoundryDestroyParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableFoundryDestroyParams) FoundrySN() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamFoundrySN))
}

type MutableFoundryDestroyParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableFoundryDestroyParams) FoundrySN() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamFoundrySN))
}

type ImmutableFoundryModifySupplyParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableFoundryModifySupplyParams) DestroyTokens() wasmtypes.ScImmutableBool {
	return wasmtypes.NewScImmutableBool(s.proxy.Root(ParamDestroyTokens))
}

func (s ImmutableFoundryModifySupplyParams) FoundrySN() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamFoundrySN))
}

func (s ImmutableFoundryModifySupplyParams) SupplyDeltaAbs() wasmtypes.ScImmutableBigInt {
	return wasmtypes.NewScImmutableBigInt(s.proxy.Root(ParamSupplyDeltaAbs))
}

type MutableFoundryModifySupplyParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableFoundryModifySupplyParams) DestroyTokens() wasmtypes.ScMutableBool {
	return wasmtypes.NewScMutableBool(s.proxy.Root(ParamDestroyTokens))
}

func (s MutableFoundryModifySupplyParams) FoundrySN() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamFoundrySN))
}

func (s MutableFoundryModifySupplyParams) SupplyDeltaAbs() wasmtypes.ScMutableBigInt {
	return wasmtypes.NewScMutableBigInt(s.proxy.Root(ParamSupplyDeltaAbs))
}

type ImmutableHarvestParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableHarvestParams) ForceMinimumIotas() wasmtypes.ScImmutableBigInt {
	return wasmtypes.NewScImmutableBigInt(s.proxy.Root(ParamForceMinimumIotas))
}

type MutableHarvestParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableHarvestParams) ForceMinimumIotas() wasmtypes.ScMutableBigInt {
	return wasmtypes.NewScMutableBigInt(s.proxy.Root(ParamForceMinimumIotas))
}

type ImmutableTransferAllowanceToParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableTransferAllowanceToParams) AgentID() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamAgentID))
}

func (s ImmutableTransferAllowanceToParams) ForceOpenAccount() wasmtypes.ScImmutableBool {
	return wasmtypes.NewScImmutableBool(s.proxy.Root(ParamForceOpenAccount))
}

type MutableTransferAllowanceToParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableTransferAllowanceToParams) AgentID() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamAgentID))
}

func (s MutableTransferAllowanceToParams) ForceOpenAccount() wasmtypes.ScMutableBool {
	return wasmtypes.NewScMutableBool(s.proxy.Root(ParamForceOpenAccount))
}

type ImmutableAccountNFTsParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableAccountNFTsParams) AgentID() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamAgentID))
}

type MutableAccountNFTsParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableAccountNFTsParams) AgentID() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamAgentID))
}

type ImmutableBalanceParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableBalanceParams) AgentID() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamAgentID))
}

type MutableBalanceParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableBalanceParams) AgentID() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamAgentID))
}

type ImmutableFoundryOutputParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableFoundryOutputParams) FoundrySN() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamFoundrySN))
}

type MutableFoundryOutputParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableFoundryOutputParams) FoundrySN() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamFoundrySN))
}

type ImmutableGetAccountNonceParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableGetAccountNonceParams) AgentID() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamAgentID))
}

type MutableGetAccountNonceParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableGetAccountNonceParams) AgentID() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamAgentID))
}

type ImmutableNftDataParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableNftDataParams) NftID() wasmtypes.ScImmutableBytes {
	return wasmtypes.NewScImmutableBytes(s.proxy.Root(ParamNftID))
}

type MutableNftDataParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableNftDataParams) NftID() wasmtypes.ScMutableBytes {
	return wasmtypes.NewScMutableBytes(s.proxy.Root(ParamNftID))
}
