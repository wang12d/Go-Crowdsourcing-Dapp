package task

import (
	"log"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/client"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/platform"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/smartcontract/ctask"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/utils/cryptograph"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/utils/ethereum"
)

var (
	id     = 0
	idLock = make(chan struct{}, 1)
)

// EvalFunc is the Evaluation function
type EvalFunc func([]byte) float64

type Task struct {
	workerRequired   *big.Int
	remainingWorkers *big.Int
	reward           *big.Int
	workerLock       chan struct{}
	encryptor        cryptograph.Encryptor // The encryptor used to encrypted the uploaded data in crowdsourcing
	description      string                // The description of crowdsourcing task
	id               string
	eval             EvalFunc
	address          common.Address // The address of deployed crowdsourcing task
	collaterals      *big.Int
	workerAddresses  []common.Address // The address of participant workers
	instance         *ctask.Ctask     // The depolyed smartcontract object
	data             [][]byte         // Data submitted by workers
}

func init() {
	idLock <- struct{}{}
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func NewTask(workerRequired, reward *big.Int, encryptor cryptograph.Encryptor,
	address common.Address, description string, eval EvalFunc, instance *ctask.Ctask) *Task {
	<-idLock
	currentID := id
	id++
	idLock <- struct{}{}
	collaterals := new(big.Int)
	collaterals.SetString(reward.String(), 10)
	remainingWorkers := new(big.Int)
	remainingWorkers.SetString(workerRequired.String(), 10)
	newTask := &Task{
		workerRequired:   workerRequired,
		remainingWorkers: remainingWorkers,
		reward:           reward,
		encryptor:        encryptor,
		description:      description,
		workerLock:       make(chan struct{}, 1),
		id:               strconv.Itoa(currentID),
		eval:             eval,
		collaterals:      collaterals,
		address:          address,
		workerAddresses:  make([]common.Address, workerRequired.Int64()),
		instance:         instance,
		data:             make([][]byte, workerRequired.Int64()),
	}
	// Get the worker lock of current worker
	newTask.workerLock <- struct{}{}
	return newTask
}

// RemainingWorkers returns the remaining workers of the task
func (t *Task) RemainingWorkers() *big.Int {
	return t.remainingWorkers
}

// WorkerRequired returns the total worker needed for the task
func (t *Task) WorkerRequired() *big.Int {
	return t.workerRequired
}

// Reward returns the reward of the task
func (t *Task) Reward() *big.Int {
	return t.reward
}

// Collateral returns the collateral a worker needed in order to
// participant the task
func (t *Task) Collateral() *big.Int {
	return t.collaterals
}

// EncKey return the encryption key of data needed
func (t *Task) Encryptor() cryptograph.Encryptor {
	return t.encryptor
}

func (t *Task) TaskLock() {
	<-t.workerLock
}

func (t *Task) TaskRelease() {
	t.workerLock <- struct{}{}
}

// Address return the address of requester who posted the task
func (t *Task) Address() common.Address {
	return t.address
}

// WorkerAddresses returns the total addresses of workers who participants the task
func (t *Task) WorkerAddresses() []common.Address {
	return t.workerAddresses
}

// Instance returns the deployed task object in Ethereum
func (t *Task) Instance() *ctask.Ctask {
	return t.instance
}

// Participating indicates whether the participating of task success
func (t *Task) Participating(opts *bind.TransactOpts, workerAddress common.Address) bool {
	if t.remainingWorkers.Cmp(big.NewInt(0)) <= 0 {
		return false
	}
	if _, err := t.instance.Register(opts, platform.CP.InstanceAddress()); err != nil {
		log.Fatalf("Worker register crowdsourcing task error: %v\n", err)
	}
	ethereum.UpdateNonce(client.CLIENT, opts, workerAddress)
	t.workerAddresses[t.remainingWorkers.Int64()-1] = workerAddress
	t.remainingWorkers.Sub(t.remainingWorkers, big.NewInt(1))
	return true
}

// Description returns the description of the task
func (t *Task) Description() string {
	return t.description
}

func (t *Task) Eval() EvalFunc {
	return t.eval
}

// ID returns the unique id the crowdsourcing task
func (t *Task) ID() string {
	return t.id
}

// SubmitData receives data submitted from worker
func (t *Task) SubmitData(opts *bind.TransactOpts, workerID int, data []byte) {
	t.data[workerID] = data
	t.instance.SubmitData(opts, data)
}

// Data returns the data submitted from workers
func (t *Task) Data() [][]byte {
	return t.data
}
