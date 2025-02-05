package block

import (
	"encoding/json"
	"time"

	"github.com/polarysfoundation/polarys-core/modules/common"
)

type Header struct {
	PrevHash        common.Hash    `json:"prevHash"`
	Height          uint64         `json:"height"`
	Nonce           uint64         `json:"nonce"`
	Coinbase        common.Address `json:"coinbase"`
	GasPrice        uint64         `json:"gasPrice"`
	GasLimit        uint64         `json:"gasLimit"`
	Difficulty      uint64         `json:"difficulty"`
	Timestamp       uint64         `json:"timestamp"`
	Signature       []byte         `json:"signature"`
	TotalDifficulty uint64         `json:"totalDifficulty"`
	Data            []byte         `json:"data"`
	ExtraData       []byte         `json:"extraData"`
}

func NewHeader(prevHash common.Hash, height, nonce, gasPrice, gasLimit, difficulty, totalDifficulty uint64, coinbase common.Address, data, extraData []byte) *Header {
	return &Header{
		PrevHash:        prevHash,
		Height:          height,
		Nonce:           nonce,
		Coinbase:        coinbase,
		GasPrice:        gasPrice,
		GasLimit:        gasLimit,
		Difficulty:      difficulty,
		Timestamp:       uint64(time.Now().Unix()),
		Signature:       make([]byte, 64),
		TotalDifficulty: totalDifficulty,
		Data:            data,
		ExtraData:       extraData,
	}
}

func copyHeader(header *Header) *Header {
	return &Header{
		PrevHash:        header.PrevHash,
		Height:          header.Height,
		Nonce:           header.Nonce,
		Coinbase:        header.Coinbase,
		GasPrice:        header.GasPrice,
		GasLimit:        header.GasLimit,
		Difficulty:      header.Difficulty,
		Timestamp:       header.Timestamp,
		Signature:       header.Signature,
		TotalDifficulty: header.TotalDifficulty,
		Data:            header.Data,
		ExtraData:       header.ExtraData,
	}
}

func (h Header) marshal() ([]byte, error) {
	return json.Marshal(h)
}

func (h *Header) unmarshal(b []byte) error {
	return json.Unmarshal(b, h)
}
