package task

type Key [32]byte

type Task struct {
	workerRequired int
	encKey         Key // The key used to encrypted the uploaded data in crowdsourcing
	reward         int
	description    string
}
