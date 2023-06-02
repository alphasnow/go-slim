package xencode

import "encoding/hex"

func HexEncode(plainText []byte) string {
	return hex.EncodeToString(plainText)
}

func HexDecode(cipherTextHex string) ([]byte, error) {
	return hex.DecodeString(cipherTextHex)
}
