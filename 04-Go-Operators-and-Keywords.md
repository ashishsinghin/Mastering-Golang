# **Go Language Basics: Operators and Keywords You Need to Know**

## **Overview**
In this lesson, you'll learn about identifiers, keywords, and the various types of operators available in Go. These are essential building blocks for writing logic and structure in Go programs.

---

## **1. Identifiers in Go**
Identifiers are names used for variables, functions, constants, types, etc.

### ✅ Rules for Identifiers:
- Must begin with a letter (A-Z or a-z) or an underscore `_`
- Followed by letters, digits, or underscores
- Cannot contain symbols like `@`, `$`, `%`, etc.
- Cannot be a Go **keyword**

### ✅ Examples:
```go
var count int        // valid
var _total int       // valid
var 1stUser string   // invalid (starts with digit)
```

---

## **2. Keywords in Go**
Go has 25 reserved keywords you **cannot** use as identifiers:
```
break    default      func       interface   select  
case     defer        go         map         struct  
chan     else         goto       package     switch  
const    fallthrough  if         range       type    
continue for          import     return      var
```
These keywords are used to define the structure and control flow of Go programs.

---

## **3. Operators in Go**
Go supports several categories of operators to perform operations on variables and values.

### 🔸 Arithmetic Operators
```go
+   // Addition
-   // Subtraction
*   // Multiplication
/   // Division
%   // Modulo (remainder)
```
**Example:**
```go
a := 10
b := 3
fmt.Println(a + b)  // 13
fmt.Println(a % b)  // 1
```

### 🔸 Comparison Operators
```go
==  // Equal
!=  // Not Equal
>   // Greater Than
<   // Less Than
>=  // Greater Than or Equal
<=  // Less Than or Equal
```
**Example:**
```go
fmt.Println(a == b)  // false
fmt.Println(a > b)   // true
```

### 🔸 Logical Operators
```go
&&  // Logical AND
||  // Logical OR
!   // Logical NOT
```
**Example:**
```go
x, y := true, false
fmt.Println(x && y) // false
fmt.Println(x || y) // true
fmt.Println(!x)     // false
```

### 🔸 Assignment Operators
```go
=   // Assign
+=  // Add and Assign
-=  // Subtract and Assign
*=  // Multiply and Assign
/=  // Divide and Assign
%=  // Modulo and Assign
```
**Example:**
```go
num := 5
num += 3  // same as num = num + 3
fmt.Println(num) // 8
```

### 🔸 Bitwise Operators (Advanced)
```go
&   // AND
|   // OR
^   // XOR
<<  // Left Shift
>>  // Right Shift
```
**Example:**
```go
a := 5  // 0101
b := 3  // 0011
fmt.Println(a & b)  // 1 (0001)
```

---

## **Practice Challenge**
Write a Go program that:
- Declares two integers
- Applies arithmetic, comparison, and logical operations
- Prints the results clearly using `fmt.Println()`

---

## **Summary**
✅ You now understand identifiers and how to name things in Go.
✅ You know Go's reserved keywords and their importance.
✅ You can apply various operators to perform calculations and logic in your programs.

---

💬 **Discussion Prompt:**
Write a mini calculator in Go using operators and share your code in the forum or comments!

