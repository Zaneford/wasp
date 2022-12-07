// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the schema definition file instead

//go:build wasm
// +build wasm

package main

import (
	"github.com/iotaledger/wasp/packages/wasmvm/wasmvmhost/go/wasmvmhost"
	"github.com/iotaledger/wasp/documentation/tutorial-examples/go/solotutorialimpl"
)

func main() {
}

func init() {
	wasmvmhost.ConnectWasmHost()
}

//export on_call
func onCall(index int32) {
	solotutorialimpl.OnDispatch(index)
}

//export on_load
func onLoad() {
	solotutorialimpl.OnDispatch(-1)
}