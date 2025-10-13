# Go Command-Line Tools and Workflows

Welcome to your beginner-friendly guide to Go’s command-line tools! If you’re new to Go or just want to get a clearer picture of how things work behind the scenes, this guide will walk you through the essential commands and concepts you’ll use every day.

---

## 🧰 Go Command Basics

The `go` command is your main tool for working with Go programs. Think of it as the Swiss Army knife for Go: it helps you **build**, **run**, **test**, and **manage** your code.

When you type `go` in your terminal, you’re invoking this tool, which understands a variety of subcommands like `run`, `build`, `test`, `mod`, and more. It handles everything from compiling your code into an executable to managing dependencies.

---

## ⚙️ Go Run & Build

### `go run`

- Use this when you want to **quickly run your Go program** without creating a binary file.
- It compiles your code in memory and runs it immediately.
- Great for testing small programs, scripts, or trying out code snippets.

**Example:**

```bash
go run main.go
```

This compiles and runs `main.go` right away.

### `go build`

- Use this when you want to **compile your program into an executable binary**.
- It creates a file (like `main` or `main.exe`) that you can run anytime without needing the Go tool.
- Useful when you want to distribute your program or run it multiple times.

**Example:**

```bash
go build main.go
./main   # Runs the compiled binary on Unix-like systems
```

---

## 🧩 Go Mod (Modules & Packages)

### What is a package?

A **package** is a way to organize your Go code into reusable components. Each folder with Go files usually represents a package, and you can import packages into your programs to use their functionality.

### What are modules?

Modules are how Go manages your project's dependencies and versions. They replace the older GOPATH system and make it easier to track and share code.

### Key commands:

- `go mod init <module-name>`  
  Initializes a new module in your project folder. This creates a `go.mod` file, which tracks your dependencies.

- `go mod tidy`  
  Cleans up your `go.mod` and `go.sum` files by adding missing and removing unused dependencies.

- `go get <package>`  
  Adds or updates a dependency package in your module.

**Example:**

```bash
go mod init github.com/yourusername/yourproject
go get github.com/some/dependency
go mod tidy
```

This sets up your module and manages dependencies smoothly.

---

## 🧪 Go Test

Go has **built-in testing** that integrates seamlessly with your code.

- Write test functions in files named with the suffix `_test.go`.
- Test functions start with `Test` and take a pointer to `testing.T`.
- Run tests easily with `go test`.

**Example test file (`math_test.go`):**

```go
package math

import "testing"

func TestAdd(t *testing.T) {
    got := Add(2, 3)
    want := 5
    if got != want {
        t.Errorf("Add(2, 3) = %d; want %d", got, want)
    }
}
```

Run the tests with:

```bash
go test
```

This command looks for all `_test.go` files and runs the tests inside.

---

## 🧱 Compile & Install Applications

### `go build`

As mentioned before, this compiles your code into an executable binary in your current directory.

### `go install`

- Compiles your code **and installs** the resulting binary into your `$GOPATH/bin` directory (or `$HOME/go/bin` by default).
- This way, you can run your program from anywhere in your terminal (if `$GOPATH/bin` is in your PATH).

**Example:**

```bash
go install github.com/yourusername/yourproject/cmd/app
```

After running this, the binary will be placed in `$GOPATH/bin`, and you can just type `app` to run it.

---

## 🧭 Summary & Real-World Tip

These commands — `go run`, `go build`, `go mod`, `go test`, and `go install` — form the backbone of every Go developer’s workflow. Whether you’re writing simple scripts, managing complex dependencies, or shipping production applications, you’ll use these tools daily.

**Pro tip:** Get comfortable with `go mod` early on. Managing dependencies well saves you headaches later. Also, remember that `go test` is your best friend for writing reliable, bug-free code.

Happy coding! 🚀
