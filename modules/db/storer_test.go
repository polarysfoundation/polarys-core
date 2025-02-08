package db_test

import (
	"log"
	"testing"

	"github.com/polarysfoundation/polarys-core/modules/common"
	"github.com/polarysfoundation/polarys-core/modules/core/block"
	"github.com/polarysfoundation/polarys-core/modules/db"
)

func TestWriteBlock(t *testing.T) {
	storer := db.NewChainStorer()

	header := &block.Header{
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

	block := block.NewBlock(header, nil)

	if block.Hash().IsEmpty() {
		t.Errorf("calculateHash failed, hash is empty")
	} else {
		log.Printf("Block hash: %s", block.Hash().String())
	}

	block.Print()

	err := storer.WriteBlock(block)
	if err != nil {
		t.Errorf("Failed to write block: %v", err)
	} else {
		log.Println("Block written successfully")
	}
}

func TestReadBlock(t *testing.T) {
	storer := db.NewChainStorer()

	header := &block.Header{
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

	block := block.NewBlock(header, nil)

	if block.Hash().IsEmpty() {
		t.Errorf("calculateHash failed, hash is empty")
	} else {
		log.Printf("Block hash: %s", block.Hash().String())
	}

	err := storer.WriteBlock(block)
	if err != nil {
		t.Errorf("Failed to write block: %v", err)
	} else {
		log.Println("Block written successfully")
	}

	readBlock, err := storer.GetBlockByHeight(block.Height())
	if err != nil {
		t.Errorf("Failed to read block: %v", err)
	} else {
		log.Printf("Block read successfully: %v", readBlock)
		readBlock.Print()
	}
}

func TestGetCurrentBlock(t *testing.T) {
	storer := db.NewChainStorer()

	header := &block.Header{
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

	block := block.NewBlock(header, nil)

	err := storer.WriteBlock(block)
	if err != nil {
		t.Errorf("Failed to write block: %v", err)
		return
	}

	currentBlock, err := storer.GetCurrentBlock()
	if err != nil {
		t.Errorf("Failed to get current block: %v", err)
	} else {
		log.Printf("Current block retrieved successfully: %v", currentBlock)
		currentBlock.Print()
	}

	if currentBlock.Hash() != block.Hash() {
		t.Errorf("Current block hash does not match written block hash")
	}
}
