package main

import (
	"sync"

	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/platform"
)

func main() {
	var x sync.WaitGroup
	for i := 0; i < 70; i++ {
		x.Add(1)
		go func() {
			platform.CP.NewAccount()
			defer x.Done()
		}()
	}
	x.Wait()
}
