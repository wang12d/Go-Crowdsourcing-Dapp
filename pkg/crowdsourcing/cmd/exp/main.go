package main

import (
	"crypto"
	crand "crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"time"

	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/requester"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/task"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/utils/encoder"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/worker"
	"github.com/wang12d/GoMarlin/marlin"
)

const (
	numberOfIteration = uint64(10)
	reward            = 5000
	mu                = 1000
	sigma             = 250
)

var (
	numberOfWorkers = []int{
		3, 5, 7, 9, 11,
	}
)

func main() {
	byteSize := 2048

	length := len(numberOfWorkers)
	registerationTimeCost, taskPublicationTimeCost := make([]time.Duration, length), make([]time.Duration, length)
	taskParticipationTimeCost, dataCollectionTimeCost := make([]time.Duration, length), make([]time.Duration, length)
	rewardingTimeCost := make([]time.Duration, length)
	onChainStorage := make([]int, length)

	sk, _ := rsa.GenerateKey(crand.Reader, byteSize)

	var registerationCost, taskPublicationCost, taskParticipationCost, dataCollectionCost, rewardingCost time.Duration
	var onChainBytes int

	var timeStart time.Time
	for i := 0; i < int(numberOfIteration); i++ {
		for n, workerRequired := range numberOfWorkers {
			onChainBytes, registerationCost, taskPublicationCost, taskParticipationCost, dataCollectionCost, rewardingCost = 0, 0, 0, 0, 0, 0
			r := requester.NewRequester(byteSize)
			timeStart = time.Now()
			r.Register()
			registerationCost += time.Since(timeStart)
			taskDescription := "Collecting the time of daliy smartphone usage"
			timeStart = time.Now()
			r.PostTask(workerRequired, reward, taskDescription)
			taskPublicationCost += time.Since(timeStart)

			taskPublicationTimeCost[n] += taskPublicationCost

			workers := make([]*worker.Worker, workerRequired)
			for ii := 0; ii < workerRequired; ii++ {
				workers[ii] = worker.NewWorker()
				timeStart = time.Now()
				workers[ii].Register(1)
				registerationCost += time.Since(timeStart)
			}
			registerationTimeCost[n] += registerationCost / time.Duration(workerRequired+1)

			evalResults := make([]marlin.EvaluationResults, workerRequired)
			for ii, w := range workers {
				daliyTime := mu + 3*sigma + rand.Uint64()%5000
				data := encoder.Uint64ToBytes(daliyTime)

				timeStart = time.Now()
				w.ParticipantTask(r.Task())
				taskParticipationCost += time.Since(timeStart)

				timeStart = time.Now()
				w.CollectData(0, data)
				w.SubmitData(0)
				dataCollectionCost += time.Since(timeStart)
				evalResults[workerRequired-ii-1] = marlin.EvaluationResults{
					uint64(daliyTime - mu + 3*sigma),
					uint64(daliyTime - mu - 3*sigma),
				}
			}

			encryptedData := r.Task().Data()
			for _, d := range encryptedData {
				hashed := sha256.Sum256(d)
				signature, err := rsa.SignPKCS1v15(crand.Reader, sk, crypto.SHA256, hashed[:])
				if err != nil {
					log.Fatalf("Obtain signature error: %v\n", err)
				}
				onChainBytes += len(d)
				onChainBytes += len(signature)
			}

			taskParticipationTimeCost[n] += taskParticipationCost / time.Duration(workerRequired)
			dataCollectionTimeCost[n] += dataCollectionCost / time.Duration(workerRequired)

			timeStart = time.Now()
			r.Rewarding(&PC{})
			rewardingCost = time.Since(timeStart)
			rewardingTimeCost[n] += rewardingCost

			onChainStorage[n] += onChainBytes
		}
	}

	for i := 0; i < length; i++ {
		fmt.Printf("%v,%v,%v,%v,%v,%v\n", registerationTimeCost[i]/time.Duration(numberOfIteration),
			taskPublicationTimeCost[i]/time.Duration(numberOfIteration), taskParticipationTimeCost[i]/time.Duration(numberOfIteration),
			dataCollectionTimeCost[i]/time.Duration(numberOfIteration), rewardingTimeCost[i]/time.Duration(numberOfIteration),
			float64(onChainStorage[i])/1024.0/float64(numberOfIteration),
		)
	}
}

type PC struct {
}

func (cp *PC) CalculateRewards(t *task.Task, reward *big.Int, workerID int) *big.Int {
	return reward
}
