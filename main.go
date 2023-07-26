package main

import (
	"os"
)

func main() {
	if os.Args[1] == "stack" {

		SizeOfShortLivedVariables()

	} else if os.Args[1] == "heap" {

		SizeOfStructs()

	}
}
