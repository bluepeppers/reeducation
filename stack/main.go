package stack

// A stack is a data structure that could contain an arbitrary data type. We'll use ints here
// The key thing about a stack is that it's LIFO (Last In, First Out). That means that if I push something onto the stack, then pop, I should get the same value back
type Stack struct {
	// TODO(you): implement this
}

// You might want to set up some stuff before anyone uses the stack? The parameter N here is the maximum size of the stack (the size is defined as the number of items in the stack at any given point)
func NewStack(n int) *Stack {
	return &Stack{}
}

func (s *Stack) Push(v int) {
}

// If the stack is empty, we should return ok = false (i.e. there's nothing here for us to pop, this is not OK)
func (s *Stack) Pop() (v int, ok bool) {
	return 0, false
}

// TODO: write stub and tests for Size() and Peek()
