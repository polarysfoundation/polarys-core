package db

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/polarysfoundation/polarys-core/modules/common"
	"github.com/polarysfoundation/polarys-core/modules/core/block"
	polarysdb "github.com/polarysfoundation/polarys_db"
)

const (
	blockByHeight        = "/chain/blocks/height/"
	blockByHash          = "/chain/blocks/hash/"
	latestBlock          = "/chain/blocks/"
	transactionsBlocks   = "/chain/blocks/transactions/"
	transactionsAccounts = "/chain/transactions/accounts/"

	blocksMiner = "/chain/blocks/miner/"
)

type ChainStorer struct {
	db *polarysdb.Database
}

func NewChainStorer() *ChainStorer {
	db, err := polarysdb.Init("", ".polarys")
	if err != nil {
		log.Printf("error: %v", err)
		panic("error initializing dabatase")
	}

	return &ChainStorer{
		db: db,
	}
}

func (s *ChainStorer) WriteBlock(block *block.Block) error {
	writeOperations := []struct {
		table string
		key   string
		value interface{}
	}{
		{blockByHeight, fmt.Sprintf("block_%d", block.Height()), block},
		{blockByHash, fmt.Sprintf("block_%d", block.Hash()), block},
		{latestBlock, "currentBlock", block},
	}

	for _, op := range writeOperations {
		if !s.db.Exist(op.key) {
			err := s.db.Create(op.table, op.key, op.value)
			if err != nil {
				return fmt.Errorf("error creating new state for key %s, err: %v", op.key, err)
			}
		} else {
			err := s.db.Write(op.table, op.key, op.value)
			if err != nil {
				return fmt.Errorf("error writing new state for key %s, err: %v", op.key, err)
			}
		}
	}

	return nil
}

func (s *ChainStorer) GetBlockByHeight(blockHeight uint64) (*block.Block, error) {
	data, ok := s.db.Read(blockByHeight, fmt.Sprintf("block_%d", blockHeight))
	if !ok {
		return nil, fmt.Errorf("error reading block at height %d", blockHeight)
	}

	blockData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error marshalling block data at height %d: %v", blockHeight, err)
	}

	var blk block.Block
	err = json.Unmarshal(blockData, &blk)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling block data at height %d: %v", blockHeight, err)
	}

	return &blk, nil
}

func (s *ChainStorer) GetBlockByHash(blockHash common.Hash) (*block.Block, error) {
	data, ok := s.db.Read(blockByHash, fmt.Sprintf("block_%d", blockHash))
	if !ok {
		return nil, fmt.Errorf("error reading block at height %d", blockHash)
	}

	blockData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error marshalling block data with hash  %s: %v", blockHash.String(), err)
	}

	var blk block.Block
	err = json.Unmarshal(blockData, &blk)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling block data with hash %s: %v", blockHash.String(), err)
	}

	return &blk, nil
}

func (s *ChainStorer) GetCurrentBlock() (*block.Block, error) {
	data, ok := s.db.Read(latestBlock, "currentBlock")
	if !ok {
		return nil, fmt.Errorf("error reading current block")
	}

	blockData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error marshalling current block: %v", err)
	}

	var blk block.Block
	err = json.Unmarshal(blockData, &blk)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling current block: %v", err)
	}

	return &blk, nil
}
