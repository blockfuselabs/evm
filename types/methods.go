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
	fmt.Println("popped: ", s.Data[lastIndex])
	s.Data = append(s.Data[:lastIndex], s.Data[lastIndex+1:]...)
}

func (m *Memory) Store(offset byte, value Byte32) {
	index := int(offset) / 32

	if index >= len(m.Data) {
		newMemoryData := make([]Byte32, index + 1)
		copy(newMemoryData, m.Data)
		m.Data = newMemoryData
	}

	m.Data[index]=value
}

func (m *Memory) Load(offset byte) Byte32 {
	index := int(offset) / 32
	// fmt.Println(index > len(m.Data)-1)
	if index > len(m.Data)-1 {
		panic("invalid memory location")
	}
	return m.Data[index]
}

// ------------------- Constructor Methods -----------------

func NewStack() *Stack {
	return &Stack{
		Data: make([]Uint256, 0),
	}
}

func NewMemory() *Memory {
	return &Memory{make([]Byte32, 0)}
}
