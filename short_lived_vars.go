package main

import (
	"fmt"
	"unsafe"
)

func SizeOfShortLivedVariables() {

	// The following variables are short lived variables which uses memory allocated on the stack.
	id := 1
	uId := uint8(1)
	string := "string"
	boolean := true

	fmt.Printf("Size of %T(%d): %d bytes\n", id, id, unsafe.Sizeof(id))
	fmt.Printf("Size of %T(%d): %d bytes\n", uId, uId, unsafe.Sizeof(uId))
	fmt.Printf("Size of %s: %d bytes\n", string, unsafe.Sizeof(string))
	fmt.Printf("Size of %T(%t): %d bytes\n", boolean, boolean, unsafe.Sizeof(boolean))
}
