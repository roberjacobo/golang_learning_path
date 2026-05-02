package main

import (
	"fmt"
)

var x int
var y string
var z bool

func main() {
	// Quotes around "x" print the literal letter, not the variable.
	// Las comillas en "x" imprimen la letra literal, no la variable.
	fmt.Println("x")
	fmt.Println("y")
	fmt.Println("z")
	fmt.Println("-------------")

	// Common escapes / Escapes comunes: \n new line, \t tab, \" quote, \\ backslash.
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)
}
