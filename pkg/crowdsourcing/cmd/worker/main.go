package main

import (
	"fmt"
	"sync"

	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/plantform"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/task"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/worker"
)

const (
	numberOfWorkers = 5
	rewards         = 1
)

func main() {
	// Create a Task with 5 workers
	t := task.NewTask(numberOfWorkers, rewards, []byte(fmt.Sprintf("%032v", "a")), "Cluster")
	plantform.CP.ReceiveTask(t)

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

	// Now finding the task
	for i := 0; i < numberOfWorkers; i++ {
		lock.Add(1)
		go func(i int) {
			defer lock.Done()
			workers[i].ParticipantTask()
		}(i)
	}
	lock.Wait()

	// Now collecting data
	data := make([][]byte, numberOfWorkers)
	for i := 0; i < numberOfWorkers; i++ {
		data[i] = []byte(fmt.Sprintf("%016x", i))
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
	for _, data := range plantform.CP.CheckData(t) {
		fmt.Println(string(data))
	}
}
