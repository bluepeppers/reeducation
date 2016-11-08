package queue

import (
	"testing"
)

// Tests that when we put items onto the queue, we can get them back, in the correct order
func TestPushThenPop(t *testing.T) {
	queue := NewQueue(16)

	values := []int{9, 8, 7, 6}

	// We're going to push these onto the queue
	for _, v := range values {
		queue.Push(v)
	}

	// When we pop them, we should get them in reverse order
	for i := range values {
		got, ok := queue.Pop()
		// Ok will be false if there is nothing left on the queue, which is not the case
		if !ok {
			t.Fatalf("Could not pop %vth item from queue", i)
		}
		expected := values[i]
		if got != expected {
			t.Fatalf("Got unexpected value from queue:\nGot: %v\nExpected: %v\n", got, expected)
		}
	}
}

// Tests that we don't get anything out of the queue when we pop it while empty
func TestPopEmpty(t *testing.T) {
	queue := NewQueue(0)

	got, ok := queue.Pop()
	if ok {
		t.Fatalf("Could pop from empty queue:\nGot: %v\n", got)
	}
}

// Tests that we can push and pop as we like
func TestPushPopPushPop(t *testing.T) {
	queue := NewQueue(3)

	var ok bool
	mustBeOK := func(ok bool, i int) {
		if !ok {
			t.Fatalf("%vth pop failed", i)
		}
	}
	expected := []int{9, 8, 7, 6, 5, 4}
	got := make([]int, len(expected))

	// Here we push and pop stuff. If you think of the Y axis here as time, and the X as the contents of the queue, this example looks to me like
	//
	//
	//              5   4
	//   8      6 6 6 5 5 4
	// 9 9 8  7 7 7 7 6 6 5 4

	queue.Push(9)
	queue.Push(8)
	got[0], ok = queue.Pop()
	mustBeOK(ok, 0)
	got[1], ok = queue.Pop()
	mustBeOK(ok, 1)

	queue.Push(7)
	queue.Push(6)
	queue.Push(5)
	got[2], ok = queue.Pop()
	mustBeOK(ok, 2)
	queue.Push(4)
	got[3], ok = queue.Pop()
	mustBeOK(ok, 3)
	got[4], ok = queue.Pop()
	mustBeOK(ok, 4)
	got[5], ok = queue.Pop()
	mustBeOK(ok, 5)

	for i := range expected {
		if expected[i] != got[i] {
			t.Fatalf("%vth pop result differed:\nGot: %v\nExpected: %v\n", i, got[i], expected[i])
		}
	}
}
