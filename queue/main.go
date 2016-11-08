package queue

// A queue is a data structure that could contain an arbitrary data type. We'll use ints here
// Queues are similar to stacks, with the major difference being that they are FIFO (First In, First Out)
type Queue struct {
}

func NewQueue(n int) *Queue {
	return &Queue{}
}

// This function should be O(1)
func (q *Queue) Push(v int) {

}

// If the queue is empty, we should return ok = false (i.e. there's nothing here for us to pop, this is not OK)
// This function should be O(1)
func (q *Queue) Pop() (v int, ok bool) {
	return 0, false
}

// TODO: write stub and tests for Size() and Peek()
