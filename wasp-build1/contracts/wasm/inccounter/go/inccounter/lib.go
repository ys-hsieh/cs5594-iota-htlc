// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package inccounter

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib"

var exportMap = wasmlib.ScExportMap{
	Names: []string{
		FuncCallIncrement,
		FuncCallIncrementRecurse5x,
		FuncEndlessLoop,
		FuncIncrement,
		FuncIncrementWithDelay,
		FuncInit,
		FuncLocalStateInternalCall,
		FuncLocalStatePost,
		FuncLocalStateSandboxCall,
		FuncPostIncrement,
		FuncRepeatMany,
		FuncTestVliCodec,
		FuncTestVluCodec,
		FuncWhenMustIncrement,
		ViewGetCounter,
		ViewGetVli,
		ViewGetVlu,
	},
	Funcs: []wasmlib.ScFuncContextFunction{
		funcCallIncrementThunk,
		funcCallIncrementRecurse5xThunk,
		funcEndlessLoopThunk,
		funcIncrementThunk,
		funcIncrementWithDelayThunk,
		funcInitThunk,
		funcLocalStateInternalCallThunk,
		funcLocalStatePostThunk,
		funcLocalStateSandboxCallThunk,
		funcPostIncrementThunk,
		funcRepeatManyThunk,
		funcTestVliCodecThunk,
		funcTestVluCodecThunk,
		funcWhenMustIncrementThunk,
	},
	Views: []wasmlib.ScViewContextFunction{
		viewGetCounterThunk,
		viewGetVliThunk,
		viewGetVluThunk,
	},
}

func OnLoad(index int32) {
	if index >= 0 {
		wasmlib.ScExportsCall(index, &exportMap)
		return
	}

	wasmlib.ScExportsExport(&exportMap)
}

type CallIncrementContext struct {
	State MutableIncCounterState
}

func funcCallIncrementThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("inccounter.funcCallIncrement")
	f := &CallIncrementContext{
		State: MutableIncCounterState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	funcCallIncrement(ctx, f)
	ctx.Log("inccounter.funcCallIncrement ok")
}

type CallIncrementRecurse5xContext struct {
	State MutableIncCounterState
}

func funcCallIncrementRecurse5xThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("inccounter.funcCallIncrementRecurse5x")
	f := &CallIncrementRecurse5xContext{
		State: MutableIncCounterState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	funcCallIncrementRecurse5x(ctx, f)
	ctx.Log("inccounter.funcCallIncrementRecurse5x ok")
}

type EndlessLoopContext struct {
	State MutableIncCounterState
}

func funcEndlessLoopThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("inccounter.funcEndlessLoop")
	f := &EndlessLoopContext{
		State: MutableIncCounterState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	funcEndlessLoop(ctx, f)
	ctx.Log("inccounter.funcEndlessLoop ok")
}

type IncrementContext struct {
	State MutableIncCounterState
}

func funcIncrementThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("inccounter.funcIncrement")
	f := &IncrementContext{
		State: MutableIncCounterState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	funcIncrement(ctx, f)
	ctx.Log("inccounter.funcIncrement ok")
}

type IncrementWithDelayContext struct {
	Params ImmutableIncrementWithDelayParams
	State  MutableIncCounterState
}

func funcIncrementWithDelayThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("inccounter.funcIncrementWithDelay")
	f := &IncrementWithDelayContext{
		Params: ImmutableIncrementWithDelayParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutableIncCounterState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.Delay().Exists(), "missing mandatory delay")
	funcIncrementWithDelay(ctx, f)
	ctx.Log("inccounter.funcIncrementWithDelay ok")
}

type InitContext struct {
	Params ImmutableInitParams
	State  MutableIncCounterState
}

func funcInitThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("inccounter.funcInit")
	f := &InitContext{
		Params: ImmutableInitParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutableIncCounterState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	funcInit(ctx, f)
	ctx.Log("inccounter.funcInit ok")
}

type LocalStateInternalCallContext struct {
	State MutableIncCounterState
}

func funcLocalStateInternalCallThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("inccounter.funcLocalStateInternalCall")
	f := &LocalStateInternalCallContext{
		State: MutableIncCounterState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	funcLocalStateInternalCall(ctx, f)
	ctx.Log("inccounter.funcLocalStateInternalCall ok")
}

type LocalStatePostContext struct {
	State MutableIncCounterState
}

func funcLocalStatePostThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("inccounter.funcLocalStatePost")
	f := &LocalStatePostContext{
		State: MutableIncCounterState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	funcLocalStatePost(ctx, f)
	ctx.Log("inccounter.funcLocalStatePost ok")
}

type LocalStateSandboxCallContext struct {
	State MutableIncCounterState
}

func funcLocalStateSandboxCallThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("inccounter.funcLocalStateSandboxCall")
	f := &LocalStateSandboxCallContext{
		State: MutableIncCounterState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	funcLocalStateSandboxCall(ctx, f)
	ctx.Log("inccounter.funcLocalStateSandboxCall ok")
}

type PostIncrementContext struct {
	State MutableIncCounterState
}

func funcPostIncrementThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("inccounter.funcPostIncrement")
	f := &PostIncrementContext{
		State: MutableIncCounterState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	funcPostIncrement(ctx, f)
	ctx.Log("inccounter.funcPostIncrement ok")
}

type RepeatManyContext struct {
	Params ImmutableRepeatManyParams
	State  MutableIncCounterState
}

func funcRepeatManyThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("inccounter.funcRepeatMany")
	f := &RepeatManyContext{
		Params: ImmutableRepeatManyParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutableIncCounterState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	funcRepeatMany(ctx, f)
	ctx.Log("inccounter.funcRepeatMany ok")
}

type TestVliCodecContext struct {
	State MutableIncCounterState
}

func funcTestVliCodecThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("inccounter.funcTestVliCodec")
	f := &TestVliCodecContext{
		State: MutableIncCounterState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	funcTestVliCodec(ctx, f)
	ctx.Log("inccounter.funcTestVliCodec ok")
}

type TestVluCodecContext struct {
	State MutableIncCounterState
}

func funcTestVluCodecThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("inccounter.funcTestVluCodec")
	f := &TestVluCodecContext{
		State: MutableIncCounterState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	funcTestVluCodec(ctx, f)
	ctx.Log("inccounter.funcTestVluCodec ok")
}

type WhenMustIncrementContext struct {
	Params ImmutableWhenMustIncrementParams
	State  MutableIncCounterState
}

func funcWhenMustIncrementThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("inccounter.funcWhenMustIncrement")
	f := &WhenMustIncrementContext{
		Params: ImmutableWhenMustIncrementParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutableIncCounterState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	funcWhenMustIncrement(ctx, f)
	ctx.Log("inccounter.funcWhenMustIncrement ok")
}

type GetCounterContext struct {
	Results MutableGetCounterResults
	State   ImmutableIncCounterState
}

func viewGetCounterThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("inccounter.viewGetCounter")
	results := wasmlib.NewScDict()
	f := &GetCounterContext{
		Results: MutableGetCounterResults{
			proxy: results.AsProxy(),
		},
		State: ImmutableIncCounterState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	viewGetCounter(ctx, f)
	ctx.Results(results)
	ctx.Log("inccounter.viewGetCounter ok")
}

type GetVliContext struct {
	Params  ImmutableGetVliParams
	Results MutableGetVliResults
	State   ImmutableIncCounterState
}

func viewGetVliThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("inccounter.viewGetVli")
	results := wasmlib.NewScDict()
	f := &GetVliContext{
		Params: ImmutableGetVliParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		Results: MutableGetVliResults{
			proxy: results.AsProxy(),
		},
		State: ImmutableIncCounterState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.Ni64().Exists(), "missing mandatory ni64")
	viewGetVli(ctx, f)
	ctx.Results(results)
	ctx.Log("inccounter.viewGetVli ok")
}

type GetVluContext struct {
	Params  ImmutableGetVluParams
	Results MutableGetVluResults
	State   ImmutableIncCounterState
}

func viewGetVluThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("inccounter.viewGetVlu")
	results := wasmlib.NewScDict()
	f := &GetVluContext{
		Params: ImmutableGetVluParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		Results: MutableGetVluResults{
			proxy: results.AsProxy(),
		},
		State: ImmutableIncCounterState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.Nu64().Exists(), "missing mandatory nu64")
	viewGetVlu(ctx, f)
	ctx.Results(results)
	ctx.Log("inccounter.viewGetVlu ok")
}