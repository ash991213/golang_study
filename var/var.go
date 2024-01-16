package main

import (
	"fmt"
)

func main() {
	shortened_var()
	multiple_var()
	func_scope()
	global_scope()
	var_reassign()
	name_shadowing()
	zero_value()
}

/** Shortened Var */

func shortened_var() {
	golang, javascript := "true", false

	fmt.Println(golang)
	fmt.Println(javascript)
}

/** Multiple Var */

func multiple_var() {
	message, hi, who, goodBye := "Hello, Go!", "Hi!", "Who R U", "GoodBye"

	fmt.Println(message, hi, who, goodBye)
}

/** Func Scope */

func func_scope() {
	var message string
	message = "Hello, Go!"

	var hi string = "Hi!"

	var who = "Who R U"

	goodBye := "GoodBye"

	fmt.Println(message, hi, who, goodBye)

}

/** Global Scope */

var message string

var hi string = "Hi!"

var who = "Who R U"

var goodBye = "GoodBye"

// ! goodBye := "GoodBye" -> Error

func global_scope() {
	message = "Hello, Go!"

	fmt.Println(message, hi, who, goodBye)
}

/** Var Reassign */

func var_reassign() {
	message := "previous message"
	message, golang := "new message", "Golang"

	fmt.Println(message, golang)

	/** Another File */

	// ! message := "Awesome" -> Error
	// ? message = "Awesome" -> True
}

/** Name Shadowing */

func name_shadowing() {
	string := 10000

	fmt.Println(string)

	// !  var message string = "Hello, Go!" -> Error [string (variable of type int) is not a type]
}

/** Zero Value */

func zero_value() {
	var (
		myInt    int
		myFloat  float64
		myString string
		myBool   bool
	)

	fmt.Println(myInt, myFloat, myString, myBool)
}
