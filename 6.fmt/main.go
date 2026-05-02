package main

import (
	"fmt"
)

// Package-level variables (declared outside any function)
// Variables a nivel de paquete (declaradas fuera de cualquier función)

// int with no value: defaults to 0
// int sin valor: por defecto es 0
var a int

// string with a value
// string con un valor
var b string = "Programa"

// bool with no value: defaults to false
// bool sin valor: por defecto es false
var c bool

// bool with value true
// bool con valor true
var d bool = true

func main() {
	// Short declaration ":=" infers the type from the value
	// Declaración corta ":=" infiere el tipo a partir del valor
	e := 42
	f := "dice hola mundo."

	// Backticks `` make a raw string: keeps line breaks and ignores escapes
	// Las comillas inversas `` crean un raw string: conserva saltos de línea e ignora escapes
	g := `El doctor dice que comer vegetales es
	saludable.`

	// Println: prints values and adds a new line at the end
	// Println: imprime valores y agrega un salto de línea al final
	fmt.Println("Hello, playground")
	fmt.Println("----------")

	fmt.Println(a) // value 0 by default / valor 0 por defecto
	a = 22
	fmt.Println(a)
	fmt.Println("----------")

	// Printf: formatted print using verbs like %v and %T
	// Printf: impresión con formato usando verbos como %v y %T
	// %v = value (default format) / valor (formato por defecto)
	// %T = type of the variable / tipo de la variable
	// \n = new line / nueva línea
	fmt.Printf("El valor de la variable a es: %v\n", a)
	fmt.Printf("El tipo de la variable a es: %T\n", a)
	fmt.Println("----------")

	fmt.Printf("El valor de la variable b es: %v\n", b)
	fmt.Printf("El tipo de la variable b es: %T\n", b)
	fmt.Println("----------")

	// Sprint: builds a string instead of printing it (returns the result)
	// Sprint: construye un string en lugar de imprimirlo (devuelve el resultado)
	s1 := fmt.Sprint("El ", b, " dice hola mundo.")
	fmt.Println(s1)
	fmt.Printf("%T\n", s1) // shows the type of s1 / muestra el tipo de s1
	fmt.Println("----------")

	// Print each variable to see its value
	// Imprime cada variable para ver su valor
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	// Println with multiple arguments adds spaces between them
	// Println con varios argumentos agrega espacios entre ellos
	fmt.Println("Mi", b, f)

	// Raw string keeps the line break from g
	// El raw string conserva el salto de línea de g
	fmt.Println(g)
}
