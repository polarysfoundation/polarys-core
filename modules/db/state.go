package db

import (
	"fmt"
	"log"
	"math/big"

	"github.com/polarysfoundation/polarys-core/modules/common"
)

const (
	chainState = "/chain/state/accounts/"
)

type State struct {
	account      common.Address
	balance      *big.Int
	nonces       uint64
	latestUpdate uint64
}

func NewState(account common.Address, balance *big.Int, nonces, latestUpdate uint64) *State {
	return &State{
		account:      account,
		balance:      balance,
		nonces:       nonces,
		latestUpdate: latestUpdate,
	}
}

func (s *State) Account() common.Address { return s.account }
func (s *State) Balance() *big.Int       { return s.balance }
func (s *State) Nonces() uint64          { return s.nonces }
func (s *State) LatestUpdate() uint64    { return s.latestUpdate }

type MState struct {
	Account      common.Address `json:"account"`
	Balance      *big.Int       `json:"balance"`
	Nonces       uint64         `json:"nonces"`
	LatestUpdate uint64         `json:"latestUpdate"`
}

func copyState(state *State) *MState {
	return &MState{
		Account:      state.account,
		Balance:      state.balance,
		Nonces:       state.nonces,
		LatestUpdate: state.latestUpdate,
	}
}

func (s *ChainStorer) WriteState(state *State) error {
	if !s.db.Exist(chainState) {
		err := s.db.Create(chainState, state.account.String(), copyState(state))
		if err != nil {
			return fmt.Errorf("error writing new state, err: %v", err)
		}
	} else {
		err := s.db.Write(chainState, state.account.String(), copyState(state))
		if err != nil {
			return fmt.Errorf("error writing new state, err: %v", err)
		}
	}

	return nil
}

func (s *ChainStorer) ReadState(account common.Address) (*State, error) {
	data, ok := s.db.Read(chainState, account.String())
	if !ok {
		return nil, fmt.Errorf("error reading state for address %s", account.String())
	}

	log.Println(data)

	dataMap, ok := data.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("error reading state for address %s", account.String())
	}

	nonces, ok := dataMap["nonces"].(float64)
	if !ok {
		return nil, fmt.Errorf("error reading state for address %s", account.String())
	}

	latestUpdate, ok := dataMap["latestUpdate"].(float64)
	if !ok {
		return nil, fmt.Errorf("error reading state for address %s", account.String())
	}

	state := &State{
		account:      common.BytesToAddress(common.ConvertInterfaceSliceToByteSlice(dataMap["account"].([]interface{}))),
		balance:      new(big.Int),
		nonces:       uint64(nonces),
		latestUpdate: uint64(latestUpdate),
	}

	if balance, ok := dataMap["balance"].(string); ok {
		state.balance.SetString(balance, 10)
	} else if balance, ok := dataMap["balance"].(float64); ok {
		state.balance.SetString(fmt.Sprintf("%.0f", balance), 10)
	} else {
		return nil, fmt.Errorf("error reading balance for address %s", account.String())
	}

	return state, nil
}
