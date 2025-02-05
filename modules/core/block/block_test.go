package block

import (
	"testing"
	"time"

	"github.com/polarysfoundation/polarys-core/modules/common"
)

func TestCalculateHash(t *testing.T) {
	header := &Header{
		PrevHash:        common.BytesToHash([]byte{1, 2, 3, 4}),
		Height:          1,
		Nonce:           12345,
		Coinbase:        common.BytesToAddress([]byte{5, 6, 7, 8}),
		GasPrice:        1000,
		GasLimit:        5000,
		Difficulty:      2,
		Timestamp:       1633024800,
		Signature:       []byte{9, 10, 11, 12},
		TotalDifficulty: 3,
		Data:            []byte("block data"),
		ExtraData:       []byte("extra data"),
	}

	block := NewBlock(header)
	block.calculateHash()

	if block.hash.IsEmpty() {
		t.Errorf("calculateHash failed, hash is empty")
	}

	expectedHash := common.HexToHash("b6f8829aec00e05173333e6ebfeea6c717ec869336aeb3889f0655c3214aee5a")
	t.Logf("Expected hash: %v", expectedHash.ToHex())
	t.Logf("Computed hash: %v", block.hash.ToHex())

	if block.hash != expectedHash {
		t.Errorf("calculateHash failed, expected %v, got %v", expectedHash, block.hash)
	}
}

func TestCalculateHashWithExistingHash(t *testing.T) {
	header := &Header{
		PrevHash:        common.BytesToHash([]byte{1, 2, 3, 4}),
		Height:          1,
		Nonce:           12345,
		Coinbase:        common.BytesToAddress([]byte{5, 6, 7, 8}),
		GasPrice:        1000,
		GasLimit:        5000,
		Difficulty:      2,
		Timestamp:       1633024800,
		Signature:       []byte{9, 10, 11, 12},
		TotalDifficulty: 3,
		Data:            []byte("block data"),
		ExtraData:       []byte("extra data"),
	}

	block := NewBlock(header)
	initialHash := common.BytesToHash([]byte{13, 14, 15, 16})
	block.hash = initialHash
	block.calculateHash()

	if block.hash == initialHash {
		t.Errorf("calculateHash failed, expected hash to be recomputed, got %v", block.hash)
	}
}

func TestCalculateHashForMiningBlock(t *testing.T) {
	header := &Header{
		PrevHash:        common.BytesToHash([]byte{1, 2, 3, 4}),
		Height:          1,
		Nonce:           12345,
		Coinbase:        common.BytesToAddress([]byte{5, 6, 7, 8}),
		GasPrice:        1000,
		GasLimit:        5000,
		Difficulty:      2,
		Timestamp:       1633024800,
		Signature:       []byte{9, 10, 11, 12},
		TotalDifficulty: 3,
		Data:            []byte("block data"),
		ExtraData:       []byte("extra data"),
	}

	block := NewBlock(header)
	block.calculateHash()

	if block.hash.IsEmpty() {
		t.Errorf("calculateHash failed, hash is empty")
	}

	t.Logf("Computed hash: %v", block.hash.ToHex())

	expectedHash := common.HexToHash("b6f8829aec00e05173333e6ebfeea6c717ec869336aeb3889f0655c3214aee5a")
	if block.hash != expectedHash {
		t.Errorf("calculateHash failed, expected %v, got %v", expectedHash, block.hash)
	}
}

func TestMineBlocks(t *testing.T) {
	for i := 0; i < 1000; i++ {
		header := &Header{
			PrevHash:        common.BytesToHash([]byte{1, 2, 3, 4}),
			Height:          uint64(i + 1),
			Nonce:           uint64(i * 12345),
			Coinbase:        common.BytesToAddress([]byte{5, 6, 7, 8}),
			GasPrice:        1000,
			GasLimit:        5000,
			Difficulty:      2,
			Timestamp:       uint64(time.Now().Unix()),
			Signature:       []byte{9, 10, 11, 12},
			TotalDifficulty: 3,
			Data:            []byte("block data"),
			ExtraData:       []byte("extra data"),
		}

		block := NewBlock(header)
		block.calculateHash()

		t.Logf("block hash: %s", block.Hash().ToHex())

		if block.hash.IsEmpty() {
			t.Errorf("calculateHash failed at block %d, hash is empty", i+1)
		}
	}
}
