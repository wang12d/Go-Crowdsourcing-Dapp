package plantform

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/task"
)

type plantform struct {
	workers     int           // The number of workers in the Plantform
	requesters  int           // The number of requesters in the Plantform
	keyIndex    int           // Key index used to get the private key
	tasks       []task.Task   // Tasks available in the Plantform
	privateKeys []string      // All of private keys
	lock        chan struct{} // The multithread lock of update index of plantform
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
			lock:        make(chan struct{}, 1),
			keyIndex:    0,
			workers:     0,
			requesters:  0,
			tasks:       make([]task.Task, 0),
			privateKeys: make([]string, numberOfAccount),
		}
	})
	// Get the mutex lock
	CP.lock <- struct{}{}
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
func (cp *plantform) NewAccount() (address string) {
	// The address must be token one by one
	var privateKey string
	<-cp.lock
	defer func() {
		cp.lock <- struct{}{}
	}()
	if cp.keyIndex >= numberOfAccount {
		return NULL
	}
	privateKey = cp.privateKeys[cp.keyIndex]
	fmt.Println("Index: ", cp.keyIndex)
	cp.keyIndex++
	return privateKey
}

// ReceiveTask receive task from requesters
// Only resigstered requester can post task
func (cp *plantform) ReceiveTask(postedTask task.Task) {
	cp.tasks = append(cp.tasks, postedTask)
}
