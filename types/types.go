package types

const (
	MaximumDepth uint = 1024
)

type Byte32 [32]byte
type Uint256 [4]uint64

type Stack struct {
	Data []byte
}

type Memory struct {
	Data []Byte32
}
type Storage struct {
	Data map[Byte32]Byte32
}

type State struct {
	Pc uint8
	Stack Stack
	Memory Memory
	Storage Storage
	Sender string
	Program []byte
	Gas uint64
	Value uint64
	Calldata []uint8
	StopFlag bool
	RevertFlag bool
	Returndata []uint8
	Logs []string
}
