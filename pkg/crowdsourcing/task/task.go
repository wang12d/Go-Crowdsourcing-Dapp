package task

import (
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)

var (
	id     = 0
	idLock = make(chan struct{}, 1)
)

type Key []byte

// The Evaluation function
type EvalFunc func([]byte) float64

type Task struct {
	workerRequired   *big.Int
	remainingWorkers *big.Int
	reward           *big.Int
	workerLock       chan struct{}
	encKey           Key // The key used to encrypted the uploaded data in crowdsourcing
	description      string
	id               string
	eval             EvalFunc
	address          common.Address
	collaterals      *big.Int
}

func init() {
	idLock <- struct{}{}
}

func NewTask(workerRequired, reward *big.Int, encKey Key, address common.Address, description string, eval EvalFunc) *Task {
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
		encKey:           encKey,
		description:      description,
		workerLock:       make(chan struct{}, 1),
		id:               strconv.Itoa(currentID),
		eval:             eval,
		collaterals:      collaterals,
		address:          address,
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
func (t *Task) EncKey() Key {
	return t.encKey
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

// Participanting indicates whether the participanting of task success
func (t *Task) Participanting() bool {
	if t.remainingWorkers.Cmp(big.NewInt(0)) < 0 {
		return false
	}
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
