package requester

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/plantform"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/task"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/utils/encoder"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/utils/ethereum"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/utils/smartcontract"
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

// evaluation is the data qulity evaluation function
func (r *Requester) evaluation(data []byte) float64 {
	numeric := encoder.Float64FromBytes(data)
	return numeric - mean
}

// isRewardable determines whether the current data quality is worth reward
func (r *Requester) isRewardable(evalResult float64) bool {
	return evalResult-3*sigma <= EPS && evalResult+3*sigma >= EPS
}

// Register registe the worker into the Crowdsourcing plantform
func (r *Requester) Register() {
	if r.state != INIT { // Only worker at INIT can register
		return
	}
	privateKey, address := plantform.CP.NewAccount()
	r.privateKey, r.address = privateKey, address
	r.publicKey = &privateKey.PublicKey
	r.state = PENDING
	r.opts = ethereum.KeyedTransactor(plantform.CP.Client(), r.privateKey,
		r.address, plantform.CP.ChianID(), big.NewInt(0))
}

// CreateTask create a task by the requester
func (r *Requester) CreateTask(workers, reward int, encKey []byte, description string) {
	if r.state != PENDING {
		return
	}
	r.task = task.NewTask(big.NewInt(int64(workers)), big.NewInt(int64(reward)),
		encKey, r.address, description, r.evaluation)
	r.state = WAITING
}

// PostTask post the task to plantform
func (r *Requester) PostTask() {
	// Before post task, requester must deposit the corresponding collaterals
	collaters := new(big.Int)
	collaters.Mul(r.task.Reward(), r.task.WorkerRequired())
	// Now publishing the task to blockchain
	if err := smartcontract.DepositCollateral(plantform.CP.Client(), r.privateKey, r.address, plantform.CP.Address(), collaters, []byte{0x01}); err != nil {
		log.Fatalf("Requester deposite collaterals error: %v\n", err)
	}
	ethereum.UpdateNonce(plantform.CP.Client(), r.opts, r.address)
	plantform.CP.ReceiveTask(r.opts, r.address, r.task)
}

// Task returns the task of the requester
func (r *Requester) Task() *task.Task {
	return r.task
}

// RewardList returns a list which indicate whether the corresponding data
// worth reward or not
func (r *Requester) Rewarding() []bool {
	rewardList := make([]bool, r.task.WorkerRequired().Int64())
	for i, data := range plantform.CP.CheckData(r.task) {
		evalReuslt := r.task.Eval()(data)
		isok := r.isRewardable(evalReuslt)
		rewardList[i] = isok
		fmt.Println(r.task.WorkerAddresses()[i])
		if _, err := plantform.CP.Instance().Rewarding(r.opts, r.task.WorkerAddresses()[i], isok); err != nil {
			log.Fatalf("Rewarding error: %v\n", err)
		}
		ethereum.UpdateNonce(plantform.CP.Client(), r.opts, r.address)
	}
	return rewardList
}
