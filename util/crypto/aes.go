package crypto

import (
	"crypto/aes"
	"crypto/cipher"
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

	aes, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", errorInitiatingCipher
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		return "", err
	}

	nonceHex := viper.GetString("AES_NONCE_KEY")
	nonce, err := hex.DecodeString(nonceHex)
	if err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nil, nonce, []byte(data), nil)
	return hex.EncodeToString(ciphertext), nil
}

// Decrypt the given encrypted and product the plain version as string
// will return error if something goes unexpected
func DecryptAES(encrypted string) (string, error) {
	key := viper.GetString("AES_ENC_KEY")
	if key == "" {
		return "", errorEmptyEncryptionKey
	}

	decoded, err := hex.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	aes, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", errorInitiatingCipher
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		return "", err
	}

	nonceHex := viper.GetString("AES_NONCE_KEY")
	nonce, err := hex.DecodeString(nonceHex)
	if err != nil {
		return "", err
	}

	plaintext, err := gcm.Open(nil, []byte(nonce), []byte(decoded), nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
