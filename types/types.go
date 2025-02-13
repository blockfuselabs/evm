package types

const (
	MaximumDepth uint = 1024
)

type Byte32 [32]byte
type Uint256 [4]uint64

type Stack struct {
	Data []Uint256
}

type Memory struct {
	Data []Byte32
}
type Storage struct {
	Data map[Byte32]Byte32
}
