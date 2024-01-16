package main

import (
	"fmt"
)

func main() {
	const_assign()
	multiple_const()
}

/** Const Assign */

func const_assign() {
	const MAX_SIZE int = 10

	// ! MAX_SIZE = 20 -> Error [cannot assign to MAX_SIZE (neither addressable nor a map index expression)]

	fmt.Println(MAX_SIZE)
}

/** Multiple Const */

func multiple_const() {
	const MAX_SIZE, LANG = 100, "Golang"

	fmt.Println(MAX_SIZE, LANG)

	const (
		MAX_HEIGHT int = 100
		MAX_WIDTH  int = 150
	)

	fmt.Println(MAX_HEIGHT, MAX_WIDTH)
}
