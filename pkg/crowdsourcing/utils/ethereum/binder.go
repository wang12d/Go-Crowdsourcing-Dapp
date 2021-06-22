package ethereum

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// KeyedTransactor create the transactor which is used to make transaction of the specific key
func KeyedTransactor(client *ethclient.Client, key *ecdsa.PrivateKey, address common.Address,
	id, value *big.Int) *bind.TransactOpts {
	auth, err := bind.NewKeyedTransactorWithChainID(
		key, id)
	if err != nil {
		log.Fatalf("keyed transactor construction failed, %v", err)
	}
	var price *big.Int
	if price, err = client.SuggestGasPrice(context.Background()); err != nil {
		log.Fatalf("get suggest gas price failed, %v", err)
	}
	auth.Nonce = big.NewInt(int64(GetNonce(context.Background(), client, address)))
	auth.GasLimit = gasLimit
	auth.Value = value
	auth.GasPrice = price
	return auth
}

// UpdateNonce update according to current state
func UpdateNonce(client *ethclient.Client, auth *bind.TransactOpts, address common.Address) {
	auth.Nonce = big.NewInt(int64(GetNonce(context.Background(), client, address)))
}
