package vmimpl

import (
	"github.com/iotaledger/wasp/packages/vm"
	"github.com/iotaledger/wasp/packages/vm/gas"
	"github.com/iotaledger/wasp/packages/vm/vmexceptions"
)

func (reqctx *requestContext) GasBurnEnable(enable bool) {
	if enable && !reqctx.shouldChargeGasFee() {
		return
	}
	reqctx.gas.burnEnabled = enable
}

func (reqctx *requestContext) gasSetBudget(gasBudget, maxTokensToSpendForGasFee uint64) {
	reqctx.gas.budgetAdjusted = gasBudget
	reqctx.gas.maxTokensToSpendForGasFee = maxTokensToSpendForGasFee
	reqctx.gas.burned = 0
}

func (reqctx *requestContext) GasBurn(burnCode gas.BurnCode, par ...uint64) {
	if !reqctx.gas.burnEnabled {
		return
	}
	g := burnCode.Cost(par...)
	reqctx.gas.burnLog.Record(burnCode, g)
	reqctx.gas.burned += g

	if reqctx.gas.burned > reqctx.gas.budgetAdjusted {
		reqctx.gas.burned = reqctx.gas.budgetAdjusted // do not charge more than the limit set by the request
		// debug.PrintStack()
		panic(vm.ErrGasBudgetExceeded)
	}

	if reqctx.vm.blockGas.burned+reqctx.gas.burned > reqctx.vm.chainInfo.GasLimits.MaxGasPerBlock {
		panic(vmexceptions.ErrBlockGasLimitExceeded) // panic if the current request gas overshoots the block limit
	}
}

func (reqctx *requestContext) GasBudgetLeft() uint64 {
	if reqctx.gas.budgetAdjusted < reqctx.gas.burned {
		return 0
	}
	return reqctx.gas.budgetAdjusted - reqctx.gas.burned
}

func (reqctx *requestContext) GasBurned() uint64 {
	return reqctx.gas.burned
}
