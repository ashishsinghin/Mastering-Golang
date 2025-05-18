# **Arrays, Slices, and Maps in Go**

## **Overview**

In this lesson, you'll explore the three most fundamental data structures in Go: **arrays**, **slices**, and **maps**. Understanding these structures is essential for working with collections of data effectively.

---

## **1. Arrays in Go**

An array is a fixed-size collection of elements of the same type.

### 🔹 Declaration:

```go
var numbers [3]int = [3]int{1, 2, 3}
```

### 🔹 Accessing Elements:

```go
fmt.Println(numbers[0]) // Output: 1
```

### 🔹 Getting Length:

```go
fmt.Println(len(numbers)) // Output: 3
```

### 🔹 Looping Through Arrays:

```go
for i, v := range numbers {
    fmt.Println(i, v)
}
```

---

## **2. Slices in Go**

Slices are dynamically-sized, flexible wrappers around arrays.

### 🔹 Declaration:

```go
nums := []int{10, 20, 30}
```

### 🔹 Using `make()`:

```go
slice := make([]int, 5) // Creates a slice of length 5
```

### 🔹 Appending Values:

```go
slice = append(slice, 40)
```

### 🔹 Slicing an Array:

```go
arr := [5]int{1, 2, 3, 4, 5}
sub := arr[1:4] // Output: [2, 3, 4]
```

---

## **3. Maps in Go**

A map is a built-in associative data type (key-value store).

### 🔹 Creating a Map:

```go
m := make(map[string]int)
m["age"] = 25
m["score"] = 90
```

### 🔹 Accessing and Updating:

```go
fmt.Println(m["age"]) // Output: 25
m["age"] = 30
```

### 🔹 Check if Key Exists:

```go
value, ok := m["score"]
if ok {
    fmt.Println("Score is:", value)
}
```

### 🔹 Deleting a Key:

```go
delete(m, "score")
```

### 🔹 Looping Through a Map:

```go
for key, val := range m {
    fmt.Println(key, val)
}
```

---

## **Practice Exercises**

1. Declare an array of 5 integers and print each element.
2. Create a slice, append multiple values, and print the final result.
3. Create a map with 3 key-value pairs. Update one value and delete another.

---

## **Summary**

✅ Arrays are fixed-size collections of the same type
✅ Slices are dynamic and more commonly used in Go
✅ Maps store key-value pairs with fast lookup capabilities
✅ `make()` is used to create slices and maps with predefined size or structure

---

## **Next Lesson:** *Structs in Go – Custom Data Types for Real-World Models* 🚀
