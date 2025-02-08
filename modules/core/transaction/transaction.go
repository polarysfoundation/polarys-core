package transaction

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"

	pm256 "github.com/polarysfoundation/pm-256"
	"github.com/polarysfoundation/polarys-core/modules/common"
)

type Transaction struct {
	tx        Tx
	hash      common.Hash
	signature []byte

	signerHash common.Hash
}

func NewTransaction(tx *Tx) *Transaction {
	transaction := &Transaction{}

	transaction.tx = *copyTx(tx)

	transaction.calculateHash()

	return transaction
}

func (tx *Transaction) AddSignerHash(h common.Hash) error {
	if h.IsEmpty() {
		return fmt.Errorf("error adding signer hash, hash is empty")
	}

	tx.signerHash = h

	return nil
}

func (tx *Transaction) MarshalJSON() ([]byte, error) {
	aux := struct {
		Tx         Tx          `json:"txBody"`
		Hash       common.Hash `json:"hash"`
		Signature  []byte      `json:"signature"`
		SignerHash common.Hash `json:"signerHash"`
	}{
		Tx:         tx.tx,
		Hash:       tx.hash,
		Signature:  tx.signature,
		SignerHash: tx.signerHash,
	}

	return json.Marshal(aux)
}

func (tx *Transaction) UnmarshalJSON(b []byte) error {
	aux := struct {
		Tx         Tx          `json:"txBody"`
		Hash       common.Hash `json:"hash"`
		Signature  []byte      `json:"signature"`
		SignerHash common.Hash `json:"signerHash"`
	}{}

	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}

	tx.hash = aux.Hash
	tx.signature = aux.Signature
	tx.signerHash = aux.SignerHash
	tx.tx = aux.Tx

	return nil

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
