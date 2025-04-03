# **Setting Up Go and Writing Your First Program**

## **Overview**
In this lesson, we will set up Go on different operating systems, configure the Go environment, and write our first Go program. By the end of this lesson, you will be able to run a simple "Hello, World!" program in Go.

---

## **1. Installing Go**
### **Step 1: Download and Install Go**
Go is available for Windows, macOS, and Linux. Follow these steps to install it:

1. Visit the official Go download page: **[https://go.dev/dl/](https://go.dev/dl/)**
2. Download the appropriate version for your operating system.
3. Run the installer and follow the on-screen instructions.

### **Step 2: Verify Installation**
After installation, open a terminal or command prompt and run:
```sh
 go version
```
✅ If installed correctly, it will display the Go version.

---

## **2. Setting Up Go Workspace & Environment Variables**

### **Step 1: Check Go Environment**
Run the following command to check Go’s environment settings:
```sh
go env
```
✅ This will display all Go environment variables.

### **Step 2: Understanding GOPATH & GOROOT**
- **GOPATH** – The workspace where Go projects are stored.
- **GOROOT** – The directory where Go is installed.
- **GOBIN** – The location where compiled binaries are stored.

To set `GOPATH`, add the following lines to your terminal profile (e.g., `~/.bashrc` or `~/.zshrc`):
```sh
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$PATH
```
Then run:
```sh
source ~/.bashrc  # or source ~/.zshrc
```

---

## **3. Writing Your First Go Program**

### **Step 1: Create a New Go File**
Create a folder for your project:
```sh
mkdir myfirstgo
cd myfirstgo
```
Create a new file named `main.go`:
```sh
touch main.go
```

### **Step 2: Write the Code**
Open `main.go` in a text editor and add the following code:
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

### **Step 3: Run the Program**
To execute the program, run:
```sh
go run main.go
```
✅ **Output:** `Hello, World!` 🎉

---

## **4. Understanding the Code**

| **Code**  | **Explanation**  |
|-----------|----------------|
| `package main`  | Defines the package as `main`, which is required for executable programs. |
| `import "fmt"`  | Imports the `fmt` package for formatted I/O (printing, scanning, etc.). |
| `func main() {}`  | The entry point of the Go program. |
| `fmt.Println("Hello, World!")`  | Prints `Hello, World!` to the console. |

---

## **Summary**
By the end of this lesson, you should be able to:
✅ Install Go on your system.
✅ Set up and verify your Go environment.
✅ Write and execute your first Go program.
✅ Understand the basic structure of a Go program.

---

## **Next Steps**
📌 **Next Lesson:** *Understanding Go Syntax, Variables, and Data Types* 🚀

💬 **Discussion:**
Were you able to install Go and run your first program? Share your experience in the community forum!

