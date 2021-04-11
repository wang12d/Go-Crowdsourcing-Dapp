package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"nju.edu/cosec/crowdsourcing/src/crowdsourcing/lib"
	"nju.edu/cosec/crowdsourcing/src/crowdsourcing/utils/ethereum"
	"nju.edu/cosec/crowdsourcing/src/crowdsourcing/utils/smartcontract/crowdsourcing"
	"time"
)

const (
	localURL = "http://localhost:7545"
	requester = "7a1dcac34e4717acc4a21742c80c20b64104846332bfecee9377bf5d5a3c2b49"
	workerA = "ad41c4a8f5e3370787d8d62c651c898942cfa4427296142a1fbfd5da2c6ccbc4"
	workerB = "d252b9f840bd7c8da0ee30e5461f397fcea69ae5c5be4a0dced040e8addb5f34"
)

func main() {
	var err error
	client, err := ethclient.Dial(localURL)
	if err != nil {
		fmt.Println(fmt.Errorf("clinet dial error: %v", err))
	}
	chainID, err := client.ChainID(context.Background())
	// Initialize worker and requester transactor bind
	requesterPrivateKey, requesterAddress := ethereum.PrivateKeyAndAddress(requester)
	workerAPrivateKey, workerAAddress := ethereum.PrivateKeyAndAddress(workerA)
	workerBPrivateKey, workerBAddress := ethereum.PrivateKeyAndAddress(workerB)
	// requesterAuth, err := bind.NewKeyedTransactorWithChainID(requesterPrivateKey, chainID)
	requesterAuth := ethereum.KeyedTransactor(client, requesterPrivateKey,
		requesterAddress, chainID, big.NewInt(0))
	if err != nil {
		log.Fatal("Binding error: ", err)
	}

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