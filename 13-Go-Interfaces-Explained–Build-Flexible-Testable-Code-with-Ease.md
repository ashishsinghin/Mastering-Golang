# **Interfaces in Go – Writing Flexible and Modular Code**

## **Overview**

This lesson introduces you to **interfaces** in Go — a powerful way to define behavior. Interfaces enable polymorphism and help you write clean, testable, and modular code.

---

## **1. What is an Interface?**

An **interface** is a type that specifies a set of method signatures.

* Any type that implements those methods **implicitly** satisfies the interface.

### 🔹 Example:

```go
type Speaker interface {
    Speak() string
}
```

Any type with a `Speak()` method satisfies the `Speaker` interface.

---

## **2. Implementing an Interface**

### 🔹 Define a concrete type:

```go
type Human struct {
    Name string
}
```

### 🔹 Add the method:

```go
func (h Human) Speak() string {
    return "Hello, I’m " + h.Name
}
```

### 🔹 Use the interface:

```go
func saySomething(s Speaker) {
    fmt.Println(s.Speak())
}

h := Human{Name: "Ashish"}
saySomething(h)
```

---

## **3. Polymorphism with Interfaces**

Interfaces allow multiple types to be used interchangeably:

```go
type Dog struct {}
func (d Dog) Speak() string { return "Woof!" }

type Cat struct {}
func (c Cat) Speak() string { return "Meow!" }

func makeItTalk(s Speaker) {
    fmt.Println(s.Speak())
}
```

---

## **4. Multi-Method Interfaces**

Interfaces can require more than one method:

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

Multiple types can implement these differently.

---

## **5. Empty Interface (`interface{}`)**

The empty interface can hold any value:

```go
var anything interface{}
anything = 42
anything = "GoLang"
anything = struct{X int}{X: 10}
```

---

## **6. Type Assertion & Type Switch**

### 🔹 Type Assertion:

```go
var val interface{} = "Hello"
str := val.(string)
```

### 🔹 Type Switch:

```go
switch v := val.(type) {
case string:
    fmt.Println("String:", v)
case int:
    fmt.Println("Integer:", v)
}
```

---

## **Practice Exercises**

1. Define an interface `Vehicle` with method `Start()`
2. Implement `Car` and `Bike` types with custom `Start()` behavior
3. Create a function that takes a `Vehicle` and prints a message
4. Store values of different types in an `interface{}` and handle them using type switch

---

## **Summary**

✅ Interfaces define sets of behaviors using method signatures
✅ Types implement interfaces implicitly — no keyword needed
✅ Interfaces allow polymorphism and flexible program design
✅ The empty interface can hold any value; type assertion and switches help recover original types
