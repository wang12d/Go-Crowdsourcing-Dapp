package worker

import (
	"crypto/ecdsa"
	"log"
	"math/big"
	"time"

	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/metrics"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/platform"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/task"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/utils/ethereum"
)

var (
	zero = big.NewInt(0)
	lock = make(chan struct{}, 1)
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	lock <- struct{}{}
}

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
	caller := metrics.GetCallerName()
	defer metrics.GetMemoryStatus(caller)
	defer metrics.TimeCost(time.Now(), caller)
	if w.state != INIT { // Only worker at INIT can register
		return
	}
	privateKey, address := platform.CP.NewAccount()
	w.privateKey, w.address = privateKey, address
	w.publicKey = &privateKey.PublicKey
	w.state = PENDING
	w.opts = ethereum.KeyedTransactor(platform.CP.Client(), w.privateKey,
		w.address, platform.CP.ChainID(), big.NewInt(0))
	<-lock
	platform.CP.Register(w.address)
	lock <- struct{}{}
}

// ParticipantTask decides whether the worker participant the current task
func (w *Worker) ParticipantTask(t *task.Task) {
	caller := metrics.GetCallerName()
	defer metrics.GetMemoryStatus(caller)
	defer metrics.TimeCost(time.Now(), caller)
	if w.state != PENDING {
		return
	}
	t.TaskLock() //
	if t.RemainingWorkers().Cmp(zero) > 0 {
		w.id = int(t.RemainingWorkers().Int64()) - 1
		w.task = t
		t.Participating(w.opts, w.address)
		t.TaskRelease()
		w.state = WORKING
		return
	}
	t.TaskRelease()
}

// CollectData collects data from surrounding environment
func (w *Worker) CollectData(data []byte) {
	if w.state != WORKING {
		return
	}
	var err error
	w.data, err = w.task.Encryptor().EncryptData(data)
	if err != nil {
		log.Fatalf("Worker collecting data error: %v\n", err)
	}
}

// SubmitData uploads the collected data to the task it participated
func (w *Worker) SubmitData() {
	caller := metrics.GetCallerName()
	defer metrics.GetMemoryStatus(caller)
	defer metrics.TimeCost(time.Now(), caller)
	if w.state != WORKING {
		return
	}
	w.task.SubmitData(w.opts, w.id, w.data)
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
