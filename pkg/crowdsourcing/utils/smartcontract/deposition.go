package smartcontract

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/utils/ethereum"
)

// DepositCollateral deposits the amount of deposits into smart contract
func DepositCollateral(client *ethclient.Client, privateKey *ecdsa.PrivateKey,
	userAddress, contractAddress common.Address, value *big.Int, data []byte) error {
	signedTx, err := ethereum.NewSignedTransaction(client, privateKey, userAddress, contractAddress, value, data)
	if err != nil {
		return err
	}
	return client.SendTransaction(context.Background(), signedTx)
}
