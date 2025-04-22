package main

import "fmt"

func main() {
	// var count int  // valid
	// var _total int // valid
	//var 1stUser string  // invalid

	a := 10
	b := 3
	fmt.Println(a + b) // 13
	fmt.Println(a % b) // 1
	fmt.Println(a / b) // 3

	fmt.Println(a == b) // false
	fmt.Println(a > b)  // true

	x, y := true, false
	fmt.Println(x && y) // false
	fmt.Println(x || y) // true
	fmt.Println(!x)     // false

	a = a + 10
	a += 10
	fmt.Println(a) // 30

	j := 5              // 0101
	k := 3              // 0011
	fmt.Println(j & k)  // 1 (0001)
	fmt.Println(j | k)  // 7 (0111)
	fmt.Println(j ^ k)  // 6 (0110)
	fmt.Println(j << 1) // 10 (1010)
	fmt.Println(j >> 1) // 2 (0010)

}
