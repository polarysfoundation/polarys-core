package keystore

import (
	"fmt"
	"sync"

	"github.com/polarysfoundation/polarys-core/modules/common"
	"github.com/polarysfoundation/polarys-core/modules/core/transaction"
)

type Wallet struct {
	accounts map[common.Address]bool
	k        map[common.Address]*Keypair
	mutex    sync.RWMutex
}

func InitWallet() *Wallet {
	accounts := GetLocalAccounts()
	if len(accounts) == 0 {
		return nil
	}

	w := &Wallet{
		accounts: make(map[common.Address]bool),
		k:        make(map[common.Address]*Keypair),
	}

	for _, acc := range accounts {
		w.accounts[acc] = false
	}

	return w
}

func (w *Wallet) SignTX(a common.Address, tx *transaction.Transaction) (*transaction.Transaction, error) {

	if w.IsLocked(a) {
		return nil, fmt.Errorf("address %s is locked", a.String())
	}

	k, ok := w.k[a]
	if !ok {
		return nil, fmt.Errorf("address %s is locked or does not exist", a.String())
	}

	return k.signTX(tx)
}

func (w *Wallet) Refresh() {
	w.scan()

	if len(w.accounts) == 0 {
		return
	}

	acc := w.Accounts()
	for _, account := range acc {
		if !w.IsLocked(account) {
			k := w.k[account]
			if k.expired() {
				k.lock()
				w.accounts[account] = false
			}
		}
	}

}

func (w *Wallet) scan() {
	accounts := GetLocalAccounts()
	if len(accounts) == 0 && len(w.accounts) > 0 {
		for a := range w.accounts {
			w.removeAddress(a)
			w.removeKeypair(a)
		}
		return
	}

	for _, a := range accounts {
		if !w.exist(a) {
			w.addAddress(a)
		}
	}

}

func (w *Wallet) exist(a common.Address) bool {
	_, exist := w.accounts[a]
	return exist
}

func (w *Wallet) addAddress(a common.Address) {
	if _, exist := w.accounts[a]; !exist {
		return
	}

	w.accounts[a] = false
}

func (w *Wallet) removeAddress(a common.Address) {
	if _, exist := w.accounts[a]; !exist {
		return
	}

	delete(w.accounts, a)
}

func (w *Wallet) removeKeypair(a common.Address) {
	if _, exist := w.k[a]; !exist {
		return
	}

	delete(w.k, a)
}

func (w *Wallet) Accounts() []common.Address {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	acc := make([]common.Address, 0)
	for account := range w.accounts {
		acc = append(acc, account)
	}

	return acc

}

func (w *Wallet) IsLocked(a common.Address) bool {
	unlocked, exist := w.accounts[a]
	if !unlocked {
		return true
	} else if !exist {
		return true
	}

	return false
}

func (w *Wallet) Unlock(a common.Address, passphrase []byte) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	unlocked, exist := w.accounts[a]
	if unlocked {
		return fmt.Errorf("account %s already unlocked", a.String())
	} else if !exist {
		return fmt.Errorf("account %s does not exist", a.String())
	}

	k, err := initKeypair(a, passphrase)
	if err != nil {
		return err
	}

	if a.String() != k.address().String() {
		return fmt.Errorf("address not match, expected %s, got %s", a.String(), k.address().String())
	}

	w.k[a] = k
	w.accounts[a] = true

	return nil
}
