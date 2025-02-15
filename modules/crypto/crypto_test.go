package crypto

import (
	"encoding/hex"
	"log"
	"testing"

	pec256 "github.com/polarysfoundation/pec-256"
	"github.com/polarysfoundation/polarys-core/modules/common"
	"github.com/stretchr/testify/assert"
)

func TestGenerateKey(t *testing.T) {
	log.Println("Starting TestGenerateKey")
	priv, pub := GenerateKey()
	log.Printf("Generated private key: %v", priv)
	log.Printf("Generated public key: %v", pub)
	assert.NotNil(t, priv, "Private key should not be nil")
	assert.NotNil(t, pub, "Public key should not be nil")
	log.Println("Finished TestGenerateKey")
}

func TestGenerateSharedKey(t *testing.T) {
	log.Println("Starting TestGenerateSharedKey")
	priv, _ := GenerateKey()
	sharedKey := GenerateSharedKey(priv)
	log.Printf("Generated shared key: %v", sharedKey)
	assert.NotNil(t, sharedKey, "Shared key should not be nil")
	log.Println("Finished TestGenerateSharedKey")
}

func TestGeneratePubkey(t *testing.T) {
	log.Println("Starting TestGeneratePubkey")
	priv, _ := GenerateKey()
	pub := GeneratePubkey(priv)
	log.Printf("Generated public key: %v", pub)
	assert.NotNil(t, pub, "Public key should not be nil")
	log.Println("Finished TestGeneratePubkey")
}

func TestPm256(t *testing.T) {
	log.Println("Starting TestPm256")
	data := []byte("test data")
	hash := Pm256(data)
	log.Printf("Generated hash: %v", hash)
	assert.NotNil(t, hash, "Hash should not be nil")
	assert.Equal(t, 32, len(hash), "Hash length should be 32 bytes")
	log.Println("Finished TestPm256")
}

func TestCreateAddress(t *testing.T) {
	log.Println("Starting TestCreateAddress")
	addr := common.BytesToAddress([]byte{0x01, 0x02, 0x03, 0x04})
	nonce := uint64(1)
	hash := common.BytesToHash([]byte{0x05, 0x06, 0x07, 0x08})
	newAddr := CreateAddress(addr, nonce, hash)
	log.Printf("Generated new address: %v", newAddr)
	assert.NotNil(t, newAddr, "New address should not be nil")
	log.Println("Finished TestCreateAddress")
}

func TestCreatePoolKey(t *testing.T) {
	log.Println("Starting TestCreatePoolKey")
	addr := common.BytesToAddress([]byte{0x01, 0x02, 0x03, 0x04})
	nonce := uint64(1)
	nodeKey := []byte{0x05, 0x06, 0x07, 0x08}
	poolKey := CreatePoolKey(addr, nonce, nodeKey)
	log.Printf("Generated pool key: %v", poolKey)
	assert.NotNil(t, poolKey, "Pool key should not be nil")
	log.Println("Finished TestCreatePoolKey")
}

func TestGetPubKeyFromPrivKey(t *testing.T) {
	log.Println("Starting TestGetPubKeyFromPrivKey")
	privKey := "623ed73a4ecd35c58b93e2031bca91af02edea968ab7fded4dde68fd7f789b45"
	privKeyBytes, err := hex.DecodeString(privKey)
	assert.Nil(t, err, "Error decoding private key hex")
	pubKey := GeneratePubkey(pec256.BytesToPrivKey(privKeyBytes))
	log.Printf("Generated public key from given private key: %v", pubKey)
	assert.NotNil(t, pubKey, "Public key should not be nil")
	log.Println("Finished TestGetPubKeyFromPrivKey")
}

func TestGenerateSharedKeyFromPrivKey(t *testing.T) {
	log.Println("Starting TestGenerateSharedKeyFromPrivKey")
	privKey := "623ed73a4ecd35c58b93e2031bca91af02edea968ab7fded4dde68fd7f789b45"
	privKeyBytes, err := hex.DecodeString(privKey)
	assert.Nil(t, err, "Error decoding private key hex")
	sharedKey := GenerateSharedKey(pec256.BytesToPrivKey(privKeyBytes))
	log.Printf("Generated shared key from given private key: %v", sharedKey)
	assert.NotNil(t, sharedKey, "Shared key should not be nil")
	log.Println("Finished TestGenerateSharedKeyFromPrivKey")
}
func TestVerify(t *testing.T) {

	log.Println("Starting TestVerify")
	for i := 0; i < 1000; i++ {
		priv, pub := GenerateKey()
		data := Pm256(priv.Bytes())
		dataBytes := common.BytesToHash(data)
		log.Printf("hash to sign: %s", dataBytes.Hex())
		r, s, err := Sign(dataBytes, priv)
		assert.Nil(t, err, "Error signing data")
		valid, err := Verify(dataBytes, r, s, pub)
		assert.Nil(t, err, "Error verifying signature")
		assert.True(t, valid, "Signature should be valid")
	}
	log.Println("Finished TestVerify")
}
