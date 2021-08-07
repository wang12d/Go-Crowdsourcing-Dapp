package client

import (
	"log"
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	localURL = "http://localhost:8545"
)

var (
	initOnce sync.Once
	CLIENT   *ethclient.Client
)

func init() {
	initOnce.Do(func() {
		var err error
		CLIENT, err = ethclient.Dial(localURL)
		if err != nil {
			log.Fatalf("Create client error: %v", err)
		}
	})
}
