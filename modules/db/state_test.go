package db_test

import (
	"log"
	"math/big"
	"testing"

	"github.com/polarysfoundation/polarys-core/modules/common"
	"github.com/polarysfoundation/polarys-core/modules/db"
	"github.com/stretchr/testify/assert"
)

func TestNewState(t *testing.T) {
	log.Println("Starting TestNewState")

	account := common.Address{0x01}
	balance := big.NewInt(1000)
	nonces := uint64(1)
	latestUpdate := uint64(1234567890)

	state := db.NewState(account, balance, nonces, latestUpdate)

	log.Println("State created:", state)

	assert.Equal(t, account, state.Account())
	assert.Equal(t, balance, state.Balance())
	assert.Equal(t, nonces, state.Nonces())
	assert.Equal(t, latestUpdate, state.LatestUpdate())

	log.Println("TestNewState completed successfully")
}

func TestWriteState(t *testing.T) {
	log.Println("Starting TestWriteState")

	account := common.Address{0x01}
	balance := big.NewInt(1000)
	nonces := uint64(1)
	latestUpdate := uint64(1234567890)
	state := db.NewState(account, balance, nonces, latestUpdate)

	log.Println("State created:", state)
	log.Println("State created for address:", state.Account().String())

	chainStorer := db.NewChainStorer()

	err := chainStorer.WriteState(state)
	assert.NoError(t, err)

	log.Println("State written to chainStorer")

	log.Println("TestWriteState completed successfully")
}

func TestReadState(t *testing.T) {
	log.Println("Starting TestReadState")

	account := common.Address{0x01}
	balance := big.NewInt(1000)
	nonces := uint64(1)
	latestUpdate := uint64(1234567890)

	chainStorer := db.NewChainStorer()

	state, err := chainStorer.ReadState(account)
	assert.NoError(t, err)

	log.Println("State read from chainStorer:", state)

	assert.Equal(t, account, state.Account())
	assert.Equal(t, balance, state.Balance())
	assert.Equal(t, nonces, state.Nonces())
	assert.Equal(t, latestUpdate, state.LatestUpdate())

	log.Println("TestReadState completed successfully")
}
