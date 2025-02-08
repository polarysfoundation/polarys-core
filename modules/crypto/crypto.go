package crypto

import (
	"math/big"

	pm256 "github.com/polarysfoundation/pm-256"
	"github.com/polarysfoundation/polarys-core/modules/common"
)

func Pm256(b []byte) []byte {
	buf := make([]byte, 32)
	h := pm256.New256()
	h.Write(b)
	h.Sum(buf[:0])

	return buf
}

func CreateAddress(a common.Address, n uint64, h common.Hash) common.Address {
	nonce := new(big.Int)
	nonce.SetUint64(n)

	data := make([]byte, len(nonce.Bytes())+len(a.Bytes())+len(h.Bytes()))
	data[0] = 0xff
	copy(data[1:], nonce.Bytes())
	copy(data[1+len(nonce.Bytes()):], a.Bytes())
	copy(data[1+len(nonce.Bytes())+len(a.Bytes()):], h.Bytes())

	return common.BytesToAddress(Pm256(data))
}

func CreatePoolKey(s common.Address, n uint64, nodeKey []byte) common.Key {
	nonce := new(big.Int)
	nonce.SetUint64(n)

	data := make([]byte, len(nonce.Bytes())+len(s.Bytes())+len(nodeKey))
	data[0] = 0xff
	copy(data[1:], nonce.Bytes())
	copy(data[1+len(nonce.Bytes()):], s.Bytes())
	copy(data[1+len(nonce.Bytes())+len(s.Bytes()):], nodeKey)

	return common.BytesToKey(Pm256(data))
}
