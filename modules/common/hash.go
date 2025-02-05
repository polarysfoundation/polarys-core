package common

import (
	"encoding/hex"
	"math/big"
)

const (
	hashLen = 32
)

type Hash [hashLen]byte

func (h *Hash) SetBytes(b []byte) {
	if len(b) > len(h) {
		b = b[len(b)-hashLen:]
	}

	copy(h[hashLen-len(b):], b)
}

func (h Hash) ToBytes() []byte {
	return h[:]
}

func (h Hash) ToString() string {
	return string(h.flexhex())
}

func (h Hash) ToHex() string {
	return string(h.hex())
}

func (h Hash) ToBigInt() *big.Int {
	return new(big.Int).SetBytes(h.ToBytes())
}

func (h Hash) flexhex() []byte {
	buf := make([]byte, len(addressPrefix)+len(h)*2)
	copy(buf[:len(addressPrefix)], []byte(addressPrefix))
	encode(buf[len(addressPrefix):], h[:])
	return buf
}

func (h Hash) IsEmpty() bool {
	return h == Hash{}
}

func (h Hash) hex() []byte {
	buf := make([]byte, len(h)*2+2)
	copy(buf[:2], []byte("0x"))
	encode(buf[2:], h[:])
	return buf
}

func BytesToHash(b []byte) Hash {
	var h Hash
	h.SetBytes(b)
	return h
}

func StringToHash(s string) Hash {
	return BytesToHash([]byte(s))
}

func BigIntToHash(n *big.Int) Hash {
	return BytesToHash(n.Bytes())
}

func HexToHash(s string) Hash {
	if has0xPrefix(s) {
		s = s[2:]
	}

	b, _ := hex.DecodeString(s)
	return BytesToHash(b)
}

func has0xPrefix(s string) bool {
	return len(s) >= 2 && s[0] == '0' && s[1] == 'x'
}

func FlexHexToHash(f string) Hash {
	return BytesToHash([]byte(f))
}
