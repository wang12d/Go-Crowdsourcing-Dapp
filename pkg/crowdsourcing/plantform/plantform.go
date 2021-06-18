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
	workers     int         // The number of workers in the Plantform
	requesters  int         // The number of requesters in the Plantform
	keyIndex    int         // Key index used to get the private key
	tasks       []task.Task // Tasks available in the Plantform
	privateKeys []string    // All of private keys
}

const (
	numberOfAccount = 60
)

var (
	CP   *plantform
	once sync.Once
)

func init() {
	once.Do(func() {
		CP = &plantform{
			keyIndex:    0,
			workers:     0,
			requesters:  0,
			tasks:       make([]task.Task, 0),
			privateKeys: make([]string, numberOfAccount),
		}
	})

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

// NewAccount generates a private key to Ethereum account which is used for digital signature
// authentication
func (cp *plantform) NewAccount() (address string) {
	if cp.keyIndex >= numberOfAccount {
		log.Fatalf("Plantform has reached the account limit\n")
	}
	privateKey := cp.privateKeys[cp.keyIndex]
	cp.keyIndex++
	return privateKey
}
