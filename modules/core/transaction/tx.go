package transaction

import (
	"encoding/json"
	"math/big"

	"github.com/polarysfoundation/polarys-core/modules/common"
)

type Tx struct {
	From  common.Address `json:"from"`
	To    common.Address `json:"to"`
	Value *big.Int       `json:"value"`
	Nonce uint64         `json:"nonce"`
	Data  []byte         `json:"data"`
}

func NewTx(from, to common.Address, value *big.Int, nonce uint64, data []byte) *Tx {
	return &Tx{
		From:  from,
		To:    to,
		Value: value,
		Nonce: nonce,
		Data:  data,
	}
}

func copyTx(tx *Tx) *Tx {
	return &Tx{
		From:  tx.From,
		To:    tx.To,
		Value: new(big.Int).Set(tx.Value),
		Nonce: tx.Nonce,
		Data:  tx.Data,
	}
}

func (tx *Tx) marshal() ([]byte, error) {
	return json.Marshal(tx)
}

func (tx *Tx) unmarshal(b []byte) error {
	return json.Unmarshal(b, &tx)
}
