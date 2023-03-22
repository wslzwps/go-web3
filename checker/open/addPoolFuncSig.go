package open

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/wslzwps/go-web3"
)

type AddPoolFunc struct {
	tokenAddr string
	web3      web3.Web3
}

func NewAddPoolFunc(tokenAddr string, web3 web3.Web3) (obj *AddPoolFunc) {
	obj = new(AddPoolFunc)
	obj.tokenAddr = tokenAddr
	obj.web3 = web3
	return
}

func (obj *AddPoolFunc) checkIfOpen(trx *types.Transaction) bool {
	//trx.Data(), trx.To(), trx.From() are all available
	return false
}
