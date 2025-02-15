package common

import (
	"encoding/hex"
	"encoding/json"
)

func decode(s string) []byte {
	b, _ := hex.DecodeString(s)
	return b
}

func has0xPrefix(s string) bool {
	return len(s) >= 2 && s[0] == '0' && s[1] == 'x'
}

func has1cxPrefix(s string) bool {
	return len(s) >= len(addressPrefix) && s[0] == '1' && s[1] == 'c' && s[2] == 'x'
}

func ConvertInterfaceSliceToByteSlice(data []interface{}) []byte {
	byteSlice := make([]byte, len(data))
	for i, v := range data {
		byteSlice[i] = byte(v.(float64))
	}
	return byteSlice
}

func Equal(slice1, slice2 []byte) bool {
	// If lengths are different, slices are not equal
	if len(slice1) != len(slice2) {
		return false
	}

	// Compare each byte
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false // Found a mismatch
		}
	}

	// All bytes are equal
	return true
}

func Serialize(v ...interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func EncodeToString(data []byte) string {
	return string(flexEncoder(data))
}

func EncodeToHex(data []byte) string {
	return string(hexEncoder(data))
}

func hexEncoder(data []byte) []byte {
	// Create a buffer with enough space for "0x" and the hexadecimal representation of the data
	buf := make([]byte, len(data)*2+2)
	copy(buf[:2], []byte("0x"))
	// Encode the data into hexadecimal and write it to the buffer starting at index 2
	hex.Encode(buf[2:], data)

	return buf
}

func flexEncoder(data []byte) []byte {
	// Create a buffer with enough space for "0x" and the hexadecimal representation of the data
	buf := make([]byte, len(data)*len(addressPrefix)+2)
	copy(buf[:len(addressPrefix)], []byte(addressPrefix))
	// Encode the data into hexadecimal and write it to the buffer starting at index 2
	hex.Encode(buf[len(addressPrefix):], data)

	return buf
}
