package common

import (
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

func (h Hash) Bytes() []byte {
	return h[:]
}

func (h Hash) String() string {
	return string(h.flexhex())
}

func (h Hash) Hex() string {
	return string(h.hex())
}

func (h Hash) BigInt() *big.Int {
	return new(big.Int).SetBytes(h.Bytes())
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

	b := decode(s)
	return BytesToHash(b)
}

func FlexHexToHash(f string) Hash {
	if has1cxPrefix(f) {
		f = f[3:]
	}
	b := decode(f)
	return BytesToHash(b)
}
