package main

// stack node
type node struct {
	value interface{}
	next  *node
}

// stack structure
type Stack struct {
	top   *node
	depth int
}

// create stack
func CreateStack() *Stack {
	return &Stack{
		top:   nil,
		depth: 0,
	}
}

// obtain the value of the top of stack
func (s *Stack) Top() interface{} {
	if s.depth == 0 {
		return nil
	}
	return s.top.value
}


//obtain the depth of stack
func (s *Stack) Depth() int {
	return s.depth
}

// push to stack
func (s *Stack) Push(value interface{}) {
	node := &node{
		value: value,
		next:  s.top,
	}
	s.top = node
	s.depth++
}

// pop the top of stack
func (s *Stack) Pop() interface{} {
	if s.depth == 0 {
		return nil
	}
	ret := s.top.value
	s.top = s.top.next
	s.depth--
	return ret
}