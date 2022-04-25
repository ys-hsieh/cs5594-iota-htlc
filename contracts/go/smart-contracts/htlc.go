// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package htlc

import "time"
import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib"
import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

func funcInit(ctx wasmlib.ScFuncContext, f *InitContext) {
    if f.Params.Owner().Exists() {
        f.State.Owner().SetValue(f.Params.Owner().Value())
        return
    }
    f.State.Owner().SetValue(ctx.ContractCreator())
    f.State.Value().SetValue(0)
    f.State.InitTime().SetValue(time.Now().Unix())
}

func funcSetOwner(ctx wasmlib.ScFuncContext, f *SetOwnerContext) {
	f.State.Owner().SetValue(f.Params.Owner().Value())
}

func funcSetReceivder(ctx wasmlib.ScFuncContext, f *SetReceivderContext) {
    f.State.Receivder().SetValue(f.Params.Receivder().Value())
}

func funcSetSecret(ctx wasmlib.ScFuncContext, f *SetSecretContext) {
    f.State.Secret().SetValue(f.Params.Secret().Value())
}

func funcSetTime(ctx wasmlib.ScFuncContext, f *SetTimeContext) {
    f.State.Time().SetValue(f.Params.Time().Value())
}

func funcSetValue(ctx wasmlib.ScFuncContext, f *SetValueContext) {
    f.State.Value().SetValue(f.Params.Value().Value())
}

func funcTransfer(ctx wasmlib.ScFuncContext, f *TransferContext) {
    if (time.Now().Unix() <= f.State.InitTime().Value() + f.State.Time().Value()) {
        if (f.Params.Secret().Value() == f.State.Secret().Value()) {
            address := wasmtypes.AddressFromBytes(f.State.Receivder().Value().Bytes())
            transfers := wasmlib.NewScTransferIotas(f.State.Value().Value())
            ctx.Send(address, transfers)
        }
    }
}

func funcWithdraw(ctx wasmlib.ScFuncContext, f *WithdrawContext) {
    if (time.Now().Unix() > f.State.InitTime().Value() + f.State.Time().Value()) {
        address := wasmtypes.AddressFromBytes(f.State.Owner().Value().Bytes())
        transfers := wasmlib.NewScTransferIotas(f.State.Value().Value())
        ctx.Send(address, transfers)
    }
}

func viewGetOwner(ctx wasmlib.ScViewContext, f *GetOwnerContext) {
	f.Results.Owner().SetValue(f.State.Owner().Value())
}

func viewGetValue(ctx wasmlib.ScViewContext, f *GetValueContext) {
    f.Results.Value().SetValue(f.State.Value().Value())
}
