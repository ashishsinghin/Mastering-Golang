# **Structs in Go – Custom Data Types for Real-World Models**

## **Overview**

In this lesson, you'll learn how to define and use **structs** in Go. Structs allow you to group related data into a single custom type, making your code more organized and aligned with real-world objects.

---

## **1. What is a Struct?**

A **struct** is a composite data type that groups multiple fields under one name.

### 🔹 Example:

```go
type User struct {
    Name string
    Age  int
}
```

---

## **2. Declaring and Using Structs**

### 🔹 Declaration and Initialization:

```go
var u User
u.Name = "Ashish"
u.Age = 28
```

### 🔹 Using Struct Literal:

```go
u := User{Name: "Ashish", Age: 28}
```

### 🔹 Field Order Initialization:

```go
u := User{"Ashish", 28}
```

---

## **3. Accessing and Updating Fields**

```go
fmt.Println(u.Name) // Output: Ashish
u.Age = 29
```

---

## **4. Structs with Pointers**

Pass struct pointers to functions to modify them:

```go
func updateAge(u *User) {
    u.Age = 30
}

user := User{"Ashish", 28}
updateAge(&user)
fmt.Println(user.Age) // Output: 30
```

---

## **5. Nested Structs**

### 🔹 Defining Nested Structs:

```go
type Address struct {
    City  string
    State string
}

type Employee struct {
    Name    string
    Address Address
}
```

### 🔹 Usage:

```go
emp := Employee{
    Name: "John",
    Address: Address{
        City:  "Mumbai",
        State: "MH",
    },
}
fmt.Println(emp.Address.City) // Output: Mumbai
```

---

## **6. Anonymous Structs**

```go
person := struct {
    Name string
    Age  int
}{
    Name: "Ravi",
    Age:  24,
}
```

---

## **7. Struct Tags and JSON Serialization**

Struct tags allow custom field names for encoding/decoding:

```go
type Student struct {
    FullName string `json:"full_name"`
    Grade    int    `json:"grade"`
}

import "encoding/json"

s := Student{"Kiran", 8}
data, _ := json.Marshal(s)
fmt.Println(string(data)) // Output: {"full_name":"Kiran","grade":8}
```

---

## **Practice Exercises**

1. Create a `Book` struct with `Title`, `Author`, and `Pages` fields.
2. Write a function that accepts a pointer to `Book` and updates the page count.
3. Create a `Company` struct that contains an `Employee` and nested `Address` struct.

---

## **Summary**

✅ Structs model real-world entities with related data fields.
✅ You can initialize, access, and update struct fields.
✅ Pointers allow modification of structs inside functions.
✅ Tags enable integration with JSON and other encoders.

---

## **Next Lesson:** *Interfaces in Go – Writing Flexible and Modular Code* 🚀
