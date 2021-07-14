package cryptograph

// Encryptor is a interface used to describe the
// data encryption, either symmetric or asymmetric
// encryption scheme.
type Encryptor interface {
	EncryptData(data []byte) ([]byte, error)
}
