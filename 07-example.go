package main

import "fmt"

func updatebyvalue(val int) int {
	val = 10 + 10
	return val
}

func updatebyreference(val *int) int {
	*val = 10 + 10
	return *val
}

func main() {
	x := 42
	ptr := &x
	println("Value of x:", x)          // 42
	println("Address of x:", ptr)      // Address of x
	println("Value at address:", *ptr) // 42
	*ptr = 100
	println("New value of x:", x) // 100
	println("Address of x:", &x)  // Address of x

	fmt.Println("Before updatebyvalue:", updatebyvalue(x))           // 100
	fmt.Println("After updatebyvalue:", x)                           // 100
	fmt.Println("Before updatebyreference:", updatebyreference(ptr)) // 10
	fmt.Println("After updatebyreference:", x)                       // 10

	m := make(map[string]int)
	fmt.Println(m) // Output: map[]
	m["score"] = 99
	fmt.Println(m["score"]) // Output: 99

	// Using new()
	ptr1 := new(int) // Returns *int
	fmt.Println("new() example:")
	fmt.Printf("Value: %d\n", *ptr1) // 0 (zero-valued)
	fmt.Printf("Address: %p\n", ptr1)

	// Using make()
	slice := make([]int, 3)    // Returns []int
	m1 := make(map[string]int) // Returns map[string]int
	ch := make(chan int)       // Returns chan int

	fmt.Println("\nmake() example:")
	fmt.Printf("Slice: %v\n", slice) // [0 0 0]
	fmt.Printf("Map: %v\n", m1)      // map[]
	fmt.Printf("Channel: %v\n", ch)  // 0xc000...

	var p *int
	fmt.Println(p) // <nil>
	if p != nil {
		fmt.Println(*p)
	} else {
		fmt.Println("p is nil")
	}

}
