package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"nju.edu/cosec/crowdsourcing/src/crowdsourcing"
	"time"
)

const (
	localURL = "http://localhost:7545"
	gasLimit = uint64(6721975)
	requester = "87a4bff39275278638b6f20e8295e7339230b10537f55e0a85a3e034d5a62659"
	workerA = "f4ca3ef5eb10ddcd50f748e9cc4552732e2cf958ef578306ddc6fe42ed5974f8"
	workerB = "2953ffbd499d269eab5e1029b8c3c7392397ba515ab8949a04902c3c4ebf6839"
)

func main() {
	var err error
	client, err := ethclient.Dial(localURL)
	if err != nil {
		fmt.Println(fmt.Errorf("clinet dial error: %v", err))
	}
	gasPrice, _ := client.SuggestGasPrice(context.Background())
	chainID, err := client.ChainID(context.Background())
	// Initialize worker and requester transactor bind
	requesterPrivateKey, requesterAddress := privateKeyAndAddress(requester)
	workerAPrivateKey, workerAAddress := privateKeyAndAddress(workerA)
	workerBPrivateKey, workerBAddress := privateKeyAndAddress(workerB)
	nonce := getNonce(context.Background(), client, requesterAddress)
	requesterAuth, err := bind.NewKeyedTransactorWithChainID(requesterPrivateKey, chainID)
	if err != nil {
		log.Fatal("Binding error: ", err)
	}
	workerAAuth, err := bind.NewKeyedTransactorWithChainID(workerAPrivateKey, chainID)
	if err != nil {
		log.Fatal("Binding error: ", err)
	}
	workerBAuth, err := bind.NewKeyedTransactorWithChainID(workerBPrivateKey, chainID)
	if err != nil {
		log.Fatal("Binding error: ", err)
	}
	requesterAuth.Value = big.NewInt(0)
	requesterAuth.Nonce = big.NewInt(int64(nonce))
	requesterAuth.GasLimit = gasLimit
	requesterAuth.GasPrice = gasPrice
	workerAAuth.GasPrice, workerBAuth.GasPrice = gasPrice, gasPrice
	workerAAuth.GasLimit, workerBAuth.GasLimit = gasLimit, gasLimit
	workerAAuth.Value, workerBAuth.Value = big.NewInt(0), big.NewInt(0)
	// Contract is deployed by requester
	contractAddress, _, instance, err := crowdsourcing.DeployCrowdsourcing(requesterAuth, client)
	if err != nil {
		log.Fatal("Contract deployment error:", err)
	}

	// Task collateral setting
	collateral := new(big.Int)
	collateral.SetString("8000000000000000000", 10)
	numberOfWorkers := big.NewInt(2)
	workerCollateral := new(big.Int)
	workerCollateral.SetString("4000000000000000000", 10)
	// Creating a transaction to deposit to the smart contract
	start := time.Now()
	if err = depositCollateral(client, requesterPrivateKey, requesterAddress, contractAddress, collateral, []byte{0x01}); err != nil {
		log.Fatal("Collateral deposition error: ", err)
	}
	if err = depositCollateral(client, workerAPrivateKey, workerAAddress, contractAddress, workerCollateral, []byte{0x00}); err != nil {
		log.Fatal("Collateral deposition error: ", err)
	}
	if err = depositCollateral(client, workerBPrivateKey, workerBAddress, contractAddress, workerCollateral, []byte{0x00}); err != nil {
		log.Fatal("Collateral deposition error: ", err)
	}
	// Requester publishes the transaction
	nonce = getNonce(context.Background(), client, requesterAddress)
	requesterAuth.Nonce = big.NewInt(int64(nonce))
	_, err = instance.PublishCrowdsourcingTask(requesterAuth, collateral, numberOfWorkers, "Simple image annotation tasks")
	if err != nil {
		log.Fatal("Task publish failed:", err)
	}
	fmt.Println("Task published, time cost:", time.Since(start))
	workerAAuth.Nonce = big.NewInt(int64(getNonce(context.Background(), client, workerAAddress)))
	if  _, err = instance.JoinCrowdsourcingTask(workerAAuth, requesterAddress); err != nil {
		log.Fatal("worker A joining the task error: ", err)
	}
	workerBAuth.Nonce = big.NewInt(int64(getNonce(context.Background(), client, workerBAddress)))
	if _, err = instance.JoinCrowdsourcingTask(workerBAuth, requesterAddress); err != nil {
		log.Fatal("worker B joining the task error: ", err)
	}
	fmt.Println("A and B have joined the task.")
	workerAAuth.Nonce = big.NewInt(int64(getNonce(context.Background(), client, workerAAddress)))
	if _, err = instance.SubmitData(workerAAuth, workerAAddress, []byte{0x00}); err != nil {
		log.Fatal("A submits data failed: ", err)
	}
	workerBAuth.Nonce = big.NewInt(int64(getNonce(context.Background(), client, workerBAddress)))
	if _, err = instance.SubmitData(workerBAuth, workerBAddress, []byte{0x01}); err != nil {
		log.Fatal("A submits data failed: ", err)
	}
	fmt.Println("Data have been submitted.")
	requesterAuth.Nonce = big.NewInt(int64(getNonce(context.Background(), client, requesterAddress)))
	if _, err = instance.Rewarding(requesterAuth, workerAAddress, true); err != nil {
		log.Fatal("Rewarding A failed: ", err)
	}
	requesterAuth.Nonce = big.NewInt(int64(getNonce(context.Background(), client, requesterAddress)))
	if _, err = instance.Rewarding(requesterAuth, workerBAddress, false); err != nil {
		log.Fatal("Rewarding B failed: ", err)
	}
	fmt.Println("Rewarding over.")
}

// getNonce return the pending nonce of the account, which is the nonce of next transaction
func getNonce(ctx context.Context, client *ethclient.Client, account common.Address) uint64 {
	nonce, err := client.PendingNonceAt(ctx, account)
	if err != nil {
		panic(fmt.Errorf("get nonce error: %v", err))
	}
	return nonce
}

// newSignedTransaction create a signed transaction using the client with private key
func newSignedTransaction(client * ethclient.Client, privateKey *ecdsa.PrivateKey, from, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	depositTx := types.LegacyTx{
		Nonce:    getNonce(context.Background(), client, from),
		GasPrice: gasPrice,
		Gas:      gasLimit,
		To:       &to,
		Value:	  value,
		Data:     data,
		V:        nil,
		R:        nil,
		S:        nil,
	}
	tx := types.NewTx(&depositTx)
	return types.SignTx(tx, types.HomesteadSigner{}, privateKey)
}

// depositCollateral deposits the amount of deposits into smart contract
func depositCollateral(client *ethclient.Client, privateKey *ecdsa.PrivateKey, userAddress, contractAddress common.Address, value *big.Int, data []byte) error {
	signedTx, err := newSignedTransaction(client, privateKey, userAddress, contractAddress, value, data)
	if err != nil {
		return err
	}
	return client.SendTransaction(context.Background(), signedTx)
}

// privateKeyAndAddress generates private key from the address from the given hex private key string
func privateKeyAndAddress(hexPrivateKey string) (*ecdsa.PrivateKey, common.Address) {
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