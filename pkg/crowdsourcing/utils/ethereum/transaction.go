package ethereum

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

// NewSignedTransaction create a signed transaction using the client with private key
func NewSignedTransaction(client * ethclient.Client, privateKey *ecdsa.PrivateKey, from, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	depositTx := types.LegacyTx{
		Nonce:    GetNonce(context.Background(), client, from),
		GasPrice: gasPrice,
		Gas:      gasLimit,
		To:       &to,
		Value:    value,
		Data:     data,
		V:        nil,
		R:        nil,
		S:        nil,
	}
	tx := types.NewTx(&depositTx)
	return types.SignTx(tx, types.HomesteadSigner{}, privateKey)
}