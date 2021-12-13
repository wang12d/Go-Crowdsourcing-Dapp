package requester

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/sha512"
	"log"
	"math"
	"math/big"
	"time"

	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/client"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/utils/cryptograph"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/utils/reward"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/metrics"
	"github.com/wang12d/GoMarlin/marlin"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/platform"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/smartcontract/ctask"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/task"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/utils/encoder"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/utils/ethereum"
)

const (
	mean  = 1000.0
	sigma = 250.0
	EPS   = 1e-8
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

type requesterKeyPair struct {
	esk *rsa.PrivateKey
	epk *rsa.PublicKey
}

func newKeyPair(dataSize int) *requesterKeyPair {
	esk, epk := cryptograph.GenerateRsaKeyPair(dataSize)
	return &requesterKeyPair{
		esk: esk,
		epk: epk,
	}
}

func (rkp *requesterKeyPair) EncryptData(data []byte) ([]byte, error) {
	hash := sha512.New()
	return rsa.EncryptOAEP(hash, cryptograph.Dummy{}, rkp.epk, data, nil)
}

func (rkp *requesterKeyPair) DecryptData(data []byte) ([]byte, error) {
	hash := sha512.New()
	return rsa.DecryptOAEP(hash, cryptograph.Dummy{}, rkp.esk, data, nil)
}

func (rkp *requesterKeyPair) GetESK() *rsa.PrivateKey {
	return rkp.esk
}

func (rkp *requesterKeyPair) GetEPK() *rsa.PublicKey {
	return rkp.epk
}

type Requester struct {
	address    common.Address
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
	state      State
	task       *task.Task
	opts       *bind.TransactOpts
	plainData  []uint // Plain data collected by crowdsourcing workers
	kp         *requesterKeyPair
	rawData    [][]byte // Raw data collected by crowdsourcing workers
}

// NewRequester returns a new requester
func NewRequester(keySize int) *Requester {
	return &Requester{
		address:    common.Address{},
		privateKey: nil,
		publicKey:  nil,
		state:      INIT,
		task:       nil,
		opts:       nil,
		plainData:  make([]uint, 0),
		kp:         newKeyPair(keySize),
		rawData:    make([][]byte, 0),
	}
}

// evaluation is the data quality evaluation function
func (r *Requester) evaluation(data []byte) float64 {
	numeric := encoder.Uint64FromBytes(data)
	return float64(numeric) - mean
}

func (r *Requester) qualityThreshold(q float64) bool {
	return q-3*sigma <= EPS && EPS <= 3*sigma+q
}

// Register register the worker into the Crowdsourcing platform
func (r *Requester) Register() {
	caller := metrics.GetCallerName()
	defer metrics.GetMemoryStatus(caller)
	defer metrics.TimeCost(time.Now(), caller)
	if r.state != INIT { // Only worker at INIT can register
		return
	}
	privateKey, address := platform.CP.NewAccount()
	r.privateKey, r.address = privateKey, address
	r.publicKey = &privateKey.PublicKey
	r.state = PENDING
	r.opts = ethereum.KeyedTransactor(client.CLIENT, r.privateKey,
		r.address, platform.CP.ChainID(), big.NewInt(0))
	platform.CP.RegisterRequester(r.address)
}

// PostTask create and post the task to platform
func (r *Requester) PostTask(workers int, reward int64, description string) {
	caller := metrics.GetCallerName()
	defer metrics.GetMemoryStatus(caller)
	defer metrics.TimeCost(time.Now(), caller)
	// Before post task, requester must deposit the corresponding collaterals
	if r.state != PENDING {
		return
	}
	_workers, _rewards := big.NewInt(int64(workers)), big.NewInt(int64(reward))

	collateral := big.NewInt(_rewards.Int64())
	// Now publishing the task to blockchain
	taskAddress, _, taskContract, err := ctask.DeployCtask(r.opts, client.CLIENT, _workers, _rewards, description)
	if err != nil {
		log.Fatalf("Publish task error: %v\n", err)
	}
	ethereum.UpdateNonce(client.CLIENT, r.opts, r.address)
	if err := ethereum.DepositCollateral(client.CLIENT, r.privateKey, r.address, platform.CP.InstanceAddress(), collateral, []byte{0x01}); err != nil {
		log.Fatalf("Requester deposite collaterals error: %v\n", err)
	}
	ethereum.UpdateNonce(client.CLIENT, r.opts, r.address)
	r.task = task.NewTask(big.NewInt(int64(workers)), big.NewInt(int64(reward)),
		r.kp, taskAddress, description, r.evaluation, taskContract)
	r.state = WAITING
	platform.CP.AddingTasks(r.task)
}

// Task returns the task of the requester
func (r *Requester) Task() *task.Task {
	return r.task
}

// Rewarding returns a list of rewards each worker will received
func (r *Requester) Rewarding(rewardingPolicy reward.Policy) []*big.Int {
	caller := metrics.GetCallerName()
	defer metrics.GetMemoryStatus(caller)
	defer metrics.TimeCost(time.Now(), caller)
	rewardList := make([]*big.Int, r.task.WorkerRequired().Int64())
	qualityList := make([]uint64, r.task.WorkerRequired().Int64())
	data := r.task.Data()
	var avgQuality float64 = 0.0
	var qualityCnt uint64 = 0
	for i, encData := range data {
		rawData, err := r.kp.DecryptData(encData)
		if err != nil {
			log.Fatalf("Rewarding error when decrypting: %v\n", err)
		}
		r.rawData = append(r.rawData, rawData)
		dataQuality := r.evaluation(rawData)
		qualityList[i] = encoder.Uint64FromBytes(rawData)
		r.plainData = append(r.plainData, uint(qualityList[i]))
		if !r.qualityThreshold(dataQuality) {
			qualityList[i] = 0
		} else {
			avgQuality += float64(qualityList[i])
			qualityCnt += 1
		}
	}
	maximumReward := big.NewInt(r.task.Reward().Int64())
	if qualityCnt > 0 {
		maximumReward.Div(maximumReward, big.NewInt(int64(qualityCnt)))
	}
	avgQuality = avgQuality / float64(qualityCnt)
	for i := range data {
		var reward int64
		if qualityList[i] != 0 {
			reward = int64(math.Ceil(math.Min(float64(qualityList[i])/avgQuality, avgQuality/float64(qualityList[i])) * float64(maximumReward.Int64())))
		} else {
			reward = 0
		}
		realReward := rewardingPolicy.CalculateRewards(r.task, big.NewInt(reward), i)
		rewardList[i] = realReward
		if _, err := platform.CP.Instance().Rewarding(r.opts, r.task.WorkerAddresses()[i], realReward, big.NewInt(reward), r.task.Address()); err != nil {
			log.Fatalf("Rewarding %v error: %v\n", i, err)
		}
		ethereum.UpdateNonce(client.CLIENT, r.opts, r.address)
	}
	return rewardList
}

// GenerateZKProof generates the zk snark proof to protect the privacy of each inputs
func (r *Requester) GeneateZKProof() (marlin.Proof, marlin.VerifyKey) {
	publicKeyPem, err := cryptograph.ExportRsaPublicKeyAsPem(r.kp.epk)
	if err != nil {
		log.Fatalf("Requester export public key error: %v\n", err)
	}
	return marlin.GenerateEncryptionZKProofAndVerifyKey(uint(mean),
		uint(sigma), r.plainData, publicKeyPem,
		cryptograph.ExportRsaPrivateKeyAsPem(r.kp.esk), r.task.Data(), r.rawData, uint(r.kp.epk.N.BitLen()),
	)
}

// Address return the blockchain address of the requester
func (r *Requester) Address() common.Address {
	return r.address
}
