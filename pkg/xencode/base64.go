package xencode

import (
	"encoding/base64"
)

func Base64Encode(plainText []byte) string {
	return base64.StdEncoding.EncodeToString(plainText)
}

func Base64Decode(cipherTextBase64 string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(cipherTextBase64)
}
