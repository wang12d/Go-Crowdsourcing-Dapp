package cryptograph

// Decryptor is a interface used to describe the
// data decryption behavior, either symmetric or asymmetric
// encryption scheme.
type Decryptor interface {
	DecryptData(data []byte) ([]byte, error)
}
