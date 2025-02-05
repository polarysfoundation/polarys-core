package transaction

import (
	"log"
	"math/big"

	pm256 "github.com/polarysfoundation/pm-256"
	"github.com/polarysfoundation/polarys-core/modules/common"
)

type Transaction struct {
	tx        Tx
	hash      common.Hash
	signature []byte
}

func NewTransaction(tx *Tx) *Transaction {
	transaction := &Transaction{}

	transaction.tx = *copyTx(tx)

	transaction.calculateHash()

	return transaction
}

func (tx *Transaction) calculateHash() {
	txCopy := &tx.tx

	var buf common.Hash
	d, err := txCopy.marshal()
	if err != nil {
		log.Printf("error: %v", err)
		panic("error marshaling header block")
	}

	h := pm256.New256()
	h.Write(d)
	h.Sum(buf[:0])
	tx.hash = buf
}

func (tx *Transaction) Hash() common.Hash    { return tx.hash }
func (tx *Transaction) From() common.Address { return tx.tx.From }
func (tx *Transaction) To() common.Address   { return tx.tx.To }
func (tx *Transaction) Value() *big.Int      { return tx.tx.Value }
func (tx *Transaction) Nonce() uint64        { return tx.tx.Nonce }
func (tx *Transaction) Signature() []byte    { return tx.signature }
func (tx *Transaction) Data() []byte         { return tx.tx.Data }
