package cryptograph

import (
	"crypto/aes"
	"crypto/cipher"
	"log"
)

// EncryptData encrypts with the designated key
func EncryptData(data, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("Encryption error: %v\n", err)
	}
	cbc := cipher.NewCBCEncrypter(block, iv)
	paddedData := PKCS5Padding(data, block.BlockSize())
	encrypted := make([]byte, len(paddedData))
	cbc.CryptBlocks(encrypted, paddedData)

	return encrypted
}