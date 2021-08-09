package metrics

import (
	"fmt"
	"os"
	"time"
)

// TimeCost measure the time since inputs in a function call
func TimeCost(t time.Time, caller string) {
	timeSince := time.Since(t)
	fmt.Fprintf(os.Stderr, "%v costs: %v\n", caller, timeSince)
}
