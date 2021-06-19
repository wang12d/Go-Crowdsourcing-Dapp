package worker

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/plantform"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/task"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/utils/ethereum"
)

type Worker struct {
	address    common.Address
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
	state      State
	task       *task.Task
	id         int
	data       []byte
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
	}
}

// Register registe the worker into the Crowdsourcing plantform
func (w *Worker) Register() {
	if w.state != INIT { // Only worker at INIT can register
		return
	}
	privateKey, address := ethereum.PrivateKeyAndAddress(plantform.CP.NewAccount())
	w.privateKey, w.address = privateKey, address
	w.publicKey = &privateKey.PublicKey
	w.state = PENDING
}

// ParticipantTask decides whether the worker participant the current task
func (w *Worker) ParticipantTask() {
	if w.state != PENDING {
		return
	}
	taskList := plantform.CP.TaskList()
	for _, t := range taskList {
		t.TaskLock() //
		if t.RemainingWorkers() > 0 {
			w.id = t.RemainingWorkers() - 1
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
	plantform.CP.SubmitTaskData(w.task, w.data, w.id)
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
