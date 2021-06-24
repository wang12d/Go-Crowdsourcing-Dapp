package ethereum

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// GetNonce return the pending nonce of the account, which is the nonce of next transaction
func GetNonce(ctx context.Context, client *ethclient.Client, account common.Address) uint64 {
	nonce, err := client.PendingNonceAt(ctx, account)
	fmt.Printf("Nonce at %v: %v\n", account, nonce)
	if err != nil {
		panic(fmt.Errorf("get nonce error: %v", err))
	}
	return nonce
}
