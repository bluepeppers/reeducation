package stack

import (
	"testing"
)

// Tests that when we put items onto the stack, we can get them back, in the correct (reverse) order
func TestPushThenPop(t *testing.T) {
	stack := NewStack(16)

	values := []int{9, 8, 7, 6}

	// We're going to push these onto the stack
	for _, v := range values {
		stack.Push(v)
	}

	// When we pop them, we should get them in reverse order
	for i := range values {
		got, ok := stack.Pop()
		// Ok will be false if there is nothing left on the stack, which is not the case
		if !ok {
			t.Fatalf("Could not pop %vth item from stack", i)
		}
		// This should be 6 on the first iteration (values[4 - 1 - 0])
		expected := values[len(values)-1-i]
		if got != expected {
			t.Fatalf("Got unexpected value from stack:\nGot: %v\nExpected: %v\n", got, expected)
		}
	}
}

// Tests that we don't get anything out of the stack when we pop it while empty
func TestPopEmpty(t *testing.T) {
	stack := NewStack(0)

	got, ok := stack.Pop()
	if ok {
		t.Fatalf("Could pop from empty stack:\nGot: %v\n", got)
	}
}

// Tests that we can push and pop as we like
func TestPushPopPushPop(t *testing.T) {
	stack := NewStack(3)

	var ok bool
	mustBeOK := func(ok bool, i int) {
		if !ok {
			t.Fatalf("%vth pop failed", i)
		}
	}
	expected := []int{8, 9, 5, 4, 6, 7}
	got := make([]int, len(expected))

	// Here we push and pop stuff. If you think of the Y axis here as time, and the X as the contents of the stack, this example looks to me like
	//
	//
	//              5   4
	//   8      6 6 6 6 6 6
	// 9 9 9  7 7 7 7 7 7 7 7

	stack.Push(9)
	stack.Push(8)
	got[0], ok = stack.Pop()
	mustBeOK(ok, 0)
	got[1], ok = stack.Pop()
	mustBeOK(ok, 1)

	stack.Push(7)
	stack.Push(6)
	stack.Push(5)
	got[2], ok = stack.Pop()
	mustBeOK(ok, 2)
	stack.Push(4)
	got[3], ok = stack.Pop()
	mustBeOK(ok, 3)
	got[4], ok = stack.Pop()
	mustBeOK(ok, 4)
	got[5], ok = stack.Pop()
	mustBeOK(ok, 5)

	for i := range expected {
		if expected[i] != got[i] {
			t.Fatalf("%vth pop result differed:\nGot: %v\nExpected: %v\n", i, got[i], expected[i])
		}
	}
}
