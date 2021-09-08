package zksnark

import "github.com/wang12d/GoMarlin/marlin"

type ZK_Gen interface {
	GenerateProofAndVerifyKey() (marlin.Proof, marlin.VerifyKey)
}

type ZK_Verify interface {
	VerifyProof(proof marlin.Proof, vk marlin.VerifyKey) bool
}
