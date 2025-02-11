package types

import "fmt"


func (s *Stack) Push(data Uint256) {
	if len(s.Data) >= int(MaximumDepth) {
		panic("stack overflow")
	}

	s.Data = append(s.Data, data)
}

func (s *Stack) Pop() {
	lastIndex := len(s.Data) - 1
	fmt.Println("popped: ",s.Data[lastIndex])
	s.Data = append(s.Data[:lastIndex], s.Data[lastIndex+1:]...)
}

// ------------------- Constructor Methods -----------------

func NewStack() *Stack {
	return &Stack{
		Data: make([]Uint256, 0),
	}
}