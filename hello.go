package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/wang12d/Go-Crowdsourcing-DApp/src/crowdsourcing/lib"
	"github.com/wang12d/Go-Crowdsourcing-DApp/src/crowdsourcing/utils/ethereum"
	"github.com/wang12d/Go-Crowdsourcing-DApp/src/crowdsourcing/utils/smartcontract/crowdsourcing"
	"log"
	"math/big"
	"time"
)

const (
	localURL = "http://localhost:7545"
	requester = "ebc9bfe431c9408f463613c281c9ff9bf475925c7b7dcee6778ca9320d62f072"
	workerA = "0d39d481bf81aa4d52bfb41c4f4e26716036f0aecd9a534353ae543875263782"
	workerB = "a117d32ed4a19a8e14694975d5ebd887f6d4f40f69177b1117590466842c0295"
)

func main() {
	var err error
	client, err := ethclient.Dial(localURL)
	if err != nil {
		log.Fatalf(fmt.Sprintf("clinet dial error: %v", err))
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal("Binding error: ", err)
	}
	// Initialize worker and requester transactor bind
	requesterPrivateKey, requesterAddress := ethereum.PrivateKeyAndAddress(requester)
	workerAPrivateKey, workerAAddress := ethereum.PrivateKeyAndAddress(workerA)
	workerBPrivateKey, workerBAddress := ethereum.PrivateKeyAndAddress(workerB)
	requesterAuth := ethereum.KeyedTransactor(client, requesterPrivateKey,
		requesterAddress, chainID, big.NewInt(0))

	// Contract is deployed by requester
	contractAddress, _, instance, err := lib.DeployCrowdsourcing(requesterAuth, client)
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
	if err = crowdsourcing.DepositCollateral(client, requesterPrivateKey, requesterAddress, contractAddress, collateral, []byte{0x01}); err != nil {
		log.Fatal("Collateral deposition error: ", err)
	}
	if err = crowdsourcing.DepositCollateral(client, workerAPrivateKey, workerAAddress, contractAddress, workerCollateral, []byte{0x00}); err != nil {
		log.Fatal("Collateral deposition error: ", err)
	}
	if err = crowdsourcing.DepositCollateral(client, workerBPrivateKey, workerBAddress, contractAddress, workerCollateral, []byte{0x00}); err != nil {
		log.Fatal("Collateral deposition error: ", err)
	}
	// Requester publishes the transaction
	ethereum.UpdateNonce(client, requesterAuth, requesterAddress)
	_, err = instance.PublishCrowdsourcingTask(requesterAuth, collateral, numberOfWorkers, "Simple image annotation tasks")
	if err != nil {
		log.Fatal("Task publish failed:", err)
	}
	fmt.Println("Task published, time cost:", time.Since(start))
	workerAAuth := ethereum.KeyedTransactor(client, workerAPrivateKey,
		workerAAddress, chainID, big.NewInt(0))
	if  _, err = instance.JoinCrowdsourcingTask(workerAAuth, requesterAddress); err != nil {
		log.Fatal("worker A joining the task error: ", err)
	}
	workerBAuth := ethereum.KeyedTransactor(client, workerBPrivateKey,
		workerBAddress, chainID, big.NewInt(0))
	if _, err = instance.JoinCrowdsourcingTask(workerBAuth, requesterAddress); err != nil {
		log.Fatal("worker B joining the task error: ", err)
	}
	fmt.Println("A and B have joined the task.")
	ethereum.UpdateNonce(client, workerAAuth, workerAAddress)
	if _, err = instance.SubmitData(workerAAuth, workerAAddress, []byte{0x00}); err != nil {
		log.Fatal("A submits data failed: ", err)
	}
	ethereum.UpdateNonce(client, workerBAuth, workerBAddress)
	if _, err = instance.SubmitData(workerBAuth, workerBAddress, []byte{0x01}); err != nil {
		log.Fatal("A submits data failed: ", err)
	}
	fmt.Println("Data have been submitted.")
	ethereum.UpdateNonce(client, requesterAuth, requesterAddress)
	if _, err = instance.Rewarding(requesterAuth, workerAAddress, true); err != nil {
		log.Fatal("Rewarding A failed: ", err)
	}
	ethereum.UpdateNonce(client, requesterAuth, requesterAddress)
	if _, err = instance.Rewarding(requesterAuth, workerBAddress, false); err != nil {
		log.Fatal("Rewarding B failed: ", err)
	}
	fmt.Println("Rewarding over.")
}