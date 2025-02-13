package types

import (
	"fmt"
)

func (s *Stack) Push(data byte) {
	if len(s.Data) >= int(MaximumDepth) {
		panic("stack overflow")
	}

	s.Data = append(s.Data, data)
}

func (s *Stack) Pop() byte {
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
func (evm *State) Add() {
	a := evm.Stack.Pop()
	b := evm.Stack.Pop()

	fmt.Println("adding...", a, "+", b)
	result := a + b

	evm.Stack.Push(byte(result))

	fmt.Printf("EVM STACK: %v\n", evm.Stack.Data[0])
	evm.GasDec(3)
	fmt.Printf("EVM Gas used for ADD Op: [%d]GAS, remaining %d\n", 3, evm.Gas)
}

func (evm *State) Sub() {
	a := evm.Stack.Pop()
	b := evm.Stack.Pop()

	result := a - b
	evm.Stack.Push(result)
}

func (evm *State) Mul() {
	a := evm.Stack.Pop()
	b := evm.Stack.Pop()

	result := a * b
	evm.Stack.Push(result)
}

func (evm *State) Div() {
	a := evm.Stack.Pop()
	b := evm.Stack.Pop()

	// handle division by 0
	if b == 0 {
		evm.Stack.Push(0)
		return
	}

	result := a / b
	evm.Stack.Push(result)
}

// ---------------------------------------------------------
func (evm *State) Peek() uint8 {
	return evm.Program[evm.Pc]
}

func (evm *State) GasDec(amount uint64) {
	if evm.Gas < amount {
		panic("Insufficient Gas for execution")
	} else {
		evm.Gas -= amount
	}
}

func (evm *State) Stop() {
	evm.StopFlag = true
}

func (evm *State) ExecuteOpcode() bool {
	// stop opcode execution if stopflag or revertFlag is true
	if evm.StopFlag || evm.RevertFlag {
		return false
	}

	// stop opcode execution if stopflag or revertFlag is true
	if evm.Pc >= uint8(len(evm.Program)) {
		return false
	}

	return true
}

func (evm *State) Reset() {
	evm.Pc = 0
	evm.Stack = *NewStack()
	evm.Memory = *NewMemory()
	evm.Storage = *NewStorage()
}

func (evm *State) Run() {
	for evm.ExecuteOpcode() {
		op := evm.Program[evm.Pc]
		switch op {
		case ADD:
			evm.Add()
			evm.Pc++
			continue
		case PUSH1:
			nextOp := evm.Program[evm.Pc+1]
			evm.Stack.Push(nextOp)
			evm.Pc += 2
			continue
		case STOP:
			evm.Stop()
			break
		default:
			fmt.Println("invalid opcode", op)
		}
	}
}

// ---------------------------------------------------------
// ------------------- Constructor Methods -----------------
// ---------------------------------------------------------
func NewStack() *Stack {
	return &Stack{
		Data: make([]byte, 0),
	}
}

func NewMemory() *Memory {
	return &Memory{make([]Byte32, 0)}
}

func NewStorage() *Storage {
	return &Storage{make(map[Byte32]Byte32, 0)}
}

func NewEvmState(sender string, program []byte, gas uint64, value uint64, calldata []uint8) *State {
	return &State{
		Pc:         0,
		Stack:      *NewStack(),
		Memory:     *NewMemory(),
		Storage:    *NewStorage(),
		Sender:     sender,
		Program:    program,
		Gas:        gas,
		Value:      value,
		Calldata:   calldata,
		StopFlag:   false,
		RevertFlag: false,
		Returndata: make([]uint8, 0),
		Logs:       make([]string, 0),
	}
}
