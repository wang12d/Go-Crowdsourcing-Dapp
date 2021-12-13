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
	"os"
	"strconv"
	"time"

	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/requester"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/task"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/utils/encoder"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/worker"
	"github.com/wang12d/GoMarlin/marlin"
)

const (
	reward = 5000
	mu     = 1000
	sigma  = 250
)

var ()

func main() {

	numberOfIteration := 5
	byteSize, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Convert to number error: %v\n", err)
	}

	numberOfWorkers := make([]int, 0)
	for _, str := range os.Args[2:] {
		nw, err := strconv.Atoi(str)
		if err != nil {
			log.Fatalf("Convert to number error: %v\n", err)
		}
		numberOfWorkers = append(numberOfWorkers, nw)
	}
	length := len(numberOfWorkers)
	registrationTimeCost, taskPublicationTimeCost := make([]time.Duration, length), make([]time.Duration, length)
	requesterRegistrationTimeCost, workerRegistrationTimeCost := make([]time.Duration, length), make([]time.Duration, length)
	taskParticipationTimeCost, dataCollectionTimeCost := make([]time.Duration, length), make([]time.Duration, length)
	rewardingTimeCost := make([]time.Duration, length)
	onChainStorage, communication := make([]int, length), make([]int, length)

	sk, _ := rsa.GenerateKey(crand.Reader, byteSize)

	var registrationCost, taskPublicationCost, taskParticipationCost, dataCollectionCost, rewardingCost time.Duration
	var requesterRegistrationCost, workerRegistrationCost time.Duration
	var onChainBytes int
	var communicationBytes int

	var timeStart time.Time
	for i := 0; i < int(numberOfIteration); i++ {
		for n, workerRequired := range numberOfWorkers {
			onChainBytes, registrationCost, taskPublicationCost, taskParticipationCost, dataCollectionCost, rewardingCost = 0, 0, 0, 0, 0, 0
			requesterRegistrationCost, workerRegistrationCost = 0, 0
			communicationBytes = 0
			r := requester.NewRequester(byteSize)
			timeStart = time.Now()
			r.Register()
			requesterRegistrationCost += time.Since(timeStart)
			registrationCost += requesterRegistrationCost
			requesterRegistrationTimeCost[n] += requesterRegistrationCost
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
				workerRegistrationCost += time.Since(timeStart)
			}
			registrationCost += workerRegistrationCost
			workerRegistrationTimeCost[n] += workerRegistrationCost / time.Duration(workerRequired)
			registrationTimeCost[n] += registrationCost / time.Duration(workerRequired+1)

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
			rewardingCost += time.Since(timeStart)
			rewardingTimeCost[n] += rewardingCost

			communicationBytes += onChainBytes / workerRequired
			onChainStorage[n] += onChainBytes
			communication[n] += communicationBytes
		}
	}

	fmt.Printf("CrowdBC: %v\n", byteSize)
	for i := 0; i < length; i++ {
		fmt.Printf("%v(%v, %v),%v,%v,%v,%v,%v,%v\n", registrationTimeCost[i]/time.Duration(numberOfIteration),
			requesterRegistrationTimeCost[i]/time.Duration(numberOfIteration), workerRegistrationTimeCost[i]/time.Duration(numberOfIteration),
			taskPublicationTimeCost[i]/time.Duration(numberOfIteration), taskParticipationTimeCost[i]/time.Duration(numberOfIteration),
			dataCollectionTimeCost[i]/time.Duration(numberOfIteration), rewardingTimeCost[i]/time.Duration(numberOfIteration),
			float64(onChainStorage[i])/1024.0/float64(numberOfIteration), float64(communication[i])/float64(numberOfIteration)/1024.0,
		)
	}
}

type PC struct {
}

func (cp *PC) CalculateRewards(t *task.Task, reward *big.Int, workerID int) *big.Int {
	return reward
}
