package worker

import (
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/plantform"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/task"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/utils/ethereum"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/utils/smartcontract"
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

// Register registe the worker into the Crowdsourcing plantform
func (w *Worker) Register() {
	if w.state != INIT { // Only worker at INIT can register
		return
	}
	privateKey, address := plantform.CP.NewAccount()
	w.privateKey, w.address = privateKey, address
	w.publicKey = &privateKey.PublicKey
	w.state = PENDING
	w.opts = ethereum.KeyedTransactor(plantform.CP.Client(), w.privateKey,
		w.address, plantform.CP.ChianID(), big.NewInt(0))
}

// ParticipantTask decides whether the worker participant the current task
func (w *Worker) ParticipantTask() {
	if w.state != PENDING {
		return
	}
	taskList := plantform.CP.TaskList()
	for _, t := range taskList {
		t.TaskLock() //
		if t.RemainingWorkers().Cmp(zero) > 0 {
			// Worker must deposit the corresponding collaterals before participant task
			if err := smartcontract.DepositCollateral(plantform.CP.Client(), w.privateKey,
				w.address, plantform.CP.Address(), t.Collateral(), []byte{0x00}); err != nil {
				log.Fatalf("Worker deposite collaterals error: %v\n", err)
			}
			ethereum.UpdateNonce(plantform.CP.Client(), w.opts, w.address)
			if _, err := plantform.CP.Instance().JoinCrowdsourcingTask(w.opts, t.Address()); err != nil {
				log.Fatalf("Worker join crowdsourcing task error: %v\n", err)
			}
			ethereum.UpdateNonce(plantform.CP.Client(), w.opts, w.address)
			w.id = int(t.RemainingWorkers().Int64()) - 1
			w.task = t
			t.Participanting()
			t.TaskRelease()
			w.state = WORKING
			return
		}
		t.TaskRelease()
	}
}

// CollectData collects data from surrounding enviroment
func (w *Worker) CollectData(data []byte) {
	if w.state != WORKING {
		return
	}
	w.data = data
}

// SubmitData uploads the collected data to the task it participanted
func (w *Worker) SubmitData() {
	if w.state != WORKING {
		return
	}
	plantform.CP.SubmitTaskData(w.opts, w.address, w.task.Address(), w.task, w.data, w.id)
	w.state = FIN
}

// PublicKey returns the public key of the worker
func (w *Worker) PublicKey() ecdsa.PublicKey {
	return *w.publicKey
}

// S returns the current state of worker
func (w *Worker) S() State {
	return w.state
}

func (w *Worker) ID() int {
	return w.id
}
