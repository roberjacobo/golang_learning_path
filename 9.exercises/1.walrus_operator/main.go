package main

import (
	"fmt"
)

// Walrus ":=" only works INSIDE a function, not at package level.
// El walrus ":=" solo funciona DENTRO de una función, no a nivel de paquete.

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)
	fmt.Println("-----")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
