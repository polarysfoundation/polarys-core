package core

import "github.com/polarysfoundation/polarys-core/modules/core/block"

type Chain struct {
	currentBlock         *block.Block
	snapshot             *block.Block
	latestConsensusBlock *block.Block
}

func (ch *Chain) CurrentBlock() *block.Block { return ch.currentBlock }
func (ch *Chain) SnapBlock() *block.Block    { return ch.snapshot }
func (ch *Chain) SafeBlock() *block.Block    { return ch.latestConsensusBlock }

func (ch *Chain) UpdateCurrentBlock(newBlock *block.Block) {
	ch.currentBlock = newBlock
}

func (ch *Chain) UpdateSnapBlock(newBlock *block.Block) {
	ch.snapshot = newBlock
}

func (ch *Chain) UpdateSafeBlock(newBlock *block.Block) {
	ch.latestConsensusBlock = newBlock
}
