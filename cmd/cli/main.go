package main

import (
	"fmt"

	"github.com/blockfuselabs/evm/types"
)

func main(){
	s := types.NewStack()
	// s.Data = append(s.Data, types.Byte32{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1})
	// s.Data = append(s.Data, types.Byte32{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1})

	// for i := 0; i < 1024; i++  {
	// 	s.Push(types.Uint256{1})
	// }
	s.Push(types.Uint256{1})
	s.Push(types.Uint256{1})
	s.Push(types.Uint256{1,1,1,1})
	s.Pop()
	fmt.Println(s.Data)
}
