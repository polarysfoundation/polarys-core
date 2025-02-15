package core

import (
	"sync"

	"github.com/polarysfoundation/polarys-core/modules/db"
)


type Blockchain struct {
	chainDB      *db.ChainStorer
	chain        *Chain

	mutex *sync.RWMutex
}
