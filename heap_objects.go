package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	ID         uint8  // 8 bits or 1 byte
	PersonalId int64  // 64 bits or 8 bytes
	Name       string // 128 bits or 16 bytes
	isAdmin    bool   // 8 bits or 1 byte
}

type OptimizedUser struct {
	ID         uint8  // 8 bits or 1 byte
	isAdmin    bool   // 8 bits or 1 byte
	PersonalId int64  // 64 bits or 8 bytes
	Name       string // 128 bits or 16 bytes
}

func SizeOfStructs() {

	// The following user and optimizedUser will be allocated on the heap.
	user := User{
		ID:         1,
		PersonalId: 101,
		Name:       "user 1",
		isAdmin:    false,
	}
	/*
		Expected size of the user: 8 + 64 + 128 + 8 = 208 bits -> 26 bytes

		Padding will add for each data to align properly in the memory. After adding padding actual the memory is shown below.

		Actual size of the user: 8 -> 64 + 64 + 128 + 8 -> 64 = 320 bits -> 40 bytes

		Because we didn't align our data to fit word sizes in user, so padding is added for properly aligning fields in memory.
	*/

	optimizedUser := OptimizedUser{
		ID:         1,
		PersonalId: 101,
		Name:       "user 1",
		isAdmin:    false,
	}
	/*
		We aligned our data to fit word sizes in optimizedUser. Now we can calculate the memory it needs.

		Actual size of the optimizedUser: (8 + 8) -> 64 + 64 + 128 = 256 bits -> 32 bytes
	*/

	fmt.Printf("Size of user: %d bytes\n", unsafe.Sizeof(user))
	fmt.Printf("Size of optimizedUser: %d bytes\n", unsafe.Sizeof(optimizedUser))
}
