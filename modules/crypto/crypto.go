package crypto

import (
	"log"
	"math/big"

	pec256 "github.com/polarysfoundation/pec-256"
	pm256 "github.com/polarysfoundation/pm-256"
	"github.com/polarysfoundation/polarys-core/modules/common"
)

var c = pec256.PEC256()

func GenerateKey() (pec256.PrivKey, pec256.PubKey) {
	priv, pub, _, err := c.GenerateKeyPair()
	if err != nil {
		log.Printf("error generating keys: %v", err)
		panic("error creating new keypair")
	}

	return priv, pub
}

func GenerateSharedKey(priv pec256.PrivKey) pec256.SharedKey {
	return c.SharedKey(priv)
}

func GeneratePubkey(priv pec256.PrivKey) pec256.PubKey {
	pub, _ := c.GetPubKey(priv)
	return pub
}

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
