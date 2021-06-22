package platform

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/task"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/utils/ethereum"
)

type platform struct {
	workers     int                 // The number of workers in the Plantform
	requesters  int                 // The number of requesters in the Plantform
	keyIndex    int                 // Key index used to get the private key
	tasks       []*task.Task        // Tasks available in the Plantform
	privateKeys []string            // All of private keys
	addressLock chan struct{}       // The multithread lock of update index of platform
	taskLock    chan struct{}       // The multithread lock for handler posted task cocurrently
	totalData   map[string][][]byte // The data needed by each task, their id as key
	client      *ethclient.Client   // The client of whole ethereum network
	address     common.Address
	instance    *crowdsourcing.Crowdsourcing
	chainID     *big.Int
}

const (
	platformKeyIndex = 0
	numberOfAccount   = 60
	NULL              = ""
	localURL          = "http://localhost:8545"
)

var (
	CP           *platform
	platformOnce sync.Once
)

func init() {
	platformOnce.Do(func() {
		client, err := ethclient.Dial(localURL)
		if err != nil {
			log.Fatalf("Create ethereum client error: %v\n", err)
		}
		// Parser accounts file
		pwd, _ := os.Getwd()
		accountFile := filepath.Join(pwd, "pkg", "crowdsourcing", "platform", "accounts.json")
		accountFileHandler, err := os.Open(accountFile)
		if err != nil {
			log.Fatalf("Read account file error: %v\n", err)
		}
		defer accountFileHandler.Close()

		accountData, _ := ioutil.ReadAll(accountFileHandler)

		var accounts map[string]interface{}

		json.Unmarshal([]byte(accountData), &accounts)

		var privateKeys []string
		switch acc := accounts["private_keys"].(type) {
		case map[string]interface{}:
			cnt := 0
			privateKeys = make([]string, len(acc))
			for _, val := range acc {
				privateKeys[cnt] = val.(string)
				cnt++
			}
		default:
			log.Fatalf("Json parser error, cannot get private keys.")
		}

		// Now deploying the platform contract with its private key
		privateKey, address := ethereum.PrivateKeyAndAddress(privateKeys[platformKeyIndex])
		chainID, err := client.ChainID(context.Background())
		if err != nil {
			log.Fatalf("Get chain ID error: %v\n", err)
		}
		platformAuth := ethereum.KeyedTransactor(client, privateKey, address, chainID, big.NewInt(0))
		platformAddress, _, platformInstance, err := crowdsourcing.DeployCrowdsourcing(platformAuth, client)
		if err != nil {
			log.Fatalf("Deploy platform error: %v\n", err)
		}
		CP = &platform{
			addressLock: make(chan struct{}, 1),
			taskLock:    make(chan struct{}, 1),
			keyIndex:    1,
			workers:     0,
			requesters:  0,
			tasks:       make([]*task.Task, 0),
			privateKeys: privateKeys,
			totalData:   make(map[string][][]byte),
			client:      client,
			address:     platformAddress,
			instance:    platformInstance,
			chainID:     chainID,
		}
		// Get the mutex lock
		CP.addressLock <- struct{}{}
		CP.taskLock <- struct{}{}
	})
}

// NewAccount returns a new private and ethereum address
func (cp *platform) NewAccount() (*ecdsa.PrivateKey, common.Address) {
	// The address must be taken one by one
	var privateKey string
	<-cp.addressLock
	defer func() {
		cp.addressLock <- struct{}{}
	}()
	if cp.keyIndex >= numberOfAccount {
		return nil, common.Address{}
	}
	privateKey = cp.privateKeys[cp.keyIndex]
	cp.keyIndex++
	return ethereum.PrivateKeyAndAddress(privateKey)
}

// Client returns the Ethereum client
func (cp *platform) Client() *ethclient.Client {
	return cp.client
}

// Address returns the address of deployed Client
func (cp *platform) Address() common.Address {
	return cp.address
}

// Instance returns the instance of crowdsourcing contract
func (cp *platform) Instance() *crowdsourcing.Crowdsourcing {
	return cp.instance
}

// ReceiveTask receive task from requesters
// Only resigstered requester can post task
func (cp *platform) ReceiveTask(opts *bind.TransactOpts, address common.Address, postedTask *task.Task) {
	<-cp.taskLock
	cp.tasks = append(cp.tasks, postedTask)
	// Publish the crowdsourcing task to blockchain
	cp.instance.PublishCrowdsourcingTask(opts, postedTask.Collateral(), postedTask.WorkerRequired(), postedTask.Description())
	ethereum.UpdateNonce(cp.client, opts, address)
	cp.totalData[postedTask.ID()] = make([][]byte, postedTask.WorkerRequired().Int64())
	defer func() {
		cp.taskLock <- struct{}{}
	}()
}

// TaskList return the current task posted by requesters to the platform
func (cp *platform) TaskList() []*task.Task {
	return cp.tasks
}

// SubmitTaskData store the data of the particular task
func (cp *platform) SubmitTaskData(opts *bind.TransactOpts, srcAddress, dstAddress common.Address,
	t *task.Task, data []byte, workerID int) {
	_, err := cp.instance.SubmitData(opts, dstAddress, data)
	if err != nil {
		log.Fatalf("Submit task data error: %v\n", err)
	}
	ethereum.UpdateNonce(cp.client, opts, srcAddress)
	cp.totalData[t.ID()][workerID] = data
}

// ParticipantCrowdsourcingTask makes the worker join the selected task
func (cp *platform) ParticipantCrowdsourcingTask(opts *bind.TransactOpts, workerAddress, taskAddress common.Address) {
	if _, err := cp.Instance().JoinCrowdsourcingTask(opts, taskAddress); err != nil {
		log.Fatalf("Worker join crowdsourcing task error: %v\n", err)
	}
	ethereum.UpdateNonce(cp.Client(), opts, workerAddress)
}

// CheckData return the data of the task
// WARNING: This function is only used for debugging
func (cp *platform) CheckData(t *task.Task) [][]byte {
	return cp.totalData[t.ID()]
}

func (cp *platform) ChainID() *big.Int {
	return cp.chainID
}