package ethereum

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// DepositCollateral deposits the amount of deposits into smart contract
func DepositCollateral(client *ethclient.Client, privateKey *ecdsa.PrivateKey,
	userAddress, contractAddress common.Address, value *big.Int, data []byte) error {
	signedTx, err := NewSignedTransaction(client, privateKey, userAddress, contractAddress, value, data)
	if err != nil {
		return err
	}
	return client.SendTransaction(context.Background(), signedTx)
}
