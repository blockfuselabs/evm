package main

import (
	// "fmt"

	"github.com/blockfuselabs/evm/types"
)

func main() {
	// s := types.NewStack()
	es := types.NewEvmState("1111", "1111", 1000, 100, []uint8{})
	// m := types.NewMemory()
	// s.Data = append(s.Data, types.Byte32{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1})
	// s.Data = append(s.Data, types.Byte32{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1})

	// for i := 0; i < 1024; i++  {
	// 	s.Push(types.Uint256{1})
	// }

	es.Stack.Push(3)
	es.Stack.Push(8)
	es.Div()
	// fmt.Println(s.Data)
	// var offset96 byte = 0x60
	// var offset32 byte = 0x20
	// var offset00 byte = 0x00
	// m.Store(offset00, types.Byte32{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
	// m.Load(offset00)
	// m.Store(offset96, types.Byte32{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
	// m.Load(offset96)
	// m.Store(offset32, types.Byte32{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
	// m.Load(offset32)
	// fmt.Println(m.Data)

	// s := types.NewStorage()

	// s.Store(types.Byte32{1}, types.Byte32{1})
	// s.Store(types.Byte32{2}, types.Byte32{12})
	// s.Store(types.Byte32{2}, types.Byte32{3, 3, 5})
	// data, err := s.Load(types.Byte32{4})
	// fmt.Printf("%+v\nError:%s\n", data[0], err)
}
