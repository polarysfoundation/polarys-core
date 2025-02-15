package keystore

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/polarysfoundation/polarys-core/modules/common"
	"github.com/polarysfoundation/polarys-core/modules/config"
	"github.com/polarysfoundation/polarys-core/modules/crypto"
)

// TestNewAndLoadKeypair verifies that a keypair can be generated,
// saved to a keystore file, and then loaded back using the correct passphrase.
// It also confirms that using an incorrect passphrase fails to load the keypair.
func TestNewAndLoadKeypair(t *testing.T) {
	// Create a temporary directory for the keystore files.
	tempDir := config.GetKeystorePath()
	t.Logf("Using temporary keystore directory: %s", tempDir)

	// Define a passphrase for encryption/decryption.
	passphrase := []byte("testpassword")

	// Generate a new keypair.
	t.Log("Creating a new keypair with the correct passphrase")
	keypair, err := NewKeypair(passphrase)
	if err != nil {
		t.Fatalf("Failed to create new keypair: %v", err)
	}

	// Compute the expected keystore file name using the public key's address.
	addr := crypto.PubKeyToAddress(keypair.pub)
	keystoreFile := filepath.Join(tempDir, addr.String()+".json")
	t.Logf("New keypair created with address: %s", addr.String())

	// Check that the keystore file was created.
	if _, err := os.Stat(keystoreFile); os.IsNotExist(err) {
		t.Errorf("Keystore file not found at %s", keystoreFile)
	} else {
		t.Logf("Keystore file exists: %s", keystoreFile)
	}

	// Load keypairs using the correct passphrase.
	t.Log("Loading keypairs with the correct passphrase")
	keypairs, err := LoadKeypairs(passphrase)
	if err != nil {
		t.Fatalf("Failed to load keypairs: %v", err)
	}

	if len(keypairs) == 0 {
		t.Error("No keypairs loaded; expected at least one")
	} else {
		t.Logf("Loaded %d keypair(s)", len(keypairs))
	}

	// Verify that the generated keypair is among the loaded ones.

	var temp *Keypair
	found := false
	for _, kp := range keypairs {
		loadedAddr := crypto.PubKeyToAddress(kp.pub)
		t.Logf("Loaded keypair with address: %s", loadedAddr.String())
		if loadedAddr.String() == addr.String() {
			temp = kp
			found = true
			break
		}
	}
	if !found {
		t.Error("The generated keypair was not found among the loaded keypairs")
	}

	t.Log("Trying reseting keypair ")

	t.Logf("Keypair priv key %v", temp.priv.String())

	temp.close()

	t.Logf("New keypair priv key %v", temp.priv.String())

	if !common.BytesToHash(temp.priv.Bytes()).IsEmpty() {
		t.Fatal("error reseting keypair")
	}

	// Attempt to load keypairs using a wrong passphrase.
	t.Log("Attempting to load keypairs with an incorrect passphrase")
	wrongPassphrase := []byte("wrongpassword")
	keypairsWrong, err := LoadKeypairs(wrongPassphrase)
	if err != nil {
		t.Fatalf("Error occurred while loading keypairs with a wrong passphrase: %v", err)
	}
	if len(keypairsWrong) > 0 {
		t.Errorf("Expected 0 keypairs with a wrong passphrase, but got %d", len(keypairsWrong))
	} else {
		t.Log("No keypairs loaded with the wrong passphrase as expected")
	}
}

func TestHasKeypair(t *testing.T) {
	t.Log("Testing HasKeypairs function")

	var keys int
	var ok bool
	if ok, keys = HasKeyspair(); !ok {
		t.Log("not keys found")
	}

	if keys > 0 {
		t.Logf("total of %d keypairs found", keys)
	}

}

func TestGetLocalKeypair(t *testing.T) {
	t.Log("Testing GetLocalKeypair function")

	addresses := GetLocalAccounts()

	if len(addresses) > 0 {
		t.Logf("total of %d keypairs found", len(addresses))

		for _, addr := range addresses {
			t.Logf("Local address: %s", addr.String())
		}
	}

}

func TestGetLocalKeypairByAddress(t *testing.T) {
	t.Log("Testing GetLocalKeypairByAddress function")

	address := common.FlexHexToAddress("1cx736de0b4edf141be4636d7094c84c9f2ab33138ff4a7e55085")

	keypair, err := GetKeypairByAddress(address, []byte("testpassword"))
	if err != nil {
		t.Errorf("error getting keypair for address %s", address.String())
	}

	addr := crypto.PubKeyToAddress(keypair.pub)

	if addr != address {
		t.Errorf("error address not match expected %s, got %s", address.String(), addr.String())
	} else {
		t.Logf("success keypair obtained, address %s", addr.String())
	}
}
