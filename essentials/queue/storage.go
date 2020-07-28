package queue

type Storage interface {
	Enqueue(batchId string, p []byte) error
	Dequeue() (p []byte, err error)
}

type onMemoryStorage struct {
	queue map[string][][]byte
}
