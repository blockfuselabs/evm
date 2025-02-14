package types

import (
	"testing"
	"fmt"
)

func TestStorage(t *testing.T) {
	storage := NewStorage()

	// ... Create a Byte32 key and value  ...//
	var key Byte32
	copy(key[:], []byte("key123"))

	var value Byte32
	copy(value[:], []byte("stored_value"))

	// ...Store the value in storage... //
	storage.Store(key, value)
	fmt.Println("Stored in storage:", value)

	//... Load the value back...//
	loadedValue, err := storage.Load(key)
	fmt.Println("Loaded from storage:", loadedValue)


	//... Check if the loaded value matches the stored value...//
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if loadedValue != value {
		t.Errorf("expected %v, got %v", value, loadedValue)
	}

	//... Test loading a non-existent key (should return an error)...//
	var invalidKey Byte32
	copy(invalidKey[:], []byte("invalid"))

	_, err = storage.Load(invalidKey)
	if err == nil {
		t.Errorf("expected error for invalid key but got none")
	}
}
