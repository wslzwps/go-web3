package open

import (
	etype "github.com/ethereum/go-ethereum/core/types"
	"github.com/wslzwps/go-web3"
)

type IOpenChecker interface {
	CheckIfOpen(trx *etype.Transaction) bool
}

type Checker struct {
	openType int
	checker  IOpenChecker
}

func NewOpenMonitor(t int, web3 web3.Web3) (obj *Checker) {
	obj = new(Checker)
	obj.openType = t

	switch obj.openType {
	case 0:
		obj.checker = nil
	case 1:
		obj.checker = nil
	case 2:
		obj.checker = nil
	case 3:
		obj.checker = nil
	}

	return
}

func (obj *Checker) CheckIfOpen(trx *etype.Transaction) bool {
	return obj.checker.CheckIfOpen(trx)
}
