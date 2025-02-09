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

func (k Key) Bytes() []byte {
	return k[:]
}

func (k Key) String() string {
	return string(k.flexhex())
}

func (k Key) Hex() string {
	return string(k.hex())
}

func (k Key) flexhex() []byte {
	buf := make([]byte, len(addressPrefix)+len(k)*2)
	copy(buf[:len(addressPrefix)], []byte(addressPrefix))
	encode(buf[len(addressPrefix):], k[:])
	return buf
}

func (k Key) hex() []byte {
	buf := make([]byte, len(k)*2+2)
	copy(buf[:2], []byte("0x"))
	encode(buf[2:], k[:])
	return buf
}
