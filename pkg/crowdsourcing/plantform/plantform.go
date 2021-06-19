package plantform

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/task"
)

type plantform struct {
	workers     int                 // The number of workers in the Plantform
	requesters  int                 // The number of requesters in the Plantform
	keyIndex    int                 // Key index used to get the private key
	tasks       []*task.Task        // Tasks available in the Plantform
	privateKeys []string            // All of private keys
	addressLock chan struct{}       // The multithread lock of update index of plantform
	taskLock    chan struct{}       // The multithread lock for handler posted task cocurrently
	totalData   map[string][][]byte // The data needed by each task, their id as key
}

const (
	numberOfAccount = 60
	NULL            = ""
)

var (
	CP   *plantform
	once sync.Once
)

func init() {
	once.Do(func() {
		CP = &plantform{
			addressLock: make(chan struct{}, 1),
			taskLock:    make(chan struct{}, 1),
			keyIndex:    0,
			workers:     0,
			requesters:  0,
			tasks:       make([]*task.Task, 0),
			privateKeys: make([]string, numberOfAccount),
			totalData:   make(map[string][][]byte),
		}
	})
	// Get the mutex lock
	CP.addressLock <- struct{}{}
	CP.taskLock <- struct{}{}
	// Parser accounts file
	pwd, _ := os.Getwd()
	accountFile := filepath.Join(pwd, "pkg", "crowdsourcing", "plantform", "accounts.json")
	accountFileHandler, err := os.Open(accountFile)
	if err != nil {
		log.Fatalf("Read account file error: %v\n", err)
	}
	defer accountFileHandler.Close()

	accountData, _ := ioutil.ReadAll(accountFileHandler)

	var accounts map[string]interface{}

	json.Unmarshal([]byte(accountData), &accounts)

	switch acc := accounts["private_keys"].(type) {
	case map[string]interface{}:
		cnt := 0
		for _, val := range acc {
			CP.privateKeys[cnt] = val.(string)
			cnt++
		}
	default:
		log.Fatalf("Json parser error, cannot get private keys.")
	}

}

// NewAccount generates a private key to Ethereum account
// which is used for digital signature authentication
func (cp *plantform) NewAccount() string {
	// The address must be token one by one
	var privateKey string
	<-cp.addressLock
	defer func() {
		cp.addressLock <- struct{}{}
	}()
	if cp.keyIndex >= numberOfAccount {
		return NULL
	}
	privateKey = cp.privateKeys[cp.keyIndex]
	cp.keyIndex++
	return privateKey
}

// ReceiveTask receive task from requesters
// Only resigstered requester can post task
func (cp *plantform) ReceiveTask(postedTask *task.Task) {
	<-cp.taskLock
	cp.tasks = append(cp.tasks, postedTask)
	cp.totalData[postedTask.ID()] = make([][]byte, postedTask.WorkerRequired())
	defer func() {
		cp.taskLock <- struct{}{}
	}()
}

// TaskList return the current task posted by requesters to the plantform
func (cp *plantform) TaskList() []*task.Task {
	return cp.tasks
}

// SubmitTaskData store the data of the particular task
func (cp *plantform) SubmitTaskData(t *task.Task, data []byte, workerID int) {
	cp.totalData[t.ID()][workerID] = data
}

// CheckData return the data of the task
// WARNING: This function is only used for debugging
func (cp *plantform) CheckData(t *task.Task) [][]byte {
	return cp.totalData[t.ID()]
}
