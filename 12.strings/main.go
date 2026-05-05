package main

import "fmt"

// ===== STRINGS =====
// ===== CADENAS =====

// Two ways to write a string literal:
// Dos formas de escribir un literal de string:
//   "..."  interpreted: escape sequences like \n and \t are processed.
//   "..."  interpretado: las secuencias de escape como \n y \t se procesan.
//   `...`  raw: everything between backticks is taken verbatim, including newlines.
//   `...`  raw: todo lo que está entre backticks se toma literal, incluso los saltos de línea.

func main() {
	s1 := "Hello World"
	s2 := `This is a splitted
  line`

	fmt.Println(s1)
	fmt.Printf("The type of s1 is: %T\n", s1)

	fmt.Println(s2)
	fmt.Printf("The type of s2 is: %T\n", s2)

	fmt.Println("")

	// A string in Go is immutable, but you can convert it to []byte to see
	// (or modify a copy of) the underlying bytes.
	// Un string en Go es inmutable, pero puedes convertirlo a []byte para ver
	// (o modificar una copia de) los bytes que lo forman.
	bs := []byte(s1)
	fmt.Println(bs) // prints the byte values, not the text / imprime los valores de los bytes, no el texto
	fmt.Printf("The type of bs is: %T\n", bs)
}
