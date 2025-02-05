package block

import (
	"testing"
	"time"

	"github.com/polarysfoundation/polarys-core/modules/common"
	"github.com/stretchr/testify/assert"
)

func TestNewHeader(t *testing.T) {
	prevHash := common.Hash{0x01, 0x02, 0x03}
	height := uint64(1)
	nonce := uint64(2)
	gasPrice := uint64(3)
	gasLimit := uint64(4)
	difficulty := uint64(5)
	totalDifficulty := uint64(6)
	coinbase := common.Address{0x07, 0x08, 0x09}
	data := []byte{0x0a, 0x0b, 0x0c}
	extraData := []byte{0x0d, 0x0e, 0x0f}

	header := NewHeader(prevHash, height, nonce, gasPrice, gasLimit, difficulty, totalDifficulty, coinbase, data, extraData)

	assert.Equal(t, prevHash, header.PrevHash)
	assert.Equal(t, height, header.Height)
	assert.Equal(t, nonce, header.Nonce)
	assert.Equal(t, gasPrice, header.GasPrice)
	assert.Equal(t, gasLimit, header.GasLimit)
	assert.Equal(t, difficulty, header.Difficulty)
	assert.Equal(t, totalDifficulty, header.TotalDifficulty)
	assert.Equal(t, coinbase, header.Coinbase)
	assert.Equal(t, data, header.Data)
	assert.Equal(t, extraData, header.ExtraData)
	assert.Equal(t, 64, len(header.Signature))
	assert.WithinDuration(t, time.Unix(int64(header.Timestamp), 0), time.Now(), time.Second)
}

func TestCopyHeader(t *testing.T) {
	original := &Header{
		PrevHash:        common.Hash{0x01, 0x02, 0x03},
		Height:          uint64(1),
		Nonce:           uint64(2),
		Coinbase:        common.Address{0x07, 0x08, 0x09},
		GasPrice:        uint64(3),
		GasLimit:        uint64(4),
		Difficulty:      uint64(5),
		Timestamp:       uint64(time.Now().Unix()),
		Signature:       []byte{0x10, 0x11, 0x12},
		TotalDifficulty: uint64(6),
		Data:            []byte{0x0a, 0x0b, 0x0c},
		ExtraData:       []byte{0x0d, 0x0e, 0x0f},
	}

	copy := copyHeader(original)

	assert.Equal(t, original, copy)
}

func TestMarshalUnmarshal(t *testing.T) {
	header := &Header{
		PrevHash:        common.Hash{0x01, 0x02, 0x03},
		Height:          uint64(1),
		Nonce:           uint64(2),
		Coinbase:        common.Address{0x07, 0x08, 0x09},
		GasPrice:        uint64(3),
		GasLimit:        uint64(4),
		Difficulty:      uint64(5),
		Timestamp:       uint64(time.Now().Unix()),
		Signature:       []byte{0x10, 0x11, 0x12},
		TotalDifficulty: uint64(6),
		Data:            []byte{0x0a, 0x0b, 0x0c},
		ExtraData:       []byte{0x0d, 0x0e, 0x0f},
	}

	data, err := header.marshal()
	assert.NoError(t, err)

	var unmarshaledHeader Header
	err = unmarshaledHeader.unmarshal(data)
	assert.NoError(t, err)
	assert.Equal(t, header, &unmarshaledHeader)
}
