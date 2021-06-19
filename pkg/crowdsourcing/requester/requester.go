package requester

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/plantform"
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
}

// NewRequester returns a new requester
func NewRequester() *Requester {
	return &Requester{
		address:    common.Address{},
		privateKey: nil,
		publicKey:  nil,
		state:      INIT,
		task:       nil,
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
	privateKey, address := ethereum.PrivateKeyAndAddress(plantform.CP.NewAccount())
	r.privateKey, r.address = privateKey, address
	r.publicKey = &privateKey.PublicKey
	r.state = PENDING
}

// CreateTask create a task by the requester
func (r *Requester) CreateTask(workers, reward int, encKey []byte, description string) {
	if r.state != PENDING {
		return
	}
	r.task = task.NewTask(workers, reward, encKey, description, r.evaluation)
}

// PostTask post the task to plantform
func (r *Requester) PostTask() {
	plantform.CP.ReceiveTask(r.task)
}

// Task returns the task of the requester
func (r *Requester) Task() *task.Task {
	return r.task
}

// RewardList returns a list which indicate whether the corresponding data
// worth reward or not
func (r *Requester) RewardList() []bool {
	rewardList := make([]bool, r.task.WorkerRequired())
	for i, data := range plantform.CP.CheckData(r.task) {
		evalReuslt := r.task.Eval()(data)
		rewardList[i] = r.isRewardable(evalReuslt)
	}
	return rewardList
}
