package platform

type Role int

const (
	Requester Role = iota
	Worker
)

type Node struct {
	role    Role
	numbers int
}

func NewRequester(workerRequired int) Node {
	return Node{
		role:    Requester,
		numbers: workerRequired,
	}
}

func NewWorker() Node {
	return Node{
		role:    Worker,
		numbers: 1,
	}
}
