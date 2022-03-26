// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmtypes from "wasmlib/wasmtypes";

export class Donation {
	amount    : u64 = 0;  // amount donated
	donator   : wasmtypes.ScAgentID = wasmtypes.agentIDFromBytes([]);  // who donated
	error     : string = "";  // error to be reported to donator if anything goes wrong
	feedback  : string = "";  // the feedback for the person donated to
	timestamp : u64 = 0;  // when the donation took place

	static fromBytes(buf: u8[]): Donation {
		const dec = new wasmtypes.WasmDecoder(buf);
		const data = new Donation();
		data.amount    = wasmtypes.uint64Decode(dec);
		data.donator   = wasmtypes.agentIDDecode(dec);
		data.error     = wasmtypes.stringDecode(dec);
		data.feedback  = wasmtypes.stringDecode(dec);
		data.timestamp = wasmtypes.uint64Decode(dec);
		dec.close();
		return data;
	}

	bytes(): u8[] {
		const enc = new wasmtypes.WasmEncoder();
		wasmtypes.uint64Encode(enc, this.amount);
		wasmtypes.agentIDEncode(enc, this.donator);
		wasmtypes.stringEncode(enc, this.error);
		wasmtypes.stringEncode(enc, this.feedback);
		wasmtypes.uint64Encode(enc, this.timestamp);
		return enc.buf();
	}
}

export class ImmutableDonation extends wasmtypes.ScProxy {

	exists(): bool {
		return this.proxy.exists();
	}

	value(): Donation {
		return Donation.fromBytes(this.proxy.get());
	}
}

export class MutableDonation extends wasmtypes.ScProxy {

	delete(): void {
		this.proxy.delete();
	}

	exists(): bool {
		return this.proxy.exists();
	}

	setValue(value: Donation): void {
		this.proxy.set(value.bytes());
	}

	value(): Donation {
		return Donation.fromBytes(this.proxy.get());
	}
}