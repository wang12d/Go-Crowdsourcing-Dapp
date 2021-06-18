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
	state      State
	task       task.Task
}

// NewWorker creates a new worker with initial value of parameters
func NewWorker() *Worker {
	return &Worker{
		address:    common.Address{},
		privateKey: nil,
		state:      INIT,
	}
}

// Register registe the worker into the Crowdsourcing plantform
func (w *Worker) Register() {
	if w.state != INIT { // Only worker at INIT can register
		return
	}
	privateKey, address := ethereum.PrivateKeyAndAddress(plantform.CP.NewAccount())
	w.privateKey, w.address = privateKey, address
	w.state = PENDING
}

// ParticipantTask decides whether the worker participant the current task
func (w *Worker) ParticipantTask() {

}

// SubmitData uploads the collected data to the task it participanted
func (w *Worker) SubmitData() {

}
