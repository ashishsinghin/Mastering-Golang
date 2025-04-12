# **Go Syntax, Variables, and Data Types**

## **Overview**
In this lesson, you will learn how Go's syntax works, how to declare variables and constants, and how to use Go's built-in data types effectively. These are the foundational concepts every Go developer needs.

---

## **1. Basic Syntax Rules in Go**
Go is known for its simplicity and strict formatting. Here are some important syntax rules:

- Curly braces `{}` are required for all code blocks (e.g., `if`, `for`, `func`).
- Parentheses are not required around conditions (`if`, `for`, etc.).
- Semicolons are **not** used at the end of statements.
- All code must be written inside a **package**, and the executable program must have a `main` function.
- Use the `gofmt` tool to auto-format your code consistently.

---

## **2. Declaring Variables in Go**

### **Using `var` keyword:**
```go
var name string = "Ashish"
var age int = 25
```

### **Using short declaration (`:=`):**
```go
language := "Go"
year := 2009
```
> ⚠️ `:=` can only be used inside functions.

### **Multiple variable declarations:**
```go
var a, b, c int = 1, 2, 3
x, y := 10, "hello"
```

---

## **3. Constants in Go**
Constants are declared using the `const` keyword and cannot be reassigned:
```go
const pi = 3.14
const language = "Golang"
```

---

## **4. Built-in Data Types**
Go provides a rich set of data types:

| Type           | Example                         | Description                          |
|----------------|----------------------------------|--------------------------------------|
| `int`, `int8`, `int64`  | `var age int = 30`         | Integer types                        |
| `float32`, `float64`    | `var price float64 = 9.99` | Floating-point numbers              |
| `bool`         | `var active bool = true`         | Boolean type                         |
| `string`       | `var name string = "Go"`         | Text (a sequence of characters)     |
| `complex64`, `complex128` | `var z complex64 = 1 + 2i` | Complex numbers (used less often)   |

---

## **5. Type Inference in Go**
Go can automatically infer the type when using short declarations:
```go
name := "Ashish"   // inferred as string
count := 5         // inferred as int
```

---

## **6. Zero Values in Go**
Variables declared without initialization get default zero values:
```go
var name string   // ""
var age int       // 0
var isActive bool // false
```
This feature helps prevent the use of uninitialized variables.

---

## **7. Code Example**
```go
package main

import "fmt"

func main() {
    var username string = "ashish"
    age := 30
    const pi = 3.1415
    isGoAwesome := true

    fmt.Println("Name:", username)
    fmt.Println("Age:", age)
    fmt.Println("Pi:", pi)
    fmt.Println("Is Go Awesome?", isGoAwesome)
}
```
**Output:**
```
Name: ashish
Age: 30
Pi: 3.1415
Is Go Awesome? true
```

---

## **Summary**
By the end of this lesson, you should:
✅ Understand how Go syntax works and why it’s strict.
✅ Know how to declare variables using `var` and `:=`.
✅ Be able to use constants.
✅ Work with Go's built-in data types and zero values.

---

## **Next Steps**
📌 **Next Lesson:** *Control Structures in Go – If, Loops, and Switch Statements* 🚀

💬 **Practice Task:**
Try declaring 3 variables (string, int, bool) and print their values using `fmt.Println()`. Share your code in the community or comments!

