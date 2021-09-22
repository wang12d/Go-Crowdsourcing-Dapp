package main

import (
	"fmt"
	"math/big"
	mrand "math/rand"
	"sync"

	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/requester"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/task"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/utils/encoder"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/worker"
)

const (
	numberOfWorkers = 5
	rewards         = 500000000
)

type PC struct {
}

func (cp *PC) CalculateRewards(t *task.Task, reward *big.Int, workerID int) *big.Int {
	return reward
}

func main() {
	// Create a Task with 5 workers
	r := requester.NewRequester(2048)
	r.Register()

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
			workers[i].Register(1)
		}(i)
	}
	lock.Wait()

	r.PostTask(numberOfWorkers, rewards, "Cluster")
	fmt.Printf("task address: %v\n", r.Task().Address().Hex())
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
		numeric := mrand.Uint64() % 1000
		fmt.Printf("%v %v %v\n", workers[i].ID(), numeric, workers[i].Address())
		data[i] = encoder.Uint64ToBytes(numeric)
	}
	for i := 0; i < numberOfWorkers; i++ {
		lock.Add(1)
		go func(i int) {
			defer lock.Done()
			workers[i].CollectData(0, data[i])
			workers[i].SubmitData(0)
		}(i)
	}
	lock.Wait()
	// Verify the uploaded data
	fmt.Println(r.Rewarding(&PC{}))

	// Verify workers can participant multiple tasks
	r = requester.NewRequester(2048)
	r.Register()

	workers = make([]*worker.Worker, numberOfWorkers)
	for i := 0; i < numberOfWorkers; i++ {
		workers[i] = worker.NewWorker()
	}

	for i := 0; i < numberOfWorkers; i++ {
		lock.Add(1)
		go func(i int) {
			defer lock.Done()
			workers[i].Register(1)
		}(i)
	}
	lock.Wait()

	r.PostTask(numberOfWorkers, rewards, "Cluster")
	fmt.Printf("task address: %v\n", r.Task().Address().Hex())
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
	data = make([][]byte, numberOfWorkers)
	for i := 0; i < numberOfWorkers; i++ {
		numeric := mrand.Uint64() % 1000
		fmt.Printf("%v %v %v\n", workers[i].ID(), numeric, workers[i].Address())
		data[i] = encoder.Uint64ToBytes(numeric)
	}
	for i := 0; i < numberOfWorkers; i++ {
		lock.Add(1)
		go func(i int) {
			defer lock.Done()
			workers[i].CollectData(0, data[i])
			workers[i].SubmitData(0)
		}(i)
	}
	lock.Wait()
	// Verify the uploaded data
	fmt.Println(r.Rewarding(&PC{}))

	pk, vk := r.GeneateZKProof()
	fmt.Printf("%v\t%v\n", len(pk), len(vk))
}
