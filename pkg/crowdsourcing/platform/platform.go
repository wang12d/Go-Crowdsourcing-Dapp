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
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/client"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/smartcontract/cplatform"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/task"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/utils/ethereum"
)

type platform struct {
	workers                 int      // The number of workers in the Plantform
	requesters              int      // The number of requesters in the Plantform
	keyIndex                int      // Key index used to get the private key
	privateKeys             []string // All of private keys of valid ganache blockchain address
	privateKey              *ecdsa.PrivateKey
	addressLock             chan struct{}               // The multithread lock of update index of platform
	totalData               map[string][][]byte         // The data needed by each task, their id as key
	address                 common.Address              // The deployed adress of smart contract
	taskParticipanted       map[string][]common.Address // Tasks' address of a worker participanted
	workerParticipationLock chan struct{}
	instance                *cplatform.Cplatform
	chainID                 *big.Int
	opts                    *bind.TransactOpts
	instanceAddress         common.Address
}

const (
	platformKeyIndex = 0
	numberOfAccount  = 60
	NULL             = ""
)

var (
	CP           *platform
	platformOnce sync.Once
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	platformOnce.Do(func() {
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
		chainID, err := client.CLIENT.ChainID(context.Background())
		if err != nil {
			log.Fatalf("Get chain ID error: %v\n", err)
		}
		platformAuth := ethereum.KeyedTransactor(client.CLIENT, privateKey, address, chainID, big.NewInt(0))
		platformAddress, _, platformInstance, err := cplatform.DeployCplatform(platformAuth, client.CLIENT)
		if err != nil {
			log.Fatalf("Deploy platform error: %v\n", err)
		}
		ethereum.UpdateNonce(client.CLIENT, platformAuth, address)
		CP = &platform{
			addressLock:             make(chan struct{}, 1),
			keyIndex:                1,
			workers:                 0,
			requesters:              0,
			privateKeys:             privateKeys,
			totalData:               make(map[string][][]byte),
			taskParticipanted:       make(map[string][]common.Address),
			workerParticipationLock: make(chan struct{}, 1),
			instanceAddress:         platformAddress,
			instance:                platformInstance,
			chainID:                 chainID,
			opts:                    platformAuth,
			privateKey:              privateKey,
			address:                 address,
		}
		// Get the mutex lock
		CP.addressLock <- struct{}{}
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

// InstanceAddress returns the address of deployed platform
func (cp *platform) InstanceAddress() common.Address {
	return cp.instanceAddress
}

// Instance returns the instance of crowdsourcing contract
func (cp *platform) Instance() *cplatform.Cplatform {
	return cp.instance
}

func (cp *platform) ChainID() *big.Int {
	return cp.chainID
}

func (cp *platform) Opts() *bind.TransactOpts {
	return cp.opts
}

// Register writes address to platform contract
func (cp *platform) Register(address common.Address) {
	if _, err := cp.instance.Register(cp.opts, address); err != nil {
		log.Fatalf("Register to platform error: %v\n", err)
	}
	ethereum.UpdateNonce(client.CLIENT, cp.opts, cp.address)
}

// AddingTask adds the task to the platform for using
func (cp *platform) WorkerParticipantTask(opts *bind.TransactOpts, t *task.Task, workerAddress common.Address) {
	if _, err := t.Instance().Register(opts, cp.instanceAddress); err != nil {
		log.Fatalf("Worker register crowdsourcing task error: %v\n", err)
	}
	ethereum.UpdateNonce(client.CLIENT, opts, workerAddress)
	t.AddingWorkers(workerAddress)
	cp.workerParticipationLock <- struct{}{}
	cp.taskParticipanted[t.Address().Hex()] = append(cp.taskParticipanted[t.Address().Hex()], workerAddress)
	<-cp.workerParticipationLock
}
