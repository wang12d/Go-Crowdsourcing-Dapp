package task

import (
	"strconv"
)

var (
	id     = 0
	idLock = make(chan struct{}, 1)
)

type Key []byte

type Task struct {
	workerRequired   int
	remainingWorkers int
	reward           int
	workerLock       chan struct{}
	encKey           Key // The key used to encrypted the uploaded data in crowdsourcing
	description      string
	id               string
}

func init() {
	idLock <- struct{}{}
}

func NewTask(workerRequired, reward int, encKey Key, description string) *Task {
	<-idLock
	currentID := id
	id++
	idLock <- struct{}{}
	newTask := &Task{
		workerRequired:   workerRequired,
		remainingWorkers: workerRequired,
		reward:           reward,
		encKey:           encKey,
		description:      description,
		workerLock:       make(chan struct{}, 1),
		id:               strconv.Itoa(currentID),
	}
	// Get the worker lock of current worker
	newTask.workerLock <- struct{}{}
	return newTask
}

// RemainingWorkers returns the remaining workers of the task
func (t *Task) RemainingWorkers() int {
	return t.remainingWorkers
}

// WorkerRequired returns the total worker needed for the task
func (t *Task) WorkerRequired() int {
	return t.workerRequired
}

// Reward returns the reward of the task
func (t *Task) Reward() int {
	return t.reward
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

// Participanting indicates whether the participanting of task success
func (t *Task) Participanting() bool {
	if t.remainingWorkers <= 0 {
		return false
	}
	t.remainingWorkers--
	return true
}

// Description returns the description of the task
func (t *Task) Description() string {
	return t.description
}

// ID returns the unique id the crowdsourcing task
func (t *Task) ID() string {
	return t.id
}
