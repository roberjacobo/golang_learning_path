# Golang Learning Path

My notes and exercises while learning Go, organized by topic. Each folder is a runnable mini-program with bilingual comments (English / Spanish).

Mis notas y ejercicios mientras aprendo Go, organizados por tema. Cada carpeta es un mini-programa ejecutable con comentarios bilingües (inglés / español).

## Structure / Estructura

| Folder / Carpeta | Topic / Tema |
|---|---|
| `1.hello/` | First program: `Hello, world` / Primer programa: `Hello, world` |
| `2.values/` | Literal values: strings, numbers, booleans / Valores literales: strings, números, booleanos |
| `3.variables/` | Variable declaration (`var`, `:=`) / Declaración de variables (`var`, `:=`) |
| `4.Constants/` | Constants with `const` / Constantes con `const` |
| `5.data_types/` | All Go types (numeric, string, bool, slice, map, struct, pointer, channel, interface, etc.) / Todos los tipos de Go (numéricos, string, bool, slice, map, struct, pointer, channel, interface, etc.) |
| `6.fmt/` | The `fmt` package: `Println`, `Printf`, `Sprint`, verbs `%v` and `%T` / Paquete `fmt`: `Println`, `Printf`, `Sprint`, verbos `%v` y `%T` |
| `7.custom_types/` | Custom types (`type Celsius float64`, `type UserID int`) / Tipos propios (`type Celsius float64`, `type UserID int`) |
| `8.conversions/` | Explicit type conversion (`int(b)`) / Conversión explícita entre tipos (`int(b)`) |
| `9.exercises/` | Short exercises by topic / Ejercicios cortos por tema |
| └─ `1.walrus_operator/` | Using `:=` / Uso de `:=` |
| └─ `2.escape_sequences/` | Escape sequences (`\n`, `\t`, etc.) / Secuencias de escape (`\n`, `\t`, etc.) |

## How to run an example

From the repo root:

## Cómo correr un ejemplo

Desde la raíz del repo:

```bash
go run ./1.hello
go run ./5.data_types
go run ./9.exercises/1.walrus_operator
```

## Requirements / Requisitos

- Go 1.21+ (tested with Go 1.26 / probado con Go 1.26)

## Comment conventions

- Comments in **English and Spanish**, one below the other.
- Only what is **new** in each file is commented. Concepts already explained in earlier folders are not repeated.

## Convenciones de comentarios

- Comentarios en **inglés y español**, uno debajo del otro.
- Solo se comenta lo que es **nuevo** en cada archivo. Conceptos ya explicados en carpetas anteriores no se repiten.
