package zksync

import "github.com/ethereum/go-ethereum/common"

type ContractAddress struct {
	MainContract string `json:"mainContract"`
	GovContract  string `json:"govContract"`
}

func (a *ContractAddress) GetMainAddress() common.Address {
	return common.HexToAddress(a.MainContract)
}

func (a *ContractAddress) GetGovAddress() common.Address {
	return common.HexToAddress(a.GovContract)
}
