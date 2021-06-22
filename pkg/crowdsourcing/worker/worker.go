package worker

import (
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/platform"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/task"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/utils/ethereum"
)

var (
	zero = big.NewInt(0)
)

type Worker struct {
	address    common.Address
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
	state      State
	task       *task.Task
	id         int
	data       []byte
	opts       *bind.TransactOpts
}

// NewWorker creates a new worker with initial value of parameters
func NewWorker() *Worker {
	return &Worker{
		address:    common.Address{},
		privateKey: nil,
		publicKey:  nil,
		state:      INIT,
		id:         -1,
		task:       nil,
		data:       nil,
		opts:       nil,
	}
}

// Register register the worker into the Crowdsourcing platform
func (w *Worker) Register() {
	if w.state != INIT { // Only worker at INIT can register
		return
	}
	privateKey, address := platform.CP.NewAccount()
	w.privateKey, w.address = privateKey, address
	w.publicKey = &privateKey.PublicKey
	w.state = PENDING
	w.opts = ethereum.KeyedTransactor(platform.CP.Client(), w.privateKey,
		w.address, platform.CP.ChainID(), big.NewInt(0))
}

// ParticipantTask decides whether the worker participant the current task
func (w *Worker) ParticipantTask() {
	if w.state != PENDING {
		return
	}
	taskList := platform.CP.TaskList()
	for _, t := range taskList {
		t.TaskLock() //
		if t.RemainingWorkers().Cmp(zero) > 0 {
			// Worker must deposit the corresponding collaterals before participant task
			if err := ethereum.DepositCollateral(platform.CP.Client(), w.privateKey,
				w.address, platform.CP.Address(), t.Collateral(), []byte{0x00}); err != nil {
				log.Fatalf("Worker deposite collaterals error: %v\n", err)
			}
			ethereum.UpdateNonce(platform.CP.Client(), w.opts, w.address)
			platform.CP.ParticipantCrowdsourcingTask(w.opts, w.address, t.Address())
			w.id = int(t.RemainingWorkers().Int64()) - 1
			w.task = t
			t.Participating(w.address)
			t.TaskRelease()
			w.state = WORKING
			return
		}
		t.TaskRelease()
	}
}

// CollectData collects data from surrounding environment
func (w *Worker) CollectData(data []byte) {
	if w.state != WORKING {
		return
	}
	w.data = data
}

// SubmitData uploads the collected data to the task it participated
func (w *Worker) SubmitData() {
	if w.state != WORKING {
		return
	}
	platform.CP.SubmitTaskData(w.opts, w.address, w.task.Address(), w.task, w.data, w.id)
	w.state = FIN
}

// PublicKey returns the public key of the worker
func (w *Worker) PublicKey() ecdsa.PublicKey {
	return *w.publicKey
}

func (w *Worker) Address() common.Address {
	return w.address
}

func (w *Worker) ID() int {
	return w.id
}
