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

	x = 42
	y = "James Bond"
	z = true

	// \t puts a tab between values; Sprintf builds the string instead of printing it.
	// \t pone un tab entre valores; Sprintf arma el string en lugar de imprimirlo.
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}
