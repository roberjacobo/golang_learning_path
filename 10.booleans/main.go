package main

import (
	"fmt"
)

// Comparison operators (==, !=, <, >, <=, >=) always return a bool.
// Los operadores de comparación (==, !=, <, >, <=, >=) siempre devuelven un bool.

func main() {
	x := 7
	y := 42

	fmt.Println(x == y) // false
	fmt.Println(x != y) // true
	fmt.Println(x < y)  // true
}
