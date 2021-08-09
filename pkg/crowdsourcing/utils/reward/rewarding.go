package reward

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/task"
)

type Policy interface {
	CalculateRewards(data []byte, t *task.Task, workerAddress common.Address) *big.Int
}
