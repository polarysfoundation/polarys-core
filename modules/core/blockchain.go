package core

import (
	"sync"

	"github.com/polarysfoundation/polarys-core/modules/core/block"
	"github.com/polarysfoundation/polarys-core/modules/db"
)

const (
	maxCachedBlock = 1000
)

type Blockchain struct {
	cachedBlocks []*block.Block
	chainDB      *db.ChainStorer
	chain        *Chain

	mutex *sync.RWMutex
}
