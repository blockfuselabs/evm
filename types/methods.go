package types

import (
	"fmt"
)

func (s *Stack) Push(data int64) {
	if len(s.Data) >= int(MaximumDepth) {
		panic("stack overflow")
	}

	s.Data = append(s.Data, data)
}

func (s *Stack) Pop() int64 {
	lastIndex := len(s.Data) - 1
	lastItem := s.Data[lastIndex]
	fmt.Println("popped: ", lastItem)
	s.Data = append(s.Data[:lastIndex], s.Data[lastIndex+1:]...)

	return lastItem
}

func (m *Memory) Store(offset byte, value Byte32) {
	index := int(offset) / 32

	if index >= len(m.Data) {
		newMemoryData := make([]Byte32, index+1)
		copy(newMemoryData, m.Data)
		m.Data = newMemoryData
	}

	m.Data[index] = value
}

func (m *Memory) Load(offset byte) Byte32 {
	index := int(offset) / 32

	if index > len(m.Data)-1 {
		panic("invalid memory location")
	}
	return m.Data[index]
}

func (s *Storage) Store(key Byte32, value Byte32) {
	s.Data[key] = value
}

func (s *Storage) Load(key Byte32) (Byte32, error) {
	data, err := s.Data[key]
	if !err {
		return Byte32{}, fmt.Errorf("Error: invalid storage key")
	}

	return data, nil
}

// ---------------------------------------------------------
// ------------------- Operations --------------------------
// ---------------------------------------------------------
func (es *State) Add() {
	a := es.Stack.Pop()
	b := es.Stack.Pop()

	result := a + b
	es.Stack.Push(result)
}

func (es *State) Sub() {
	a := es.Stack.Pop()
	b := es.Stack.Pop()

	result := a - b
	es.Stack.Push(result)
}

func (es *State) Mul() {
	a := es.Stack.Pop()
	b := es.Stack.Pop()

	result := a * b
	es.Stack.Push(result)
}

func (es *State) Div() {
	a := es.Stack.Pop()
	b := es.Stack.Pop()

	// handle division by 0
	if b == 0 {es.Stack.Push(0); return}

	result := a / b
	es.Stack.Push(result)
}

// ---------------------------------------------------------
// ------------------- Constructor Methods -----------------
// ---------------------------------------------------------
func NewStack() *Stack {
	return &Stack{
		Data: make([]int64, 0),
	}
}

func NewMemory() *Memory {
	return &Memory{make([]Byte32, 0)}
}

func NewStorage() *Storage {
	return &Storage{make(map[Byte32]Byte32, 0)}
}

func NewEvmState(sender string, program string, gas uint64, value uint64, calldata []uint8) *State {
	return &State{
		Pc: 0,
		Stack: *NewStack(),
		Memory: *NewMemory(),
		Storage: *NewStorage(),
		Sender: sender,
		Program: program,
		Gas: gas,
		Value: value,
		Calldata: calldata,
		Stop_flag: false,
		Revert_flag: false,
		Returndata: make([]uint8, 0),
		Logs: make([]string, 0),
	}
}