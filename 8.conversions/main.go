package main

import (
	"fmt"
)

// ===== TYPE CONVERSIONS =====
// ===== CONVERSIONES DE TIPO =====

// In Go, types do NOT convert by themselves. You must convert them on purpose.
// En Go, los tipos NO se convierten solos. Tienes que convertirlos a propósito.
// Syntax: newType(value) — like int(b) or float64(a).
// Sintaxis: nuevoTipo(valor) — como int(b) o float64(a).

var a int

type dinero int

var b dinero

func main() {
	a = 42
	fmt.Print(a)
	fmt.Printf("%T\n", a)

	b = 1000
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// Without int(b) Go gives: "cannot use b (type dinero) as type int".
	// Sin int(b) Go marca: "cannot use b (type dinero) as type int".
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
