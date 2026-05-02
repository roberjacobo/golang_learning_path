package main

import "fmt"

// ===== CUSTOM TYPES =====
// ===== TIPOS PROPIOS (PERSONALIZADOS) =====

// In Go you can create your own type from another type.
// En Go puedes crear tu propio tipo a partir de otro tipo.
// This gives the value a clearer name and a stronger meaning.
// Esto le da al valor un nombre más claro y un significado más fuerte.

type Celsius float64
type UserID int

// Udemy
var a int

type Money int

var b Money

func main() {
	var bodyTemp Celsius = 37.0
	var id UserID = 1001

	fmt.Println("body temp:", bodyTemp, "°C")
	fmt.Println("user id:", id)

	// You CANNOT mix Celsius with a plain float64 directly.
	// NO puedes mezclar Celsius con un float64 normal directamente.
	var raw float64 = 25.5
	// var wrong Celsius = raw      // ❌ error: type mismatch / error: tipos no coinciden
	var roomTemp Celsius = Celsius(raw) // ✅ explicit conversion / conversión explícita
	fmt.Println("room temp:", roomTemp, "°C")
	fmt.Println("----------------")

	// Untyped literal 2000 fits into Money automatically.
	// El literal 2000 sin tipo entra en Money automáticamente.
	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b = 2000
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
