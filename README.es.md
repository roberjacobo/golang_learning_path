# Golang Learning Path

[English](README.md) | **Español**

Mis notas y ejercicios mientras aprendo Go, organizados por tema. Cada carpeta es un mini-programa ejecutable con comentarios bilingües (inglés / español).

## Estructura

| Carpeta | Tema |
|---|---|
| `1.hello/` | Primer programa: `Hello, world` |
| `2.values/` | Valores literales: strings, números, booleanos |
| `3.variables/` | Declaración de variables (`var`, `:=`) |
| `4.Constants/` | Constantes con `const` |
| `5.data_types/` | Todos los tipos de Go (numéricos, string, bool, slice, map, struct, pointer, channel, interface, etc.) |
| `6.fmt/` | Paquete `fmt`: `Println`, `Printf`, `Sprint`, verbos `%v` y `%T` |
| `7.custom_types/` | Tipos propios (`type Celsius float64`, `type UserID int`) |
| `8.conversions/` | Conversión explícita entre tipos (`int(b)`) |
| `9.exercises/` | Ejercicios cortos por tema |
| └─ `1.walrus_operator/` | Uso de `:=` |
| └─ `2.escape_sequences/` | Secuencias de escape (`\n`, `\t`) y `fmt.Sprintf` |
| └─ `3.types/` | Repaso: tipo propio + conversión explícita |
| `10.booleans/` | Operadores de comparación que devuelven `bool` |
| `11.number_types/` | Tipos numéricos, overflow en compilación y paquete `runtime` |

## Cómo correr un ejemplo

Desde la raíz del repo:

```bash
go run ./1.hello
go run ./5.data_types
go run ./9.exercises/1.walrus_operator
```

## Requisitos

- Go 1.21+ (probado con Go 1.26)

## Convenciones de comentarios

- Comentarios en **inglés y español**, uno debajo del otro.
- Solo se comenta lo que es **nuevo** en cada archivo. Conceptos ya explicados en carpetas anteriores no se repiten.
