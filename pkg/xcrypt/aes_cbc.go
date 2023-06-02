package xcrypt

import (
	"encoding/base64"
	"encoding/json"
	"github.com/forgoer/openssl"
)

// AesCbc The length of the key can be 16/24/32 characters (128/192/256 bits)
type AesCbc struct {
	Secret []byte
	Iv     []byte
}

func NewAesCbc(config *Config) *AesCbc {
	return &AesCbc{
		Secret: config.GetSecret(),
		Iv:     config.GetSecret(),
	}
}

func (c *AesCbc) Encrypt(data []byte) ([]byte, error) {
	return openssl.AesCBCEncrypt(data, c.Secret, c.Iv, openssl.PKCS7_PADDING)
}

func (c *AesCbc) Decrypt(data []byte) ([]byte, error) {
	return openssl.AesCBCDecrypt(data, c.Secret, c.Iv, openssl.PKCS7_PADDING)
}

func (c *AesCbc) EncryptBase64(obj interface{}) (string, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	res, err := c.Encrypt(data)
	if err != nil {
		return "", err
	}
	resStr := base64.StdEncoding.EncodeToString(res)
	return resStr, nil
}

func (c *AesCbc) DecryptBase64(data string, obj interface{}) error {
	dataByte, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return err
	}
	dataObj, err := c.Decrypt(dataByte)
	if err != nil {
		return err
	}
	return json.Unmarshal(dataObj, obj)
}
