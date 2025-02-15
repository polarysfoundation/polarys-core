package transaction

import (
	"log"
	"math/big"
	"testing"

	"github.com/polarysfoundation/polarys-core/modules/common"
)

func TestNewTx(t *testing.T) {
	log.Println("Starting TestNewTx")
	from := common.BytesToAddress([]byte{1, 2, 3, 4})
	to := common.BytesToAddress([]byte{5, 6, 7, 8})
	value := big.NewInt(1000)
	nonce := uint64(1)
	data := []byte("transaction data")

	tx := NewLegacyTx(from, to, value, nonce, data, []byte(""))

	log.Printf("Created transaction: %+v", tx)

	if tx.From != from {
		t.Errorf("expected from address %v, got %v", from, tx.From)
	}
	if tx.To != to {
		t.Errorf("expected to address %v, got %v", to, tx.To)
	}
	if tx.Value.Cmp(value) != 0 {
		t.Errorf("expected value %v, got %v", value, tx.Value)
	}
	if tx.Nonce != nonce {
		t.Errorf("expected nonce %v, got %v", nonce, tx.Nonce)
	}
	if string(tx.Data) != string(data) {
		t.Errorf("expected data %v, got %v", data, tx.Data)
	}
	log.Println("Finished TestNewTx")
}

func TestCopyTx(t *testing.T) {
	log.Println("Starting TestCopyTx")
	from := common.BytesToAddress([]byte{1, 2, 3, 4})
	to := common.BytesToAddress([]byte{5, 6, 7, 8})
	value := big.NewInt(1000)
	nonce := uint64(1)
	data := []byte("transaction data")

	tx := NewLegacyTx(from, to, value, nonce, data, []byte(""))

	txCopy := copyTx(tx)

	log.Printf("Original transaction: %+v", tx)
	log.Printf("Copied transaction: %+v", txCopy)

	if txCopy.From != from {
		t.Errorf("expected from address %v, got %v", from, txCopy.From)
	}
	if txCopy.To != to {
		t.Errorf("expected to address %v, got %v", to, txCopy.To)
	}
	if txCopy.Value.Cmp(value) != 0 {
		t.Errorf("expected value %v, got %v", value, txCopy.Value)
	}
	if txCopy.Nonce != nonce {
		t.Errorf("expected nonce %v, got %v", nonce, txCopy.Nonce)
	}
	if string(txCopy.Data) != string(data) {
		t.Errorf("expected data %v, got %v", data, txCopy.Data)
	}
	log.Println("Finished TestCopyTx")
}

func TestMarshalUnmarshalTx(t *testing.T) {
	log.Println("Starting TestMarshalUnmarshalTx")
	from := common.BytesToAddress([]byte{1, 2, 3, 4})
	to := common.BytesToAddress([]byte{5, 6, 7, 8})
	value := big.NewInt(1000)
	nonce := uint64(1)
	data := []byte("transaction data")

	tx := NewLegacyTx(from, to, value, nonce, data, []byte(""))
	marshaledTx, err := tx.marshal()
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	log.Printf("Marshaled transaction: %x", marshaledTx)

	var unmarshaledTx Tx
	err = unmarshaledTx.unmarshal(marshaledTx)
	if err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	log.Printf("Unmarshaled transaction: %+v", unmarshaledTx)

	if unmarshaledTx.From != from {
		t.Errorf("expected from address %v, got %v", from, unmarshaledTx.From)
	}
	if unmarshaledTx.To != to {
		t.Errorf("expected to address %v, got %v", to, unmarshaledTx.To)
	}
	if unmarshaledTx.Value.Cmp(value) != 0 {
		t.Errorf("expected value %v, got %v", value, unmarshaledTx.Value)
	}
	if unmarshaledTx.Nonce != nonce {
		t.Errorf("expected nonce %v, got %v", nonce, unmarshaledTx.Nonce)
	}
	if string(unmarshaledTx.Data) != string(data) {
		t.Errorf("expected data %v, got %v", data, unmarshaledTx.Data)
	}
	log.Println("Finished TestMarshalUnmarshalTx")
}
