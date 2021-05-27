package registry_pkg

import (
	"github.com/iotaledger/goshimmer/packages/ledgerstate"
	"github.com/iotaledger/wasp/packages/parameters"
	flag "github.com/spf13/pflag"
)

const (
	// CfgBindAddress defines the config flag of the web API binding address.
	CfgRewardAddress = "reward.address"
)

func InitFlags() {
	flag.String(CfgRewardAddress, "", "reward address for this Wasp node. Empty (default) means no rewards are collected")
}

func GetFeeDestination(chainID *ledgerstate.AliasAddress) ledgerstate.Address {
	//TODO
	ret, err := ledgerstate.AddressFromBase58EncodedString(parameters.GetString(CfgRewardAddress))
	if err != nil {
		return nil
	}
	return ret
}
