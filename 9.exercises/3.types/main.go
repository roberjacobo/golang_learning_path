package main

import (
	"fmt"
)

// Recap: custom type + explicit conversion (folders 7 and 8).
// Repaso: tipo propio + conversión explícita (carpetas 7 y 8).

type Hot_Dog int

var x Hot_Dog
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("The type of X is: %T\n", x)

	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("The type of Y is: %T\n", y)
}
