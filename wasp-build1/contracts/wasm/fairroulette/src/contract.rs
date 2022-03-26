// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]

use wasmlib::*;
use crate::*;

pub struct ForcePayoutCall {
	pub func: ScFunc,
}

pub struct ForceResetCall {
	pub func: ScFunc,
}

pub struct PayWinnersCall {
	pub func: ScFunc,
}

pub struct PlaceBetCall {
	pub func: ScFunc,
	pub params: MutablePlaceBetParams,
}

pub struct PlayPeriodCall {
	pub func: ScFunc,
	pub params: MutablePlayPeriodParams,
}

pub struct LastWinningNumberCall {
	pub func: ScView,
	pub results: ImmutableLastWinningNumberResults,
}

pub struct RoundNumberCall {
	pub func: ScView,
	pub results: ImmutableRoundNumberResults,
}

pub struct RoundStartedAtCall {
	pub func: ScView,
	pub results: ImmutableRoundStartedAtResults,
}

pub struct RoundStatusCall {
	pub func: ScView,
	pub results: ImmutableRoundStatusResults,
}

pub struct ScFuncs {
}

impl ScFuncs {
    pub fn force_payout(_ctx: &dyn ScFuncCallContext) -> ForcePayoutCall {
        ForcePayoutCall {
            func: ScFunc::new(HSC_NAME, HFUNC_FORCE_PAYOUT),
        }
    }

    pub fn force_reset(_ctx: &dyn ScFuncCallContext) -> ForceResetCall {
        ForceResetCall {
            func: ScFunc::new(HSC_NAME, HFUNC_FORCE_RESET),
        }
    }

    pub fn pay_winners(_ctx: &dyn ScFuncCallContext) -> PayWinnersCall {
        PayWinnersCall {
            func: ScFunc::new(HSC_NAME, HFUNC_PAY_WINNERS),
        }
    }

    pub fn place_bet(_ctx: &dyn ScFuncCallContext) -> PlaceBetCall {
        let mut f = PlaceBetCall {
            func: ScFunc::new(HSC_NAME, HFUNC_PLACE_BET),
            params: MutablePlaceBetParams { proxy: Proxy::nil() },
        };
        ScFunc::link_params(&mut f.params.proxy, &f.func);
        f
    }

    pub fn play_period(_ctx: &dyn ScFuncCallContext) -> PlayPeriodCall {
        let mut f = PlayPeriodCall {
            func: ScFunc::new(HSC_NAME, HFUNC_PLAY_PERIOD),
            params: MutablePlayPeriodParams { proxy: Proxy::nil() },
        };
        ScFunc::link_params(&mut f.params.proxy, &f.func);
        f
    }

    pub fn last_winning_number(_ctx: &dyn ScViewCallContext) -> LastWinningNumberCall {
        let mut f = LastWinningNumberCall {
            func: ScView::new(HSC_NAME, HVIEW_LAST_WINNING_NUMBER),
            results: ImmutableLastWinningNumberResults { proxy: Proxy::nil() },
        };
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn round_number(_ctx: &dyn ScViewCallContext) -> RoundNumberCall {
        let mut f = RoundNumberCall {
            func: ScView::new(HSC_NAME, HVIEW_ROUND_NUMBER),
            results: ImmutableRoundNumberResults { proxy: Proxy::nil() },
        };
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn round_started_at(_ctx: &dyn ScViewCallContext) -> RoundStartedAtCall {
        let mut f = RoundStartedAtCall {
            func: ScView::new(HSC_NAME, HVIEW_ROUND_STARTED_AT),
            results: ImmutableRoundStartedAtResults { proxy: Proxy::nil() },
        };
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn round_status(_ctx: &dyn ScViewCallContext) -> RoundStatusCall {
        let mut f = RoundStatusCall {
            func: ScView::new(HSC_NAME, HVIEW_ROUND_STATUS),
            results: ImmutableRoundStatusResults { proxy: Proxy::nil() },
        };
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }
}
