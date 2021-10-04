package encoder

import (
	"encoding/binary"
)

func Uint64FromBytes(bytes []byte) uint64 {
	return binary.LittleEndian.Uint64(bytes)
}

func Uint64ToBytes(intVal uint64) []byte {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, intVal)
	return bytes
}
