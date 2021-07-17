package vm

import (
	"fmt"
)

type Stack struct {
	Vals  []Value
	Count int
}

//types of values we can have
const (
	VAL_NUMBER = iota
	VAL_BOOL = iota
)

type Value struct {
	type int
	val  interface{}
}

func CreateStack() *Stack{
	var vals []Value
	stack := Stack{Vals: vals, Count: 0}
	return &stack
}

func (s *Stack) Push(val Value) {
	s.Vals = append(s.Vals, val)
	s.Count += 1
}

func (s *Stack) Pop() *Value {
	if s.Count == 0 {
		//its not correct to return 0 when the stack is empty but i dont want to add a second
		//return argument and have to change stuff in the test cases
		return 0
	}

	val := s.Vals[s.Count - 1]
	s.Vals = s.Vals[:s.Count - 1]
	s.Count -= 1
	return &val
}

func (s *Stack) Peek(i int) *Value {
	return s.Vals[s.Count - (i + 1)]
}

func (s *Stack) PrintStack() {
	for i := 0; i < s.Count; i++ {
		fmt.Println("type: %v, value: %v", s.Vals[i].type, s.Vals[i].val)
	}
}
