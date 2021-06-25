package cryptograph

import (
	"crypto/aes"
	"crypto/cipher"
	"log"
)

// DecryptData decrypts data with designated key
func DecryptData(data, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("Decryption error: %v\n", err)
	}
	cbc := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(data))
	cbc.CryptBlocks(decrypted, data)
	return PKCS5Trimming(decrypted)
}