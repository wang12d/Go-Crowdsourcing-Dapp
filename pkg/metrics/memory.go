package metrics

import (
	"fmt"
	"runtime"
)

const (
	unit = 1024.0
)

func toMegaBytes(bits uint64) float64 {
	return float64(bits) / unit / unit
}


// GetMemoryStatus returns the runtime memory usage of caller
func GetMemoryStatus(caller string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%v. Heap memory: %v MiBs. Stack ory: %v MiBs\n", caller, toMegaBytes(m.Alloc), toMegaBytes(m.StackInuse))
	runtime.GC()
}
