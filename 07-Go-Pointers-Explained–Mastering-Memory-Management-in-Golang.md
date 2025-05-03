# **Understanding Pointers and Memory Management in Go**

## **Overview**

This lesson will help you understand how pointers work in Go and how they impact memory management. You'll also explore the difference between `new()` and `make()` and how to pass data by reference.

---

## **1. What is a Pointer?**

A **pointer** is a variable that stores the memory address of another variable.

---

## **2. Declaring and Using Pointers**

### 🔹 Declaration:

```go
var ptr *int
```

### 🔹 Assigning a Value:

```go
x := 42
ptr = &x
fmt.Println(ptr)  // Memory address
fmt.Println(*ptr) // Value at the address (42)
```

---

## **3. Dereferencing a Pointer**

Dereferencing means accessing the value at the memory address:

```go
*ptr = 100
fmt.Println(x) // Output: 100
```

---

## **4. Pass by Value vs Pass by Reference**

### 🔹 Pass by Value:

```go
func update(val int) {
    val = 10
}
```

This won't change the original value.

### 🔹 Pass by Reference:

```go
func update(val *int) {
    *val = 10
}
```

This modifies the original variable.

Usage:

```go
x := 5
update(&x)
fmt.Println(x) // Output: 10
```

---

## **5. Using `new()`**

Allocates memory and returns a pointer to the zero value of the type.

```go
ptr := new(int)
*ptr = 20
fmt.Println(*ptr) // Output: 20
```

---

## **6. Using `make()`**

Used for creating **slices, maps, and channels**:

```go
m := make(map[string]int)
m["score"] = 99
fmt.Println(m["score"]) // Output: 99
```

---

## **7. Nil Pointers**

Pointers are `nil` if not initialized:

```go
var p *int
fmt.Println(p) // <nil>
```

Always check before dereferencing:

```go
if p != nil {
    fmt.Println(*p)
}
```

---

## **Practice Exercises**

1. Write a function to update a value using a pointer
2. Create a map using `make()` and store at least two key-value pairs
3. Demonstrate the difference between passing by value and passing by reference

---

## **Summary**

✅ Understand pointers and how they store memory addresses
✅ Use `*` to dereference and `&` to reference
✅ Know when to use `new()` and `make()`
✅ Modify values through function arguments using pointers

---

## **Next Lesson:** *Go Packages and the Standard Library – Organizing Code in Go* 🚀
