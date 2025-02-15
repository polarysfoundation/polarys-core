package keystore

import (
	"log"
	"math/big"
	"testing"

	"github.com/polarysfoundation/polarys-core/modules/common"
	"github.com/polarysfoundation/polarys-core/modules/core/transaction"
)

func TestWalletInitialization(t *testing.T) {
	log.Println("Starting Wallet Initialization Test")
	w := InitWallet()
	if w == nil {
		t.Fatalf("Expected Wallet instance, got nil")
	}
	log.Printf("Wallet initialized with %d accounts\n", len(w.Accounts()))
}

func TestWalletScan(t *testing.T) {
	log.Println("Starting Wallet Scan Test")
	w := InitWallet()
	if w == nil {
		t.Fatalf("Wallet initialization failed")
	}
	w.Refresh()
	log.Printf("Wallet scanned, accounts: %d\n", len(w.Accounts()))
}

func TestWalletLocking(t *testing.T) {
	log.Println("Starting Wallet Locking Test")
	w := InitWallet()
	if w == nil {
		t.Fatalf("Wallet initialization failed")
	}

	address := common.FlexHexToAddress("1cx736de0b4edf141be4636d7094c84c9f2ab33138ff4a7e55085")
	if w.IsLocked(address) {
		log.Printf("Address %s is locked as expected\n", address)
	} else {
		t.Errorf("Expected address %s to be locked\n", address)
	}
}

func TestWalletUnlocking(t *testing.T) {
	log.Println("Starting Wallet Unlocking Test")
	w := InitWallet()
	if w == nil {
		t.Fatalf("Wallet initialization failed")
	}

	address := common.FlexHexToAddress("1cx736de0b4edf141be4636d7094c84c9f2ab33138ff4a7e55085")
	passphrase := []byte("testpassword")
	err := w.Unlock(address, passphrase)
	if err != nil {
		t.Errorf("Failed to unlock wallet: %s\n", err)
	}

	if w.IsLocked(address) {
		t.Errorf("Expected address %s to be unlocked\n", address)
	} else {
		log.Printf("Address %s unlocked successfully\n", address)
	}
}

func TestWalletSign(t *testing.T) {
	log.Println("Starting Wallet Unlocking Test")
	w := InitWallet()
	if w == nil {
		t.Fatalf("Wallet initialization failed")
	}

	address := common.FlexHexToAddress("1cx736de0b4edf141be4636d7094c84c9f2ab33138ff4a7e55085")
	passphrase := []byte("testpassword")
	err := w.Unlock(address, passphrase)
	if err != nil {
		t.Errorf("Failed to unlock wallet: %s\n", err)
	}

	if w.IsLocked(address) {
		t.Errorf("Expected address %s to be unlocked\n", address)
	} else {
		log.Printf("Address %s unlocked successfully\n", address)
	}

	tx := transaction.NewLegacyTx(address, common.FlexHexToAddress("1cx660a440474ebdfdf1c4df9af04ecb4e3a81e145c68a58bf833"), big.NewInt(1e10), 5, []byte{1, 2, 3}, []byte("hola"))
	transaction := transaction.NewTransaction(tx)

	nTx, err := w.SignTX(address, transaction)
	if err != nil {
		t.Errorf("Failed sign transaction %v", err)
	}

	nTx.Print()
}
