package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/utils/ethereum"
)

const (
	localURL  = "http://localhost:8545"
	requester = "97e0ba71d9ba8a41bf31c74b80484d58aa400ef2084a5ae15174d2394ccb3124"
	workerA   = "ef75b6f88870c312a7ae94c75472e43a2f13d1703e41811089cfd1aa67f8b5cd"
	workerB   = "b556fa3cf0b2629beda1ef1e6c1f484587079ea862cb3b39f34394811a635f56"
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
	if err = ethereum.DepositCollateral(client, requesterPrivateKey, requesterAddress, contractAddress, collateral, []byte{0x01}); err != nil {
		log.Fatal("Collateral deposition error: ", err)
	}
	if err = ethereum.DepositCollateral(client, workerAPrivateKey, workerAAddress, contractAddress, workerCollateral, []byte{0x00}); err != nil {
		log.Fatal("Collateral deposition error: ", err)
	}
	if err = ethereum.DepositCollateral(client, workerBPrivateKey, workerBAddress, contractAddress, workerCollateral, []byte{0x00}); err != nil {
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
	if _, err = instance.JoinCrowdsourcingTask(workerAAuth, requesterAddress); err != nil {
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
