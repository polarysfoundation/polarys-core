package common

import (
	"crypto/rand"
	"io"
	"math/big"

	pm256 "github.com/polarysfoundation/pm-256"
)

const (
	addressLen    = 25
	addressPrefix = "1cx"
	hextable      = "0123456789abcdef"
)

type Address [addressLen]byte

func BytesToAddress(b []byte) Address {
	var addr Address
	addr.SetBytes(b)
	return addr
}

func BigIntToAddress(n *big.Int) Address {
	return BytesToAddress(n.Bytes())
}

func StringToAddress(s string) Address {
	if len(s) > len(addressPrefix) && s[:len(addressPrefix)] == addressPrefix {
		s = s[len(addressPrefix):]
	}
	return BytesToAddress([]byte(s))
}

func HexToAddress(s string) Address {
	if len(s) > 1 && s[:2] == "0x" {
		s = s[2:]
	} else {
		return Address{}
	}
	return BytesToAddress([]byte(s))
}

func FlexHexToAddress(f string) Address {
	return BytesToAddress([]byte(f))
}

func (a Address) ToString() string {
	return string(a.flexhex())
}

func (a Address) ToBytes() []byte {
	return a[:]
}

func (a Address) ToHex() string {
	return string(a.hex())
}

func (a Address) ToBigInt() *big.Int {
	return new(big.Int).SetBytes(a.ToBytes())
}

func (a Address) Length() int {
	return len(a)
}

func (a Address) flexhex() []byte {
	buf := make([]byte, len(addressPrefix)+len(a)*2)
	copy(buf[:len(addressPrefix)], []byte(addressPrefix))
	encode(buf[len(addressPrefix):], a[:])
	return buf
}

func (a Address) hex() []byte {
	buf := make([]byte, len(a)*2+2)
	copy(buf[:2], []byte("0x"))
	encode(buf[2:], a[:])
	return buf
}

func encode(dst, src []byte) int {
	j := 0
	for _, v := range src {
		dst[j] = hextable[v>>4]
		dst[j+1] = hextable[v&0x0f]
		j += 2
	}
	return len(src) * 2
}


func (a *Address) SetBytes(b []byte) {
	if len(b) > len(a) {
		b = b[len(b)-addressLen:]
	}

	copy(a[addressLen-len(b):], b)
}

func GenerateAddress() Address {
	b := make([]byte, addressLen)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		panic(err)
	}

	h := pm256.Sum256(b)
	addr := BytesToAddress(h[:])

	return addr
}
