package main

import (
	"fmt"

	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/platform"
)

func main() {
	privateKeys := platform.CP.PrivateKeys()
	for _, privateKey := range privateKeys {
		fmt.Printf("%v\n", privateKey)
	}
}
