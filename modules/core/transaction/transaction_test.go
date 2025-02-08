package transaction

import (
	"math/big"
	"testing"

	"github.com/polarysfoundation/polarys-core/modules/common"
)

func TestCalculateHash(t *testing.T) {
	tx := &Tx{
		From:  common.BytesToAddress([]byte{1, 2, 3, 4}),
		To:    common.BytesToAddress([]byte{5, 6, 7, 8}),
		Value: big.NewInt(1000),
		Nonce: 1,
		Data:  []byte("transaction data"),
	}

	t.Logf("Creating transaction: From=%v, To=%v, Value=%v, Nonce=%v, Data=%s", tx.From, tx.To, tx.Value, tx.Nonce, tx.Data)

	transaction := NewTransaction(tx)
	transaction.calculateHash()

	t.Logf("Computed hash: %v", transaction.hash.Hex())

	if transaction.hash.IsEmpty() {
		t.Errorf("calculateHash failed, hash is empty")
	}

	expectedHash := common.HexToHash("0xe49bedb384f288c746895019ad8af44598ada9ecef31d76bc0fe8e2b1de331be") // Replace with the actual expected hash value
	t.Logf("Expected hash: %v", expectedHash.Hex())

	if transaction.hash != expectedHash {
		t.Errorf("calculateHash failed, expected %v, got %v", expectedHash, transaction.hash)
	}
}

func TestCalculateHashWithExistingHash(t *testing.T) {
	tx := &Tx{
		From:  common.BytesToAddress([]byte{1, 2, 3, 4}),
		To:    common.BytesToAddress([]byte{5, 6, 7, 8}),
		Value: big.NewInt(1000),
		Nonce: 1,
		Data:  []byte("transaction data"),
	}

	t.Logf("Creating transaction: From=%v, To=%v, Value=%v, Nonce=%v, Data=%s", tx.From, tx.To, tx.Value, tx.Nonce, tx.Data)

	transaction := NewTransaction(tx)
	initialHash := common.BytesToHash([]byte{13, 14, 15, 16})
	transaction.hash = initialHash

	t.Logf("Initial hash set: %v", initialHash.Hex())

	transaction.calculateHash()

	t.Logf("Computed hash: %v", transaction.hash.Hex())

	if transaction.hash == initialHash {
		t.Errorf("calculateHash failed, expected hash to be recomputed, got %v", transaction.hash)
	}
}

func TestCalculateHashForMultipleTransactions(t *testing.T) {
	for i := 0; i < 100; i++ {
		tx := &Tx{
			From:  common.BytesToAddress([]byte{1, 2, 3, 4}),
			To:    common.BytesToAddress([]byte{5, 6, 7, 8}),
			Value: big.NewInt(int64(i * 1000)),
			Nonce: uint64(i + 1),
			Data:  []byte("transaction data"),
		}

		t.Logf("Creating transaction %d: From=%v, To=%v, Value=%v, Nonce=%v, Data=%s", i+1, tx.From, tx.To, tx.Value, tx.Nonce, tx.Data)

		transaction := NewTransaction(tx)
		transaction.calculateHash()

		t.Logf("Transaction %d hash: %s", i+1, transaction.Hash().Hex())

		if transaction.hash.IsEmpty() {
			t.Errorf("calculateHash failed at transaction %d, hash is empty", i+1)
		}
	}
}
