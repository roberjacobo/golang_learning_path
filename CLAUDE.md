# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository purpose

Personal Go learning notebook. Each numbered folder (`1.hello/`, `2.values/`, … `12.strings/`) is an independent runnable mini-program demonstrating one Go concept. `9.exercises/` holds short practice programs grouped by topic. `notes/` holds longer-form study notes in Markdown.

This is **not** a library or application — there are no shared packages between folders. Every example uses `package main` with its own `main()`.

## Common commands

Run from the repo root (paths use `./` because the module is `github.com/roberjacobo/golang_learning_path`):

```bash
go run ./1.hello                       # run a top-level example
go run ./9.exercises/1.walrus_operator # run a nested exercise
go build ./...                         # compile everything to verify it builds
go vet ./...                           # static checks across all examples
go mod tidy                            # after adding/removing imports
```

There is no test suite.

## Conventions for new examples

- **Bilingual comments (English / Spanish), one below the other.** Every conceptual comment block has both languages. See `7.custom_types/main.go` for the canonical style.
- **Only comment what is new in that file.** Concepts already explained in earlier-numbered folders are not re-explained — readers progress through folders in order.
- The file inside a numbered folder is named either `main.go` (most recent convention) or matches the topic (`hello.go` in `1.hello/`). Either is acceptable; prefer `main.go` for new folders.
- New numbered folders should also be added to the table in `README.md` (and ideally `README.es.md`).

## Go version

`go.mod` declares `go 1.26.2`. The only non-stdlib dependency is `rsc.io/quote/v4`, used by `1.hello/`.
