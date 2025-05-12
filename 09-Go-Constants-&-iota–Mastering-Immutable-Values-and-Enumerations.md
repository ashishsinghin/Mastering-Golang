# **Constants and iota in Go**

## **Overview**

In this lesson, you'll learn how to use constants in Go and how `iota` simplifies the creation of enumerated values. You'll also explore how Go handles typed vs untyped constants and compile-time evaluations.

---

## **1. What is a Constant?**

A **constant** is a fixed value that cannot be changed after its declaration. It is created using the `const` keyword and must be assigned a compile-time known value.

### 🔹 Example:

```go
const pi = 3.14159
const language = "Go"
```

---

## **2. Typed vs Untyped Constants**

### 🔹 Untyped Constant:

```go
const x = 5  // inferred type
```

* Adapts based on usage context

### 🔹 Typed Constant:

```go
const y int = 5  // explicitly typed
```

* Enforces a specific type

---

## **3. Grouped Constant Declarations**

```go
const (
    a = 1
    b = 2
    c = 3
)
```

---

## **4. What is `iota`?**

`iota` is a special identifier that simplifies declaration of related constants. It starts at `0` and increments automatically.

### 🔹 Example:

```go
const (
    First = iota  // 0
    Second        // 1
    Third         // 2
)
```

---

## **5. Skipping Values with `_`**

```go
const (
    _ = iota
    KB = 1 << (10 * iota)  // 1 << 10 = 1024
    MB                     // 1 << 20 = 1048576
    GB                     // 1 << 30 = 1073741824
)
```

Useful for defining byte sizes.

---

## **6. Bit Flags with iota**

Use `iota` to define bit flags in powers of 2:

```go
const (
    Read = 1 << iota  // 1
    Write             // 2
    Execute           // 4
)
```

Used in permission settings, feature flags, etc.

---

## **7. iota Resets in New `const` Blocks**

```go
const (
    A = iota  // 0
    B         // 1
)

const (
    X = iota  // 0 again
    Y         // 1
)
```

Each `const` block starts with a fresh `iota = 0`

---

## **Practice Exercises**

1. Declare a constant for Pi and calculate the area of a circle.
2. Use `iota` to define user roles: `Guest`, `User`, `Admin`, `SuperAdmin`
3. Use bit shifting with `iota` to define `KB`, `MB`, `GB`, `TB`

---

## **Summary**

✅ You now understand:

* The role and syntax of constants in Go
* The difference between typed and untyped constants
* How `iota` simplifies enumerations
* How to apply bitwise logic using `iota`

