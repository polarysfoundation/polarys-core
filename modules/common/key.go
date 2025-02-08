package common

const (
	keyLen = 12
)

type Key [keyLen]byte

func BytesToKey(b []byte) Key {
	var k Key
	k.SetBytes(b)
	return k
}



func (k *Key) SetBytes(b []byte) {
	if len(b) > len(k) {
		b = b[len(b)-keyLen:]
	}

	copy(k[keyLen-len(b):], b)
}