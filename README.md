# Golang Learning Path

**English** | [Español](README.es.md)

My notes and exercises while learning Go, organized by topic. Each folder is a runnable mini-program with bilingual comments (English / Spanish).

## Structure

| Folder | Topic |
|---|---|
| `1.hello/` | First program: `Hello, world` |
| `2.values/` | Literal values: strings, numbers, booleans |
| `3.variables/` | Variable declaration (`var`, `:=`) |
| `4.Constants/` | Constants with `const` |
| `5.data_types/` | All Go types (numeric, string, bool, slice, map, struct, pointer, channel, interface, etc.) |
| `6.fmt/` | The `fmt` package: `Println`, `Printf`, `Sprint`, verbs `%v` and `%T` |
| `7.custom_types/` | Custom types (`type Celsius float64`, `type UserID int`) |
| `8.conversions/` | Explicit type conversion (`int(b)`) |
| `9.exercises/` | Short exercises by topic |
| └─ `1.walrus_operator/` | Using `:=` |
| └─ `2.escape_sequences/` | Escape sequences (`\n`, `\t`) and `fmt.Sprintf` |
| └─ `3.types/` | Recap: custom type + explicit conversion |

## How to run an example

From the repo root:

```bash
go run ./1.hello
go run ./5.data_types
go run ./9.exercises/1.walrus_operator
```

## Requirements

- Go 1.21+ (tested with Go 1.26)

## Comment conventions

- Comments in **English and Spanish**, one below the other.
- Only what is **new** in each file is commented. Concepts already explained in earlier folders are not repeated.
