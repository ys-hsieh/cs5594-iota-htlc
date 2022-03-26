// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]
#![allow(unused_imports)]

use wasmlib::*;
use crate::*;

#[derive(Clone)]
pub struct ImmutableCallOnChainResults {
	pub(crate) proxy: Proxy,
}

impl ImmutableCallOnChainResults {
    pub fn int_value(&self) -> ScImmutableInt64 {
		ScImmutableInt64::new(self.proxy.root(RESULT_INT_VALUE))
	}
}

#[derive(Clone)]
pub struct MutableCallOnChainResults {
	pub(crate) proxy: Proxy,
}

impl MutableCallOnChainResults {
    pub fn int_value(&self) -> ScMutableInt64 {
		ScMutableInt64::new(self.proxy.root(RESULT_INT_VALUE))
	}
}

#[derive(Clone)]
pub struct ImmutableGetMintedSupplyResults {
	pub(crate) proxy: Proxy,
}

impl ImmutableGetMintedSupplyResults {
    pub fn minted_color(&self) -> ScImmutableColor {
		ScImmutableColor::new(self.proxy.root(RESULT_MINTED_COLOR))
	}

    pub fn minted_supply(&self) -> ScImmutableUint64 {
		ScImmutableUint64::new(self.proxy.root(RESULT_MINTED_SUPPLY))
	}
}

#[derive(Clone)]
pub struct MutableGetMintedSupplyResults {
	pub(crate) proxy: Proxy,
}

impl MutableGetMintedSupplyResults {
    pub fn minted_color(&self) -> ScMutableColor {
		ScMutableColor::new(self.proxy.root(RESULT_MINTED_COLOR))
	}

    pub fn minted_supply(&self) -> ScMutableUint64 {
		ScMutableUint64::new(self.proxy.root(RESULT_MINTED_SUPPLY))
	}
}

#[derive(Clone)]
pub struct ImmutableRunRecursionResults {
	pub(crate) proxy: Proxy,
}

impl ImmutableRunRecursionResults {
    pub fn int_value(&self) -> ScImmutableInt64 {
		ScImmutableInt64::new(self.proxy.root(RESULT_INT_VALUE))
	}
}

#[derive(Clone)]
pub struct MutableRunRecursionResults {
	pub(crate) proxy: Proxy,
}

impl MutableRunRecursionResults {
    pub fn int_value(&self) -> ScMutableInt64 {
		ScMutableInt64::new(self.proxy.root(RESULT_INT_VALUE))
	}
}

#[derive(Clone)]
pub struct ImmutableTestChainOwnerIDFullResults {
	pub(crate) proxy: Proxy,
}

impl ImmutableTestChainOwnerIDFullResults {
    pub fn chain_owner_id(&self) -> ScImmutableAgentID {
		ScImmutableAgentID::new(self.proxy.root(RESULT_CHAIN_OWNER_ID))
	}
}

#[derive(Clone)]
pub struct MutableTestChainOwnerIDFullResults {
	pub(crate) proxy: Proxy,
}

impl MutableTestChainOwnerIDFullResults {
    pub fn chain_owner_id(&self) -> ScMutableAgentID {
		ScMutableAgentID::new(self.proxy.root(RESULT_CHAIN_OWNER_ID))
	}
}

#[derive(Clone)]
pub struct ImmutableFibonacciResults {
	pub(crate) proxy: Proxy,
}

impl ImmutableFibonacciResults {
    pub fn int_value(&self) -> ScImmutableInt64 {
		ScImmutableInt64::new(self.proxy.root(RESULT_INT_VALUE))
	}
}

#[derive(Clone)]
pub struct MutableFibonacciResults {
	pub(crate) proxy: Proxy,
}

impl MutableFibonacciResults {
    pub fn int_value(&self) -> ScMutableInt64 {
		ScMutableInt64::new(self.proxy.root(RESULT_INT_VALUE))
	}
}

#[derive(Clone)]
pub struct ImmutableGetCounterResults {
	pub(crate) proxy: Proxy,
}

impl ImmutableGetCounterResults {
    pub fn counter(&self) -> ScImmutableInt64 {
		ScImmutableInt64::new(self.proxy.root(RESULT_COUNTER))
	}
}

#[derive(Clone)]
pub struct MutableGetCounterResults {
	pub(crate) proxy: Proxy,
}

impl MutableGetCounterResults {
    pub fn counter(&self) -> ScMutableInt64 {
		ScMutableInt64::new(self.proxy.root(RESULT_COUNTER))
	}
}

#[derive(Clone)]
pub struct MapStringToImmutableInt64 {
	pub(crate) proxy: Proxy,
}

impl MapStringToImmutableInt64 {
    pub fn get_int64(&self, key: &str) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.proxy.key(&string_to_bytes(key)))
    }
}

#[derive(Clone)]
pub struct ImmutableGetIntResults {
	pub(crate) proxy: Proxy,
}

impl ImmutableGetIntResults {
    pub fn values(&self) -> MapStringToImmutableInt64 {
		MapStringToImmutableInt64 { proxy: self.proxy.clone() }
	}
}

#[derive(Clone)]
pub struct MapStringToMutableInt64 {
	pub(crate) proxy: Proxy,
}

impl MapStringToMutableInt64 {
    pub fn clear(&self) {
        self.proxy.clear_map();
    }

    pub fn get_int64(&self, key: &str) -> ScMutableInt64 {
        ScMutableInt64::new(self.proxy.key(&string_to_bytes(key)))
    }
}

#[derive(Clone)]
pub struct MutableGetIntResults {
	pub(crate) proxy: Proxy,
}

impl MutableGetIntResults {
    pub fn values(&self) -> MapStringToMutableInt64 {
		MapStringToMutableInt64 { proxy: self.proxy.clone() }
	}
}

#[derive(Clone)]
pub struct MapStringToImmutableString {
	pub(crate) proxy: Proxy,
}

impl MapStringToImmutableString {
    pub fn get_string(&self, key: &str) -> ScImmutableString {
        ScImmutableString::new(self.proxy.key(&string_to_bytes(key)))
    }
}

#[derive(Clone)]
pub struct ImmutableGetStringValueResults {
	pub(crate) proxy: Proxy,
}

impl ImmutableGetStringValueResults {
    pub fn vars(&self) -> MapStringToImmutableString {
		MapStringToImmutableString { proxy: self.proxy.clone() }
	}
}

#[derive(Clone)]
pub struct MapStringToMutableString {
	pub(crate) proxy: Proxy,
}

impl MapStringToMutableString {
    pub fn clear(&self) {
        self.proxy.clear_map();
    }

    pub fn get_string(&self, key: &str) -> ScMutableString {
        ScMutableString::new(self.proxy.key(&string_to_bytes(key)))
    }
}

#[derive(Clone)]
pub struct MutableGetStringValueResults {
	pub(crate) proxy: Proxy,
}

impl MutableGetStringValueResults {
    pub fn vars(&self) -> MapStringToMutableString {
		MapStringToMutableString { proxy: self.proxy.clone() }
	}
}

#[derive(Clone)]
pub struct ImmutableTestChainOwnerIDViewResults {
	pub(crate) proxy: Proxy,
}

impl ImmutableTestChainOwnerIDViewResults {
    pub fn chain_owner_id(&self) -> ScImmutableAgentID {
		ScImmutableAgentID::new(self.proxy.root(RESULT_CHAIN_OWNER_ID))
	}
}

#[derive(Clone)]
pub struct MutableTestChainOwnerIDViewResults {
	pub(crate) proxy: Proxy,
}

impl MutableTestChainOwnerIDViewResults {
    pub fn chain_owner_id(&self) -> ScMutableAgentID {
		ScMutableAgentID::new(self.proxy.root(RESULT_CHAIN_OWNER_ID))
	}
}

#[derive(Clone)]
pub struct ImmutableTestSandboxCallResults {
	pub(crate) proxy: Proxy,
}

impl ImmutableTestSandboxCallResults {
    pub fn sandbox_call(&self) -> ScImmutableString {
		ScImmutableString::new(self.proxy.root(RESULT_SANDBOX_CALL))
	}
}

#[derive(Clone)]
pub struct MutableTestSandboxCallResults {
	pub(crate) proxy: Proxy,
}

impl MutableTestSandboxCallResults {
    pub fn sandbox_call(&self) -> ScMutableString {
		ScMutableString::new(self.proxy.root(RESULT_SANDBOX_CALL))
	}
}