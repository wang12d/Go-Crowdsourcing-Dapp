package requester

import (
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/platform"
	ctask "github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/smartcontract/task"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/task"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/utils/encoder"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/utils/ethereum"
)

const (
	mean  = 1000.0
	sigma = 250.0
	EPS   = 1e-8
)

type Requester struct {
	address    common.Address
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
	state      State
	task       *task.Task
	opts       *bind.TransactOpts
}

// NewRequester returns a new requester
func NewRequester() *Requester {
	return &Requester{
		address:    common.Address{},
		privateKey: nil,
		publicKey:  nil,
		state:      INIT,
		task:       nil,
		opts:       nil,
	}
}

// evaluation is the data quality evaluation function
func (r *Requester) evaluation(data []byte) float64 {
	numeric := encoder.Float64FromBytes(data)
	return numeric - mean
}

// isReadable determines whether the current data quality is worth reward
func (r *Requester) isReward(evalResult float64) bool {
	return evalResult-3*sigma <= EPS && evalResult+3*sigma >= EPS
}

// Register register the worker into the Crowdsourcing platform
func (r *Requester) Register() {
	if r.state != INIT { // Only worker at INIT can register
		return
	}
	privateKey, address := platform.CP.NewAccount()
	r.privateKey, r.address = privateKey, address
	r.publicKey = &privateKey.PublicKey
	r.state = PENDING
	r.opts = ethereum.KeyedTransactor(platform.CP.Client(), r.privateKey,
		r.address, platform.CP.ChainID(), big.NewInt(0))
}

// CreateTask create a task by the requester
// func (r *Requester) CreateTask(workers, reward int, encKey []byte, description string) {
// 	if r.state != PENDING {
// 		return
// 	}
// 	r.task = task.NewTask(big.NewInt(int64(workers)), big.NewInt(int64(reward)),
// 		encKey, r.address, description, r.evaluation)
// 	r.state = WAITING
// }

// PostTask create and post the task to platform
func (r *Requester) PostTask(workers, reward int, encKey []byte, description string) {
	// Before post task, requester must deposit the corresponding collaterals
	if r.state != PENDING {
		return
	}
	_workers, _rewards := big.NewInt(int64(workers)), big.NewInt(int64(reward))

	collateral := new(big.Int)
	collateral.Mul(_workers, _rewards)
	// Now publishing the task to blockchain
	taskAddress, _, taskContract, err := ctask.DeployTask(r.opts, platform.CP.Client(), _workers, _rewards, description)
	ethereum.UpdateNonce(platform.CP.Client(), r.opts, r.address)
	if err != nil {
		log.Fatalf("Publish task error: %v\n", err)
	}
	if err := ethereum.DepositCollateral(platform.CP.Client(), r.privateKey, r.address, taskAddress, collateral, []byte{0x01}); err != nil {
		log.Fatalf("Requester deposite collaterals error: %v\n", err)
	}
	ethereum.UpdateNonce(platform.CP.Client(), r.opts, r.address)
	r.task = task.NewTask(big.NewInt(int64(workers)), big.NewInt(int64(reward)),
		encKey, taskAddress, description, r.evaluation, taskContract)
	r.state = WAITING
}

// Task returns the task of the requester
func (r *Requester) Task() *task.Task {
	return r.task
}

// Rewarding returns a list which indicate whether the corresponding data
// worth reward or not
func (r *Requester) Rewarding() []bool {
	rewardList := make([]bool, r.task.WorkerRequired().Int64())
	data := r.task.Data()
	for i, data := range data {
		evalResult := r.task.Eval()(data)
		isok := r.isReward(evalResult)
		rewardList[i] = isok
		if _, err := r.task.Instance().Rewarding(r.opts, r.task.WorkerAddresses()[i], isok); err != nil {
			log.Fatalf("Rewarding error: %v\n", err)
		}
		ethereum.UpdateNonce(platform.CP.Client(), r.opts, r.address)
	}
	return rewardList
}
