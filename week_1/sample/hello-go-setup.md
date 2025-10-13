# 🚀 Day 1 — Setting Up Go and Your First Project

> **by [@Wavedidwhat](https://twitter.com/Wavedidwhat)**  
> Today, we’ll set up Go, create your first Go program, and understand the basics of how Go projects are structured.

---

## ⚙️ Step 1 — Install Go

🔗 **Official Go Downloads:** [https://go.dev/dl/](https://go.dev/dl/)

Download and install Go for your OS (macOS, Windows, or Linux).

After installation, verify with:

```bash
go version
```
You should see something like:

```
go version go1.20.3 darwin/amd64
```

---

## 📁 Step 2 — Create Your Project Folder

Open your terminal and create a new directory for your project:

```bash
mkdir hello-go
cd hello-go
```

---

## 🧩 Step 3 — Initialize Your Go Module

Initialize a new Go module. This sets up your project to manage dependencies:

```bash
go mod init github.com/yourusername/hello-go
```

Replace `github.com/yourusername/hello-go` with your own module path.

---

## ✍️ Step 4 — Write Your First Go Program

Create a file named `main.go`:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

This program prints "Hello, Go!" to the console.

---

## 📦 Step 5 — Create a Package

Let’s create a simple package to greet users.

1. Create a folder named `greet`:

```bash
mkdir greet
```

2. Inside `greet`, create a file `greet.go`:

```go
package greet

import "fmt"

// Hello prints a greeting message
func Hello(name string) {
    fmt.Printf("Hello, %s!\n", name)
}
```

3. Modify `main.go` to use this package:

```go
package main

import "github.com/yourusername/hello-go/greet"

func main() {
    greet.Hello("Go Learner")
}
```

Remember to replace `github.com/yourusername/hello-go` with your module path.

---

## ▶️ Step 6 — Run Your Code

Run your program with:

```bash
go run main.go
```

You should see:

```
Hello, Go Learner!
```

---

## 🧠 Explanation

- **Modules:** Go modules manage your project dependencies and versions.
- **Packages:** Packages organize your code into reusable components.
- **main package:** Entry point of your application.
- **fmt package:** Standard library package for formatted I/O.
- **go run:** Compiles and runs your Go program.

---

## 📚 Learning Resources

- [A Tour of Go](https://tour.golang.org/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go by Example](https://gobyexample.com/)
- [Go Modules Reference](https://blog.golang.org/using-go-modules)

---

## 🤔 Reflection

Take a moment to reflect:

- Did you successfully install Go?
- Can you explain what a Go module is?
- How do packages help organize your code?
- What did you find interesting or challenging?

Feel free to jot down your thoughts or share them with your peers!

---

Happy coding! 🎉
