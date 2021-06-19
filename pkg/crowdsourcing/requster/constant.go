package requester

type State int

const (
	INIT = iota
	PENDING
	FIN
)
