package txpool

import (
	"math/big"
	"sync"

	"github.com/polarysfoundation/polarys-core/modules/common"
	"github.com/polarysfoundation/polarys-core/modules/core/transaction"
)

type Status int

const (
	Pending Status = iota
	Queued
	Failed
)

type TxPool struct {
	poolKey   common.Key
	feeBurned *big.Int
	signer    *TxSigner

	Transactions map[Status]*transaction.Transaction
	mutex        *sync.RWMutex
}

func InitTxPool(txSigner *TxSigner, localTx []*transaction.Transaction) *TxPool {
	
	return &TxPool{
		feeBurned: big.NewInt(0),
	}
}
