package requester

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/plantform"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/task"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/utils/ethereum"
)

type Requester struct {
	address    common.Address
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
	state      State
	task       *task.Task
}

// NewRequester returns a new requester
func NewRequester() *Requester {
	return &Requester{
		address:    common.Address{},
		privateKey: nil,
		publicKey:  nil,
		state:      INIT,
		task:       nil,
	}
}

// Register registe the worker into the Crowdsourcing plantform
func (r *Requester) Register() {
	if r.state != INIT { // Only worker at INIT can register
		return
	}
	privateKey, address := ethereum.PrivateKeyAndAddress(plantform.CP.NewAccount())
	r.privateKey, r.address = privateKey, address
	r.publicKey = &privateKey.PublicKey
	r.state = PENDING
}

func (r *Requester) CreateTask(workers, reward int, encKey []byte, description string) {
	if r.state != PENDING {
		return
	}
	r.task = task.NewTask(workers, reward, encKey, description)
}

func (r *Requester) PostTask() {
	plantform.CP.ReceiveTask(r.task)
}
