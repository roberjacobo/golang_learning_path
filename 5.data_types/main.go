package main

import "fmt"

// ===== BASIC TYPES =====
// ===== TIPOS BÁSICOS =====

// --- Signed integers (can be negative) ---
// --- Enteros con signo (pueden ser negativos) ---

// int8: -128 to 127
// int8: de -128 a 127
var i8 int8 = 21

// int16: -32,768 to 32,767
// int16: de -32,768 a 32,767
var i16 int16 = 1000

// int32: about -2 billion to 2 billion
// int32: aproximadamente de -2 mil millones a 2 mil millones
var i32 int32 = 100000

// int64: very big whole numbers
// int64: números enteros muy grandes
var i64 int64 = 9000000000

// int: size depends on the computer (32 or 64 bits). Most used.
// int: el tamaño depende de la computadora (32 o 64 bits). El más usado.
var i int = 42

// --- Unsigned integers (only zero or positive) ---
// --- Enteros sin signo (solo cero o positivos) ---

// uint8: 0 to 255 (same as byte)
// uint8: de 0 a 255 (igual que byte)
var u8 uint8 = 255

// uint16: 0 to 65,535
// uint16: de 0 a 65,535
var u16 uint16 = 60000

// uint32: 0 to about 4 billion
// uint32: de 0 a unos 4 mil millones
var u32 uint32 = 4000000000

// uint64: 0 to a huge number
// uint64: de 0 a un número enorme
var u64 uint64 = 18000000000

// uint: like int but only positive
// uint: como int pero solo positivo
var u uint = 7

// uintptr: holds a memory address as a number (rare, low level)
// uintptr: guarda una dirección de memoria como número (raro, bajo nivel)
var up uintptr = 0

// --- Aliases ---
// --- Alias ---

// byte: same as uint8, used for raw data and bytes of a file
// byte: igual que uint8, se usa para datos crudos y bytes de un archivo
var b byte = 'A' // 'A' is the number 65 / 'A' es el número 65

// rune: same as int32, used for one character (Unicode)
// rune: igual que int32, se usa para un carácter (Unicode)
var r rune = 'Ñ'

// --- Floating point (numbers with decimals) ---
// --- Punto flotante (números con decimales) ---

// float32: less precise, less memory
// float32: menos preciso, usa menos memoria
var f32 float32 = 3.14

// float64: more precise, default for decimals
// float64: más preciso, el predeterminado para decimales
var f64 float64 = 3.141592653589793

// --- Complex numbers (real + imaginary part) ---
// --- Números complejos (parte real + parte imaginaria) ---

// complex64: uses two float32
// complex64: usa dos float32
var c64 complex64 = 1 + 2i

// complex128: uses two float64
// complex128: usa dos float64
var c128 complex128 = complex(3, 4) // 3 + 4i

// --- Boolean (true or false) ---
// --- Booleano (verdadero o falso) ---
var ok bool = true

// --- String (text) ---
// --- String (texto) ---
var name string = "Roberto"

// ===== COMPOSITE TYPES =====
// ===== TIPOS COMPUESTOS =====

// Array: fixed size list of items of the same type
// Array: lista de tamaño fijo con elementos del mismo tipo
var nums [3]int = [3]int{10, 20, 30}

// Slice: a list that can grow or shrink
// Slice: una lista que puede crecer o achicarse
var fruits []string = []string{"apple", "banana", "kiwi"}

// Map: key-value pairs, like a dictionary
// Map: pares llave-valor, como un diccionario
var ages map[string]int = map[string]int{"Ana": 30, "Luis": 25}

// Struct: a group of fields with names, like a small object
// Struct: un grupo de campos con nombre, como un objeto pequeño
type Person struct {
	Name string
	Age  int
}

var p Person = Person{Name: "Sara", Age: 28}

// Pointer: holds the memory address of another variable
// Puntero: guarda la dirección de memoria de otra variable
var x int = 100
var ptr *int = &x // & means "address of" / & significa "dirección de"

// Function: a function is also a value with a type
// Función: una función también es un valor con un tipo
var greet func(string) string = func(n string) string {
	return "Hello, " + n
}

// Channel: a pipe to send values between goroutines
// Canal: una tubería para enviar valores entre goroutines
var ch chan int = make(chan int, 1) // buffered channel, size 1 / canal con buffer, tamaño 1

// Interface: a set of methods. empty interface{} can hold any value.
// Interface: un conjunto de métodos. interface{} vacía puede guardar cualquier valor.
var anything interface{} = "I can be anything"

func main() {
	// Basic numbers / Números básicos
	fmt.Println("int8:", i8, "int16:", i16, "int32:", i32, "int64:", i64, "int:", i)
	fmt.Println("uint8:", u8, "uint16:", u16, "uint32:", u32, "uint64:", u64, "uint:", u, "uintptr:", up)

	// byte and rune / byte y rune
	fmt.Printf("byte: %d (char %c)\n", b, b)
	fmt.Printf("rune: %d (char %c)\n", r, r)

	// Floats and complex / Decimales y complejos
	fmt.Println("float32:", f32, "float64:", f64)
	fmt.Println("complex64:", c64, "complex128:", c128)

	// Bool and string / Booleano y texto
	fmt.Println("bool:", ok, "string:", name)

	// Composite types / Tipos compuestos
	fmt.Println("array:", nums)
	fmt.Println("slice:", fruits)
	fmt.Println("map:", ages)
	fmt.Println("struct:", p)
	// * means "value at address" / * significa "valor en la dirección"
	fmt.Println("pointer value:", ptr, "-> points to:", *ptr)

	// Function value / Valor de función
	fmt.Println("function:", greet("Go"))

	// Channel: send a value, then read it
	// Canal: enviar un valor, luego leerlo
	ch <- 5   // send 5 into the channel / envía 5 al canal
	v := <-ch // read from the channel / lee del canal
	fmt.Println("channel got:", v)

	// Interface
	fmt.Println("interface holds:", anything)

	// Zero values: every type has a default value when not set
	// Valores cero: cada tipo tiene un valor por defecto cuando no se asigna
	var zInt int       // 0
	var zFloat float64 // 0
	var zBool bool     // false / falso
	var zStr string    // "" (empty string / cadena vacía)
	var zPtr *int      // nil (no address / sin dirección)
	fmt.Println("zero values ->", zInt, zFloat, zBool, zStr, zPtr)
}
