package types

const (
	MaximumDepth uint = 1024
)

type Byte32 [32]byte
type Uint256 [4]uint64

type Stack struct {
	Data []int64
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
	Program string
	Gas uint64
	Value uint64
	Calldata []uint8
	Stop_flag bool
	Revert_flag bool
	Returndata []uint8
	Logs []string
}
