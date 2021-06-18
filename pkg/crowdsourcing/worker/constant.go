package worker

type State int

const (
	INIT    State = iota
	PENDING       // PENDING means waiting for registeration
	FIN
)
