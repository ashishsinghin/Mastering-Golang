# **Control Structures in Go – If, Loops, and Switch**

## **Overview**
In this lesson, we explore Go's control flow mechanisms including `if`, `for`, and `switch`. These are essential for writing logic, making decisions, and iterating over data in your Go programs.

---

## **1. If, Else If, and Else Statements**
Go uses simple and readable `if-else` syntax for decision making.

### ✅ Syntax:
```go
if condition {
    // do something
} else if anotherCondition {
    // do something else
} else {
    // fallback case
}
```

### 🧪 Example:
```go
age := 18

if age >= 18 {
    fmt.Println("You can vote.")
} else if age >= 13 {
    fmt.Println("You're a teenager.")
} else {
    fmt.Println("You're a child.")
}
```

---

## **2. For Loops in Go**
Go provides a single looping construct: the `for` loop. It can be used in several styles.

### 🔸 Standard For Loop:
```go
for i := 1; i <= 5; i++ {
    fmt.Println(i)
}
```

### 🔸 Go-style While Loop:
```go
i := 1
for i <= 5 {
    fmt.Println(i)
    i++
}
```

### 🔸 Infinite Loop:
```go
for {
    fmt.Println("This loop runs forever")
    break // avoid actual infinite loop
}
```

---

## **3. Range Loops (Iterating Collections)**
Use `range` to iterate over slices, arrays, maps, and strings.

### 🔸 Looping through a Slice:
```go
numbers := []int{10, 20, 30}
for index, value := range numbers {
    fmt.Println("Index:", index, "Value:", value)
}
```

### 🔸 Looping through a Map:
```go
user := map[string]string{"name": "Ashish", "city": "Delhi"}
for key, value := range user {
    fmt.Println(key, "=", value)
}
```

---

## **4. Switch Statements**
Switch is useful for cleaner multi-condition logic.

### ✅ Basic Syntax:
```go
switch value {
case option1:
    // do something
case option2:
    // do something else
default:
    // fallback
}
```

### 🧪 Example:
```go
day := "Monday"

switch day {
case "Monday":
    fmt.Println("Start of the week!")
case "Friday":
    fmt.Println("Almost the weekend!")
default:
    fmt.Println("Midweek grind!")
}
```

### 🔸 Switch Without Expression:
```go
score := 85

switch {
case score >= 90:
    fmt.Println("Grade: A")
case score >= 75:
    fmt.Println("Grade: B")
default:
    fmt.Println("Grade: C")
}
```

---

## **Practice Challenge**
Write a Go program that:
1. Accepts a number
2. Uses `if` to check if it's positive, negative, or zero
3. Uses a `for` loop to count from 1 up to the number
4. Uses `switch` to print a message based on the range

---

## **Summary**
✅ You now understand:
- How to use conditional logic with `if`, `else if`, and `else`
- How to write various forms of `for` loops
- How to use `range` to iterate over collections
- How to write efficient and readable `switch` statements

---


💬 **Discussion Prompt:**
Write a program using all three constructs (`if`, `for`, `switch`) and share it with the community!

