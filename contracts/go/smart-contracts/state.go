// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package htlc

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type ImmutablehtlcState struct {
	proxy wasmtypes.Proxy
}

func (s ImmutablehtlcState) InitTime() wasmtypes.ScImmutableInt64 {
	return wasmtypes.NewScImmutableInt64(s.proxy.Root(StateInitTime))
}

func (s ImmutablehtlcState) Owner() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(StateOwner))
}

func (s ImmutablehtlcState) Receivder() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(StateReceivder))
}

func (s ImmutablehtlcState) Secret() wasmtypes.ScImmutableHash {
	return wasmtypes.NewScImmutableHash(s.proxy.Root(StateSecret))
}

func (s ImmutablehtlcState) Time() wasmtypes.ScImmutableInt64 {
	return wasmtypes.NewScImmutableInt64(s.proxy.Root(StateTime))
}

func (s ImmutablehtlcState) Value() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.proxy.Root(StateValue))
}

type MutablehtlcState struct {
	proxy wasmtypes.Proxy
}

func (s MutablehtlcState) AsImmutable() ImmutablehtlcState {
	return ImmutablehtlcState(s)
}

func (s MutablehtlcState) InitTime() wasmtypes.ScMutableInt64 {
	return wasmtypes.NewScMutableInt64(s.proxy.Root(StateInitTime))
}

func (s MutablehtlcState) Owner() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(StateOwner))
}

func (s MutablehtlcState) Receivder() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(StateReceivder))
}

func (s MutablehtlcState) Secret() wasmtypes.ScMutableHash {
	return wasmtypes.NewScMutableHash(s.proxy.Root(StateSecret))
}

func (s MutablehtlcState) Time() wasmtypes.ScMutableInt64 {
	return wasmtypes.NewScMutableInt64(s.proxy.Root(StateTime))
}

func (s MutablehtlcState) Value() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.proxy.Root(StateValue))
}
