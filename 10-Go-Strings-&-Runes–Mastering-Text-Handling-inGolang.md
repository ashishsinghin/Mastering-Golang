# **Working with Strings and Runes in Go**

## **Overview**

In this lesson, you'll learn how to handle and manipulate text in Go using strings and runes. Strings in Go are UTF-8 encoded, and working with them properly involves understanding both byte-level and character-level operations.

---

## **1. What is a String in Go?**

* A **string** in Go is a read-only slice of bytes.
* String literals are enclosed in double quotes (`"..."`).
* Go strings support **Unicode characters**, encoded in UTF-8.

### 🔹 Example:

```go
message := "Hello, भारत"
fmt.Println(message) // Output: Hello, भारत
```

---

## **2. Common String Operations**

### ✅ Length of a string (in bytes):

```go
len("GoLang") // Output: 6
```

### ✅ Concatenation:

```go
first := "Go"
second := "Lang"
result := first + second // Output: GoLang
```

### ✅ Accessing characters (as bytes):

```go
text := "GoLang"
fmt.Println(text[0])        // Output: 71 (byte value)
fmt.Println(string(text[0])) // Output: "G"
```

---

## **3. The `strings` Package**

Go’s `strings` package provides many useful functions.

### 🔹 Example functions:

```go
strings.ToUpper("go")          // "GO"
strings.ToLower("GoLang")      // "golang"
strings.Contains("go", "o")    // true
strings.HasPrefix("Go", "G")   // true
strings.HasSuffix("Go", "o")   // true
strings.Split("a,b,c", ",")     // ["a", "b", "c"]
strings.Replace("go go go", "go", "stop", 2) // "stop stop go"
```

---

## **4. Strings and Unicode (Runes)**

### What is a Rune?

* A **rune** in Go is an alias for `int32`.
* It represents a single **Unicode code point**.
* Runes are used for proper multi-byte character handling.

### 🔹 Iterating by rune:

```go
for i, r := range "Hello, भारत" {
    fmt.Printf("Index: %d, Rune: %c\n", i, r)
}
```

---

## **5. Converting Between Strings, Bytes, and Runes**

### 🔹 String → Byte Slice:

```go
s := "Go"
bytes := []byte(s) // [71 111]
```

### 🔹 Byte Slice → String:

```go
str := string(bytes)
```

### 🔹 String → Rune Slice:

```go
runes := []rune("भारत")
fmt.Println(runes) // [2349 2366 2352 2340]
```

---

## **Practice Exercises**

1. Create a function that counts characters (not bytes) in a string using `rune`.
2. Use `strings.Split()` to split a sentence into words.
3. Iterate through the string "Hello भारत" and print each rune's index and value.

---

## **Summary**

✅ Strings are slices of bytes in Go and support UTF-8 encoding.
✅ The `strings` package makes string manipulation easy.
✅ Use `rune` to handle multi-byte Unicode characters properly.
✅ Convert between strings, bytes, and runes for advanced manipulation.

---

## **Next Lesson:** *Arrays, Slices, and Maps – Core Data Structures in Go* 🚀
