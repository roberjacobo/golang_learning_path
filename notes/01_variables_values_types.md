# Quiz Notes: Variables, Values, and Types

**English** | [Español abajo](#notas-quiz-variables-valores-y-tipos)

## Core vocabulary

- **Statement (Declaración)** → smallest element of a program that expresses an action to execute. Usually starts a line of code. Example: `x := 5`.
- **Expression (Expresión)** → combination of values, constants, variables, operators, and functions that evaluates to another value. Example: `2 + 3` evaluates to `5`.
- **Operator** → the symbol that performs the action (e.g. `+` in `2 + 2`).
- **Operand** → the values the operator works on (e.g. the two `2`s in `2 + 2`).
- **Identifier** → the name given to a variable, function, or constant.
- **Keyword (palabra clave)** → words reserved by Go (`var`, `func`, `package`, `const`, `type`, `import`, etc.). They have a fixed purpose and cannot be repurposed.

## Brackets in Go

| Symbol | Name (EN) | Name (ES) |
|---|---|---|
| `( )` | parentheses | paréntesis |
| `{ }` | braces / curly braces | llaves |
| `[  ]` | square brackets | corchetes |

## Scope

The **scope** of a variable is the part of the code that has access to it (read or write). True/Verdadero.

## Types

- **Primitive type (primitivo)** → built into the language. In Go: `int`, `float64`, `string`, `bool`, `byte`, `rune`, etc. → ✅ `int` IS primitive, `string` IS primitive.
- **Composite type (compuesto)** → built from other types: arrays, slices, maps, structs, channels, interfaces, functions, pointers.
- **Zero value (valor cero)** → when you declare with `var` and no value, Go assigns a default: `int=0`, `float64=0`, `bool=false`, `string=""`, pointers/maps/slices/interfaces=`nil`.
- **Custom type (tipo propio)** → you can create your own with `type Name UnderlyingType` (e.g. `type Money int`).
- **Underlying type (tipo subyacente / raíz)** → every custom type has one (e.g. `Money`'s underlying type is `int`).
- **Conversion (NOT casting)** → Go says "conversion": `int(b)`. Required between different named types.

## The empty interface

- `interface{}` → empty interface. **Every** Go value satisfies it.
- `...interface{}` → variadic parameter. The `...` means the function accepts any number of values of that type.

## Program structure

- Entry point: `func main()` inside `package main`. ✅ True.
- `package`, `var`, `func`, `import` → all reserved keywords.
- Short declaration `:=` → **only inside functions**, NOT at package level. ❌ False to "anywhere".

## Calling functions from packages

- Syntax: `package.identifier` → e.g. `fmt.Println()` calls `Println` from the `fmt` package.
- `fmt.Sprintf()` → returns a formatted string (instead of printing it).

## Idiomatic Go

- "Idiomatic Go" = code written in the natural style and conventions of Go (clear names, no unnecessary abstractions, follows Effective Go guidelines).
- **Good package names** (per Effective Go): **short**, **concise**, **evocative** (corto, conciso, evocativo).

## Unused values and variables

- `_` (underscore) → the **blank identifier**. Tells the compiler "ignore this value".
- Unused variables in Go → ❌ **NO**. The compiler errors. (Unused **imports** also error.)

## Documentation and resources

- **golang.org** → official site, language spec, blog, downloads.
- **pkg.go.dev** (replaced godoc.org) → searchable package documentation.
- **Go Playground** (play.golang.org) → run Go code in the browser.
- **Golang Bridge forum** (https://forum.golangbridge.org/) → community Q&A.

## Meta-language

Yes — there is a language we use to talk about the language (statement, expression, operand, scope, etc.). That is what this whole vocabulary list is.

## Quick true/false cheatsheet

| Statement | Answer |
|---|---|
| `int` is a primitive type in Go | ✅ True |
| `string` is a primitive type in Go | ✅ True |
| `package` is NOT reserved | ❌ False (it IS reserved) |
| `var` is reserved | ✅ True |
| `+` is the operator in `2 + 2` | ✅ True |
| The `2`s are operands | ✅ True |
| `:=` works at package level | ❌ False |
| `func main()` lives in `package main` | ✅ True |
| You can have unused variables | ❌ No |
| Go says "casting" | ❌ False (Go says **conversion**) |
| You can create your own type | ✅ Yes |
| Custom types have an underlying type | ✅ True |
| Every value satisfies `interface{}` | ✅ True |

---

# Notas Quiz: Variables, Valores y Tipos

[English above](#quiz-notes-variables-values-and-types) | **Español**

## Vocabulario clave

- **Declaración (statement)** → el elemento más pequeño de un programa que expresa una acción a ejecutar. Normalmente inicia una línea de código.
- **Expresión (expression)** → combinación de valores, constantes, variables, operadores y funciones que evalúa a otro valor. Ejemplo: `2 + 3` evalúa a `5`.
- **Operador** → el símbolo que ejecuta la acción (`+` en `2 + 2`).
- **Operando** → los valores sobre los que actúa el operador (los dos `2` en `2 + 2`).
- **Identificador** → el nombre que se le da a una variable, función o constante.
- **Palabra clave (keyword)** → palabras reservadas por Go (`var`, `func`, `package`, `const`, `type`, `import`, etc.). Tienen un propósito fijo y no se pueden reutilizar para algo distinto.

## Símbolos

| Símbolo | Nombre |
|---|---|
| `( )` | paréntesis |
| `{ }` | llaves |
| `[  ]` | corchetes |

## Scope

El **scope** (alcance) de una variable es la parte del código que puede leerla o escribirla. ✅ Verdadero.

## Tipos

- **Tipo primitivo** → tipo básico construido dentro del lenguaje. En Go: `int`, `float64`, `string`, `bool`, `byte`, `rune`. → ✅ `int` SÍ es primitivo, `string` SÍ es primitivo.
- **Tipo compuesto** → te permite componer valores a partir de otros tipos: arrays, slices, maps, structs, channels, interfaces, funciones, punteros.
- **Valor cero (zero value)** → cuando declaras con `var` sin asignar, Go pone un valor por defecto: `int=0`, `float64=0`, `bool=false`, `string=""`, punteros/maps/slices/interfaces=`nil`.
- **Tipo propio** → lo creas con `type Nombre TipoSubyacente` (ej. `type Money int`).
- **Tipo subyacente / raíz** → cada tipo propio tiene uno (el de `Money` es `int`).
- **Conversión (no "casting")** → en Go decimos **conversión**: `int(b)`. Es obligatoria entre tipos con nombre distinto.

## La interfaz vacía

- `interface{}` → interfaz vacía. **Todo** valor en Go la satisface.
- `...interface{}` → parámetro variádico. Los `...` significan que la función acepta una cantidad variable de valores de ese tipo.

## Estructura de un programa

- Punto de entrada: `func main()` dentro de `package main`. ✅ Verdadero.
- `package`, `var`, `func`, `import` → todas son palabras reservadas.
- Operador de declaración corta `:=` → **solo dentro de funciones**, NO a nivel de paquete. ❌ Falso si dice "en cualquier parte".

## Llamar funciones desde paquetes

- Sintaxis: `paquete.identificador` → ej. `fmt.Println()` llama a `Println` del paquete `fmt`.
- `fmt.Sprintf()` → devuelve un string con formato (en lugar de imprimirlo). Útil cuando quieres "imprimir a un string" y guardarlo en una variable.

## Go idiomático

- "Código de Go idiomático" = código escrito en el estilo natural y convenciones de Go (nombres claros, sin abstracciones innecesarias, sigue Effective Go).
- **Buenos nombres de paquetes** (según Effective Go): **corto**, **conciso**, **evocativo**.

## Valores y variables sin usar

- `_` (guion bajo / underscore) → el **identificador en blanco** ("blank identifier"). Le dice al compilador "ignora este valor".
- Variables sin usar → ❌ **NO** se permite. El compilador da error. (Los **imports** sin usar también dan error.)

## Documentación y recursos

- **golang.org** → sitio oficial: spec del lenguaje, blog, descargas.
- **pkg.go.dev** (reemplazó a godoc.org) → documentación buscable de paquetes.
- **Go Playground** (play.golang.org) → corre Go en el navegador.
- **Foro Golang Bridge** (https://forum.golangbridge.org/) → preguntas de la comunidad.

## Meta-lenguaje

Sí — hay un lenguaje que usamos para hablar del lenguaje (declaración, expresión, operando, scope, etc.). Toda esta lista de vocabulario es justamente eso.

## Cheatsheet rápido (V/F)

| Afirmación | Respuesta |
|---|---|
| `int` es primitivo en Go | ✅ Verdadero |
| `string` es primitivo en Go | ✅ Verdadero |
| `package` NO es reservada | ❌ Falso (SÍ es reservada) |
| `var` es reservada | ✅ Verdadero |
| `+` es el operador en `2 + 2` | ✅ Verdadero |
| Los `2` son operandos | ✅ Verdadero |
| `:=` se puede usar a nivel de paquete | ❌ Falso |
| `func main()` vive en `package main` | ✅ Verdadero |
| Puedes tener variables sin usar | ❌ No |
| En Go decimos "casting" | ❌ Falso (decimos **conversión**) |
| Puedes crear tu propio tipo | ✅ Sí |
| Los tipos propios tienen tipo subyacente | ✅ Verdadero |
| Todo valor satisface `interface{}` | ✅ Verdadero |
