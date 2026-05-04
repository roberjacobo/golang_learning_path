package main

import (
	"fmt"
	"runtime"
)

// Overflow on a constant is caught at COMPILE time, not runtime.
// El overflow en una constante se detecta en tiempo de COMPILACIÓN, no en ejecución.

var test uint8
var test_1 uint16

func main() {
	x := 42
	y := 42.32
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	test = 255
	// test = 256 // Error: cannot use 256 (untyped int constant) as uint8 value in assignment (overflows) (compiler NumericOverflow)
	fmt.Printf("%T\n", test)

	test_1 = 65535
	// test_1 = 65536 // Error: cannot use 65536 (untyped int constant) as uint16 value in assignment (overflows) (compiler NumericOverflow)
	fmt.Printf("%T\n", test_1)

	// runtime package exposes info about the running program / environment.
	// El paquete runtime expone información del programa / entorno en ejecución.
	fmt.Println(runtime.GOOS)   // operating system / sistema operativo
	fmt.Println(runtime.GOARCH) // CPU architecture / arquitectura del CPU
}
