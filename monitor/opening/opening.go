package opening

import (
	etype "github.com/ethereum/go-ethereum/core/types"
	"github.com/wslzwps/go-web3"
)

type IMonitor interface {
	isOpen(trx *etype.Transaction) bool
}

type OpenMonitor struct {
	openType     int
	web3         web3.Web3
	monitor IMonitor
}

func NewOpenMonitor(t int, web3 web3.Web3) (obj *OpenMonitor) {
	obj = new(OpenMonitor)
	obj.openType = t
	obj.web3 = web3

	switch obj.openType {
	case 0:
		obj.monitor = nil
	case 1:
		obj.monitor = nil
	case 2:
		obj.monitor = nil
	case 3:
		obj.monitor = nil
	}

	return
}

func (obj *OpenMonitor) IsOpening(trx *etype.Transaction) bool {
	return obj.monitor.isOpen(trx)
}
