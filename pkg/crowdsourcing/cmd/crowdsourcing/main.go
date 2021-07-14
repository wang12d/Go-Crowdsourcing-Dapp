package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"fmt"
	mrand "math/rand"
	"sync"

	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/requester"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/utils/encoder"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/worker"
)

const (
	numberOfWorkers = 5
	rewards         = 1
)

type epk rsa.PublicKey
type esk rsa.PrivateKey

func (pk *epk) EncryptData(data []byte) ([]byte, error) {
	hash := sha512.New()
	ppk := rsa.PublicKey(*pk)
	return rsa.EncryptOAEP(hash, rand.Reader, &ppk, data, nil)
}

func (sk *esk) DecryptData(data []byte) ([]byte, error) {
	hash := sha512.New()
	ssk := rsa.PrivateKey(*sk)
	return rsa.DecryptOAEP(hash, rand.Reader, &ssk, data, nil)
}

func main() {
	// Create a Task with 5 workers
	r := requester.NewRequester()
	r.Register()

	sk, _ := rsa.GenerateKey(rand.Reader, 2048)
	pk := epk(sk.PublicKey)
	xsk := esk(*sk)
	// Create five workers
	workers := make([]*worker.Worker, numberOfWorkers)
	for i := 0; i < numberOfWorkers; i++ {
		workers[i] = worker.NewWorker()
	}

	// Now register in parallel
	var lock sync.WaitGroup
	for i := 0; i < numberOfWorkers; i++ {
		lock.Add(1)
		go func(i int) {
			defer lock.Done()
			workers[i].Register()
		}(i)
	}
	lock.Wait()

	r.PostTask(numberOfWorkers, rewards, 1, &pk, "Cluster")
	// Now finding the task
	for i := 0; i < numberOfWorkers; i++ {
		lock.Add(1)
		go func(i int) {
			defer lock.Done()
			workers[i].ParticipantTask(r.Task())
		}(i)
	}
	lock.Wait()

	// Now collecting data
	data := make([][]byte, numberOfWorkers)
	for i := 0; i < numberOfWorkers; i++ {
		numeric := mrand.Float64() * 1000.0
		fmt.Printf("%v %v %v\n", workers[i].ID(), numeric, workers[i].Address())
		data[i] = encoder.Float64ToBytes(numeric)
	}
	for i := 0; i < numberOfWorkers; i++ {
		lock.Add(1)
		go func(i int) {
			defer lock.Done()
			workers[i].CollectData(data[i])
			workers[i].SubmitData()
		}(i)
	}
	lock.Wait()
	// Verify the uploaded data
	fmt.Println(r.Rewarding(&xsk))
}
