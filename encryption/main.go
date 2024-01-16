package encryption

import (
	"encoding/base64"
)

func EncryptBase64(plainText string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(plainText))
	return encoded
}

func DecryptBase64(encodedText string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(encodedText)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}
