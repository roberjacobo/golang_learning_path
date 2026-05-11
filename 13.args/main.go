package main

import (
	"fmt"
	"os"
)

// ===== COMMAND-LINE ARGUMENTS =====
// ===== ARGUMENTOS DE LÍNEA DE COMANDOS =====

// os.Args is a slice of strings holding the command-line arguments.
// os.Args[0] is the program name itself; the real arguments start at os.Args[1].

// os.Args es un slice de strings que contiene los argumentos de la línea de comandos.
// os.Args[0] es el nombre del programa; los argumentos reales empiezan en os.Args[1].

func main() {
	var s, sep string

	// Loop starts at 1 to skip os.Args[0] (the program name).
	// El bucle empieza en 1 para saltar os.Args[0] (el nombre del programa).
	for i := 1; i < len(os.Args); i++ {
		// Idiom: sep starts as "" so no leading space is added on the first
		// iteration; after that it becomes " " to join the remaining arguments.
		// Idioma: sep empieza como "" para que no se añada un espacio inicial en
		// la primera iteración; después se vuelve " " para unir los siguientes.
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
