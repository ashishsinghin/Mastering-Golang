# **Go Packages and the Standard Library**

## **Overview**

In this lesson, you'll learn how to structure Go code using packages. You will also explore Go's powerful standard library and how to create and import your own custom packages.

---

## **1. What is a Package?**

A **package** is a way to organize and reuse code in Go.

* Every Go file starts with a `package` declaration.

```go
package main
```

---

## **2. The `main` Package and Entry Point**

The `main` package is required to build an executable Go program.

* Must contain a `main()` function:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

---

## **3. Importing Standard Library Packages**

Go provides a robust standard library. Some common packages include:

* `fmt` – formatting I/O
* `strings` – string manipulation
* `sort` – sorting slices
* `time` – working with time
* `io/ioutil` – reading and writing files

Example:

```go
import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println(strings.ToUpper("golang"))
}
```

---

## **4. Creating and Using Custom Packages**

### ➤ Step-by-step:

1. Create a new folder named `calculator`:

```
/calculator
    math.go
```

2. In `math.go`:

```go
package calculator

func Add(a, b int) int {
    return a + b
}
```

3. In your `main.go`:

```go
package main

import (
    "fmt"
    "myproject/calculator"
)

func main() {
    result := calculator.Add(2, 3)
    fmt.Println("Sum:", result)
}
```

---

## **5. Practice Exercises**

1. Use the `time` package to print the current time.
2. Create a new package with a `Multiply(a, b int) int` function and use it in `main.go`.
3. Use the `strings` package to convert text to lowercase and uppercase.

---

## **6. Best Practices**

* Use lowercase names for unexported (private) functions and uppercase for exported (public) functions.
* Group related logic into the same package.
* Keep packages small and single-purpose.

---

## **Summary**

✅ Organize code using packages for modularity
✅ The `main` package is the entry point for executables
✅ Use Go's rich standard library to simplify tasks
✅ Create and import your own reusable packages

---

## **Next Lesson:** *Constants and iota in Go – Powerful Compile-Time Values* 🚀
