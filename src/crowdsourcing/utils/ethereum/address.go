package ethereum

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

// PrivateKeyAndAddress generates private key from the address from the given hex private key string
func PrivateKeyAndAddress(hexPrivateKey string) (*ecdsa.PrivateKey, common.Address) {
	privateKey, err := crypto.HexToECDSA(hexPrivateKey)
	if err != nil {
		log.Fatal("Key construction failed", err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Convert to public key failed:", publicKeyECDSA)
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return privateKey, address
}