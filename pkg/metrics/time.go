package metrics

import (
	"fmt"
	"time"
)

// TimeCost measure the time since inputs in a function call
func TimeCost(t time.Time, caller string) {
	timeSince := time.Since(t)
	fmt.Printf("%v costs: %v\n", caller, timeSince)
}
