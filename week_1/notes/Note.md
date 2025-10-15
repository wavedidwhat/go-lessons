# Week 1 — Go Basics, Packages & Imports

Welcome to Week 1! This is the moment to ignite your Go journey with energy and focus. This week is all about *thinking in Go* — building your intuition around Go's core concepts and sharpening your skills by practicing every single day.

!note Goal for the Week
By the end of this week, you will:
- Have a clear understanding of Go's basic syntax and structure
- Know how packages and imports work (and avoid common pitfalls)
- Be comfortable writing, running, and debugging Go programs
- Understand Go's tooling ecosystem (`go fmt`, `go vet`, `go test`)
- Build muscle memory through daily, focused practice

## Overview

- Master Go's syntax and core constructs (variables, types, functions)
- Organize code with packages and understand import paths
- Work with Go's project structure (`go.mod`, package organization)
- Learn control flow (if/else, loops, switches)
- Understand collections (arrays, slices, maps)
- Write tests and understand Go's testing conventions

!tip Mindset for the Week
This week is your launchpad. Approach every exercise with curiosity and determination. **Don't just code — think in Go.** Push through challenges, experiment often, break things intentionally, then fix them. Celebrate every small win. Read error messages carefully — they're your teachers.

---

## Day 1: Hello Go & Environment Setup

**Project:** `hello-go`

**Focus:** Get Go running on your machine and write your first program

**Tasks:**
- Install Go and verify with `go version`
- Create your first `main.go` with "Hello, World!"
- Run with `go run main.go`
- Build executable with `go build`

**Concepts:** `package main`, `func main()`, `fmt.Println()`, difference between `go run` and `go build`

**Challenge:** Modify to print your name and current date

!warning Common Gotchas
- Not setting up GOPATH correctly (use modules instead!)
- Forgetting `package main` declaration

---

## Day 2: Variables, Constants & Types

**Project:** `what-type`

**Focus:** Understanding Go's type system and variable declarations

**Tasks:**
- Declare variables using `var`, `:=`, and `const`
- Explore basic types: `int`, `float64`, `string`, `bool`
- Understand zero values
- Practice type conversions

**Concepts:** Variable declaration patterns, type inference, constants vs variables, explicit type conversions

**Challenge:** Create a program that converts temperatures between Celsius and Fahrenheit

!warning Common Gotchas
**Variable Shadowing** — Using `:=` in nested scopes creates new variables:
```go
// ❌ Wrong: variable shadowing
count := 1
if true {
    count := 2  // Creates NEW variable, doesn't update outer
}
fmt.Println(count) // Prints 1, not 2!

// ✅ Right: update existing variable
count := 1
if true {
    count = 2  // Updates the outer variable
}
fmt.Println(count) // Prints 2
```

!note Zero Values
Uninitialized variables have defaults: `0` for numbers, `""` for strings, `false` for bools, `nil` for pointers/slices/maps

!warning Type Conversions Are Explicit
Go doesn't do implicit casting — you must convert explicitly:
```go
var a int = 1
var b float64 = 2.5
c := float64(a) + b  // Must convert int to float64
```

---

## Day 3: Functions & Control Flow

**Project:** `control-flow`

**Focus:** Writing functions and controlling program logic

**Tasks:**
- Write functions with parameters and return values
- Use `if/else` statements
- Master `for` loops (Go's only loop!)
- Understand `switch` statements

**Concepts:** Multiple return values, named return values, early returns for error handling

**Challenge:** Write a function that checks if a number is prime

!tip Always Check Errors
Go functions often return `(result, error)` — always check the error:
```go
result, err := someFunction()
if err != nil {
    log.Fatal(err)
}
```

---

## Day 4: Arrays, Slices & Maps

**Project:** `collections`

**Focus:** Working with Go's collection types

**Tasks:**
- Create and manipulate arrays and slices
- Use `append()` and understand slice growth
- Work with maps (key-value pairs)
- Use `range` for iteration

**Concepts:** Array vs slice differences, slice backing arrays, map initialization with `make()`

**Challenge:** Build a word frequency counter

!warning Nil Maps Panic on Write
Must use `make()` before adding entries:
```go
// ❌ Wrong: nil map
var m map[string]int
m["key"] = 1  // PANIC: assignment to entry in nil map

// ✅ Right: initialize first
m := make(map[string]int)
m["key"] = 1
```

!warning Range Loop Variable Capture
Loop variables are reused — capture them when using goroutines or closures:
```go
// ❌ Wrong: all goroutines see last value
for _, v := range items {
    go func() {
        fmt.Println(v)  // Bug: prints last value repeatedly
    }()
}

// ✅ Right: pass as parameter
for _, v := range items {
    go func(v int) {
        fmt.Println(v)
    }(v)
}
```

!note Append Behavior
`append()` may reallocate the slice, changing its underlying array. Always capture the return value:
```go
slice = append(slice, newItem)
```

---

## Day 5: Structs & Methods

**Project:** `person-model`

**Focus:** Creating custom types and attaching behavior

**Tasks:**
- Define structs to model data
- Add methods to structs
- Understand pointer vs value receivers
- Practice struct composition

**Concepts:** Struct declaration and initialization, method receivers, when to use pointer receivers

**Challenge:** Create a `BankAccount` struct with deposit/withdraw methods

!tip Pointer Receivers for Mutation
Use `*T` to modify struct state:
```go
// ❌ Won't modify original
func (p Person) UpdateName(name string) {
    p.Name = name  // Only updates copy
}

// ✅ Modifies original with pointer receiver
func (p *Person) UpdateName(name string) {
    p.Name = name
}
```

!note Capitalization for Export
`Name` is exported (public), `name` is unexported (private)

---

## Day 6: Packages & Imports

**Project:** `package-practice`

**Focus:** Organizing code across multiple packages

**Tasks:**
- Create custom packages
- Understand import paths
- Use `go.mod` for module management
- Learn exported vs unexported identifiers

**Concepts:** Package declaration, module paths vs import paths, using `go mod init` and `go mod tidy`

**Challenge:** Refactor previous projects into separate packages

!warning Exported vs Unexported Names
Only capitalized names are accessible from other packages:
```go
// constants/constants.go
package constants

// ❌ Unexported - won't work from other packages
const userName = "wave"

// ✅ Exported - accessible everywhere
const UserName = "wave"
```

!tip Module Setup
```bash
# Initialize module
go mod init github.com/yourusername/yourproject

# Add dependencies
go mod tidy

# For local development, add to go.mod:
replace github.com/yourusername/yourproject => ../local-path
```

!warning Import Path Must Match Module
Check your `go.mod` file — the import path must match the module declaration

!warning Circular Imports Not Allowed
If package A imports B and B imports A, you'll get an error. Refactor common code into a third package.

---

## Day 7: Testing & Tools

**Project:** `testing-practice`

**Focus:** Writing tests and using Go's toolchain

**Tasks:**
- Write unit tests with `testing` package
- Run tests with `go test`
- Use `go fmt`, `go vet`, and `gofmt`
- Understand test file naming (`_test.go`)

**Concepts:** Test function naming, table-driven tests, test coverage, race detection

**Challenge:** Write tests for all functions from Day 3-5

!note Test File Naming
Test files must end in `_test.go`:
```go
// calculator_test.go
package calculator

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    
    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}
```

!tip Essential Go Commands (Run These Daily!)
```bash
# Format your code
go fmt ./...

# Check for suspicious code
go vet ./...

# Tidy dependencies
go mod tidy

# Run tests
go test ./...

# Run tests with race detector
go test -race ./...

# Run tests without cache
go test -count=1 ./...
```

---

## Top 10 Week 1 Gotchas

!warning Keep This Handy!
1. **Exported vs Unexported** — Capitalize to export: `UserName` not `userName`
2. **Nil Maps** — Must `make()` before writing: `m := make(map[string]int)`
3. **Variable Shadowing** — `:=` creates new variables, use `=` to update
4. **Type Conversions** — Always explicit: `float64(intVar) + floatVar`
5. **Zero Values** — Variables have defaults: `0`, `""`, `false`, `nil`
6. **Range Loop Capture** — Pass loop variables as parameters to closures
7. **Import Paths** — Must match `go.mod` module path
8. **Circular Imports** — Not allowed; refactor into third package
9. **Error Checking** — Always check `if err != nil`
10. **Append Behavior** — May reallocate slice; capture return value

---

## Resources

!tip Official Resources
- [Go Tour](https://go.dev/tour/) — Interactive Go lessons
- [Effective Go](https://go.dev/doc/effective_go) — Idiomatic patterns
- [Go by Example](https://gobyexample.com/) — Practical code examples
- [Go Playground](https://go.dev/play/) — Test code online
- [Go Documentation](https://pkg.go.dev/) — Standard library reference

---

## End of Week Checklist

- [ ] Go installed and `go version` works
- [ ] Written "Hello, World!" program
- [ ] Comfortable with variables, types, and constants
- [ ] Written functions with multiple return values
- [ ] Used arrays, slices, and maps
- [ ] Created structs with methods
- [ ] Built custom packages with proper exports
- [ ] Written and run unit tests
- [ ] Used `go fmt`, `go vet`, and `go mod tidy`
- [ ] Read through common gotchas

!note Reflection Questions
- What Go concepts clicked immediately? Which need more practice?
- How comfortable are you with packages and imports?
- What gotchas did you hit? How did you solve them?
- Which Go idioms feel natural vs. awkward?
- What will you do differently in Week 2?

!tip Next Week Preview
Interfaces, error handling, and concurrency basics. Get ready!