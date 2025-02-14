package main

import (
	"fmt"

	"github.com/blockfuselabs/evm/types"
)

func main() {
	evm := types.NewEvmState("1111", []byte{0x60, 0x03, 0x60, 0x02, 0x01, 0x00}, 1000, 100, []uint8{})
	var b types.Byte32               // Create a zero-initialized Byte32
	b[31] = 0xFD               // Set the last byte to 0xFD

	fmt.Printf("Byte32: %x\n", b)
	evm.Run()
	fmt.Printf("Remaining Gas: %d\nCurrent Stack State: %v\n", evm.Gas, evm.Stack.Data)
}
