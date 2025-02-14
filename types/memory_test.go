package types

import (
	"testing"
	"fmt"
)

func TestMemory(t *testing.T) {
	mem := NewMemory()

	//...Create a Byte32 value...//
	var value Byte32
	copy(value[:], []byte("hello"))

	//....Store the value at offset 32 ....//
	mem.Store(32, value)
	fmt.Println("Stored in memory:", value)


	//... Load the value back from offset 32...//
	loadedValue := mem.Load(32)
	fmt.Println("Loaded from memory:", loadedValue)


	// ... Check if the loaded value matches the stored value...//
	if loadedValue != value {
		t.Errorf("expected %v, got %v", value, loadedValue)
	}

	// ...Test out-of-bounds access ...//
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic on invalid memory location but got none")
		}
	}()
	mem.Load(100) 
}
