package reward

import (
	"math/big"

	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/task"
)

type Policy interface {
	CalculateRewards(t *task.Task, reward *big.Int, workerID int) *big.Int
}
