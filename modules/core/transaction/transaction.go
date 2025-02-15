package transaction

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"

	"github.com/polarysfoundation/polarys-core/modules/common"
	"github.com/polarysfoundation/polarys-core/modules/crypto"
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

	return transaction
}

func (tx *Transaction) SignTransaction(signature []byte) *Transaction {
	auxTx := copyTransaction(tx)

	auxTx.signature = signature

	auxTx.calculateHash()

	return auxTx
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

	if len(tx.signature) == 0 {
		log.Print("empty signature cannot calculate hash")
		return
	}

	b, err := txCopy.marshal()
	if err != nil {
		log.Printf("error: %v", err)
		panic("error marshaling tx body")
	}

	buff := make([]byte, len(b)+len(tx.signature))

	copy(buff[len(b):], b)
	copy(buff[len(b)+len(tx.signature):], tx.signature)

	h := crypto.Pm256(buff)
	tx.hash = common.BytesToHash(h)
}

func (tx *Transaction) Hash() common.Hash    { return tx.hash }
func (tx *Transaction) From() common.Address { return tx.tx.From }
func (tx *Transaction) To() common.Address   { return tx.tx.To }
func (tx *Transaction) Value() *big.Int      { return tx.tx.Value }
func (tx *Transaction) Nonce() uint64        { return tx.tx.Nonce }
func (tx *Transaction) Signature() []byte    { return tx.signature }
func (tx *Transaction) Data() []byte         { return tx.tx.Data }
func (tx *Transaction) Type() Version        { return tx.tx.Type }
func (tx *Transaction) Payload() []byte      { return tx.tx.Payload }

func (tx *Transaction) Print() {
	log.Printf("From: %s", tx.From().String())
	log.Printf("To: %s", tx.To().String())
	log.Printf("Value: %d", tx.Value())
	log.Printf("Nonce: %d", tx.Nonce())
	log.Printf("Signature: %s", common.EncodeToString(tx.Signature()))
	log.Printf("Data: %s", common.EncodeToString(tx.Data()))
	log.Printf("Type: %d", tx.Type())
	log.Printf("Hash: %s", tx.Hash().String())
	log.Printf("Payload: %s", string(tx.Payload()))
}

func copyTransaction(tx *Transaction) *Transaction {
	return &Transaction{
		tx:         tx.tx,
		hash:       tx.hash,
		signature:  tx.signature,
		signerHash: tx.signerHash,
	}
}
