package common

import (
	"log"
	"math/big"
	"testing"
)

func TestBytesToAddress(t *testing.T) {
	t.Log("Starting TestBytesToAddress")
	b := []byte("testaddress")
	addr := BytesToAddress(b)
	t.Logf("Converted bytes to address: %v", addr)
	if string(addr[len(addr)-len(b):]) != "testaddress" {
		t.Errorf("Expected 'testaddress', got '%s'", string(addr[:len(b)]))
	}
}

func TestBigIntToAddress(t *testing.T) {
	t.Log("Starting TestBigIntToAddress")
	n := big.NewInt(123456789)
	addr := BigIntToAddress(n)
	t.Logf("Converted big.Int to address: %v", addr)
	expected := BytesToAddress(n.Bytes())
	if addr != expected {
		t.Errorf("Expected '%v', got '%v'", expected, addr)
	}
}

func TestStringToAddress(t *testing.T) {
	t.Log("Starting TestStringToAddress")
	s := "1cxtestaddress"
	addr := StringToAddress(s)
	t.Logf("Converted string to address: %v", addr)
	expected := BytesToAddress([]byte("testaddress"))
	if addr != expected {
		t.Errorf("Expected '%v', got '%v'", expected, addr)
	}
}

func TestHexToAddress(t *testing.T) {
	t.Log("Starting TestHexToAddress")
	s := "0x7465737461646472657373"
	addr := HexToAddress(s)
	t.Logf("Converted hex to address: %v", addr)
	expected := BytesToAddress(addr.ToBytes())
	if addr != expected {
		t.Errorf("Expected '%v', got '%v'", expected, addr)
	}
}

func TestAddressToString(t *testing.T) {
	t.Log("Starting TestAddressToString")
	b := GenerateAddress()
	addr := BytesToAddress(b.ToBytes())
	t.Logf("Converted address to string: %s", addr.ToString())
	expected := b.ToString()
	addrStr := addr.ToString()
	log.Print(len(addrStr[len(addrStr)-len(expected):]))
	log.Print(len(expected))
	log.Printf("Bytes of result: %v", []byte(addrStr))
	log.Printf("Bytes of expected: %v", []byte(expected))
	if addrStr != expected {
		t.Errorf("Expected '%s', got '%s'", expected, addr.ToString())
	}
}

func TestAddressToBytes(t *testing.T) {
	t.Log("Starting TestAddressToBytes")
	b := GenerateAddress()
	addr := BytesToAddress(b.ToBytes())
	t.Logf("Converted address to bytes: %v", addr.ToBytes())
	if addr.ToString() != b.ToString() {
		t.Errorf("Expected 'testaddress', got '%s'", string(addr.ToBytes()[:len(b)]))
	}
}

func TestAddressToHex(t *testing.T) {
	t.Log("Starting TestAddressToHex")
	b := GenerateAddress()
	addr := BytesToAddress(b.ToBytes())
	t.Logf("Converted address to hex: %s", addr.ToHex())
	expected := b.ToHex()
	if addr.ToHex() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, addr.ToHex())
	}
}

func TestAddressToBigInt(t *testing.T) {
	t.Log("Starting TestAddressToBigInt")
	b := []byte("testaddress")
	addr := BytesToAddress(b)
	t.Logf("Converted address to big.Int: %v", addr.ToBigInt())
	expected := new(big.Int).SetBytes(b)
	if addr.ToBigInt().Cmp(expected) != 0 {
		t.Errorf("Expected '%v', got '%v'", expected, addr.ToBigInt())
	}
}

func TestSetBytesAddress(t *testing.T) {
	t.Log("Starting TestSetBytesAddress")
	var addr Address
	b := GenerateAddress()
	addr.SetBytes(b.ToBytes())
	t.Logf("Set bytes to address: %v", addr)
	if addr.ToString() != b.ToString() {
		t.Errorf("Expected 'testaddress', got '%s'", string(addr[:len(b)]))
	}
}
