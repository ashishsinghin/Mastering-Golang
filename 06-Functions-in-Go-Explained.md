# **Functions in Go – Writing Reusable Code**

## **Overview**
In this lesson, you'll learn how to create and use functions in Go. Functions are essential for writing clean, reusable, and organized code.

---

## **1. What is a Function?**
A **function** is a block of code that performs a specific task. It can take inputs (parameters), process them, and return outputs (return values). Functions improve readability, reusability, and testing.

---

## **2. Defining and Calling a Function**

### ✅ Syntax:
```go
func functionName(parameter type) returnType {
    // body
}
```

### 🧪 Example:
```go
func greet(name string) {
    fmt.Println("Hello,", name)
}

func main() {
    greet("Ashish")
}
```

---

## **3. Functions with Return Values**

### 🔹 Single Return:
```go
func square(n int) int {
    return n * n
}
```
```go
result := square(5)
fmt.Println("Square:", result)
```

---

## **4. Functions with Multiple Return Values**

### 🔹 Example:
```go
func divide(a, b int) (int, int) {
    return a / b, a % b
}
```
```go
quotient, remainder := divide(10, 3)
fmt.Println("Quotient:", quotient)
fmt.Println("Remainder:", remainder)
```

---

## **5. Named Return Values**
Go allows you to name return values.

### 🔹 Example:
```go
func add(a, b int) (sum int) {
    sum = a + b
    return
}
```

---

## **6. Variadic Functions**
Use `...` to accept a variable number of arguments.

### 🔹 Example:
```go
func sum(nums ...int) int {
    total := 0
    for _, val := range nums {
        total += val
    }
    return total
}
```
```go
fmt.Println(sum(1, 2, 3)) // 6
```

---

## **7. Recursive Functions**
Functions that call themselves are called **recursive**.

### 🔹 Example:
```go
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}
```

---

## **Practice Challenge**
1. Write a function to return the square of a number.
2. Write a function that returns the minimum and maximum of two values.
3. Write a recursive function to compute the nth Fibonacci number.

---

## **Summary**
✅ You can now:
- Define and call functions in Go
- Use parameters and return values
- Work with multiple and named return values
- Create variadic and recursive functions

---

## **Next Lesson:** *Understanding Pointers and Memory Management in Go* 🚀

💬 **Discussion Prompt:**
Create a function-based calculator with add, subtract, multiply, and divide operations. Share it with the community!

