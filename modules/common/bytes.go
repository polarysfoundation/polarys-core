package common

import "encoding/hex"

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