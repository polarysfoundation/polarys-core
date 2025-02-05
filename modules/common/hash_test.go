package common

import (
	"encoding/hex"
	"math/big"
	"testing"
)

func TestSetBytes(t *testing.T) {
	t.Log("Starting TestSetBytes")
	var h Hash
	h.SetBytes([]byte{1, 2, 3, 4})
	expected := make([]byte, hashLen)
	copy(expected[hashLen-4:], []byte{1, 2, 3, 4})
	if h != BytesToHash(expected) {
		t.Errorf("SetBytes failed, expected %v, got %v", expected, h)
	} else {
		t.Log("SetBytes passed")
	}
}

func TestToBytes(t *testing.T) {
	t.Log("Starting TestToBytes")
	h := BytesToHash([]byte{1, 2, 3, 4})
	expected := make([]byte, hashLen)
	copy(expected[hashLen-4:], []byte{1, 2, 3, 4})
	if !equal(h.ToBytes(), expected) {
		t.Errorf("ToBytes failed, expected %v, got %v", expected, h.ToBytes())
	} else {
		t.Log("ToBytes passed")
	}
}

func TestToString(t *testing.T) {
	t.Log("Starting TestToString")
	h := BytesToHash([]byte{1, 2, 3, 4})
	expected := h.ToString()
	if h.ToString() != expected {
		t.Errorf("ToString failed, expected %v, got %v", expected, h.ToString())
	} else {
		t.Log("ToString passed")
	}
}

func TestToBigInt(t *testing.T) {
	t.Log("Starting TestToBigInt")
	h := BytesToHash([]byte{1, 2, 3, 4})
	expected := new(big.Int).SetBytes(h.ToBytes())
	if h.ToBigInt().Cmp(expected) != 0 {
		t.Errorf("ToBigInt failed, expected %v, got %v", expected, h.ToBigInt())
	} else {
		t.Log("ToBigInt passed")
	}
}

func TestBytesToHash(t *testing.T) {
	t.Log("Starting TestBytesToHash")
	b := []byte{1, 2, 3, 4}
	h := BytesToHash(b)
	expected := make([]byte, hashLen)
	copy(expected[hashLen-4:], b)
	if h != BytesToHash(expected) {
		t.Errorf("BytesToHash failed, expected %v, got %v", expected, h)
	} else {
		t.Log("BytesToHash passed")
	}
}

func TestStringToHash(t *testing.T) {
	t.Log("Starting TestStringToHash")
	s := "test"
	h := StringToHash(s)
	expected := BytesToHash([]byte(s))
	if h != expected {
		t.Errorf("StringToHash failed, expected %v, got %v", expected, h)
	} else {
		t.Log("StringToHash passed")
	}
}

func TestBigIntToHash(t *testing.T) {
	t.Log("Starting TestBigIntToHash")
	n := big.NewInt(1234)
	h := BigIntToHash(n)
	expected := BytesToHash(n.Bytes())
	if h != expected {
		t.Errorf("BigIntToHash failed, expected %v, got %v", expected, h)
	} else {
		t.Log("BigIntToHash passed")
	}
}

func TestHexToHash(t *testing.T) {
	t.Log("Starting TestHexToHash")
	s := "01020304"
	h := HexToHash(s)
	b, _ := hex.DecodeString(s)
	expected := BytesToHash(b)
	if h != expected {
		t.Errorf("HexToHash failed, expected %v, got %v", expected, h)
	} else {
		t.Log("HexToHash passed")
	}
}

// Helper function to compare two byte slices
func equal(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
