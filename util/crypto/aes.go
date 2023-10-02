package crypto

import (
	"crypto/aes"
	"encoding/hex"
	"errors"

	"github.com/spf13/viper"
)

var (
	errorEmptyEncryptionKey = errors.New("empty encryption key")
	errorInitiatingCipher   = errors.New("error initiating cipher")
)

// Encrypt the given data and produce the encrypted data as string
// will return error if something goes unexpected
func EncryptAES(data string) (string, error) {
	key := viper.GetString("AES_ENC_KEY")
	if key == "" {
		return "", errorEmptyEncryptionKey
	}

	cipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", errorInitiatingCipher
	}

	out := make([]byte, len(data))
	cipher.Encrypt(out, []byte(data))

	return hex.EncodeToString(out), nil
}

// Decrypt the given encrypted and product the plain version as string
// will return error if something goes unexpected
func DecryptAES(encrypted string) (string, error) {
	key := viper.GetString("AES_ENC_KEY")
	if key == "" {
		return "", errorEmptyEncryptionKey
	}

	cipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", errorInitiatingCipher
	}

	out := make([]byte, len(encrypted))
	cipher.Decrypt(out, []byte(encrypted))

	return string(out), nil
}
