package main

import (
	"fmt"
)

const Pi = 3.14
const lang = "Go"
const x int = 10

const (
	First  = iota // 0
	Second        // 1
	Third         // 2
)

const (
	_  = iota
	KB = 1 << (10 * iota) // 1 << 10 = 1024
	// comment

	MB // 1 << 20 = 1048576
	GB // 1 << 30 = 1073741824
)

func main() {
	fmt.Println("Value of Pi is ", Pi)
	fmt.Printf("Type of Pi is %T\n", Pi)
	fmt.Printf("Type of lang is %T\n", lang)
	fmt.Println("Value of lang is ", lang)
	fmt.Printf("Type of x is %T\n", x)
	fmt.Println("Value of x is ", x)
	fmt.Println("Value of First is ", First)
	fmt.Println("Value of Second is ", Second)
	fmt.Println("Value of Third is ", Third)
	fmt.Println("Value of KB is ", KB)
	fmt.Println("Value of MB is ", MB)
	fmt.Println("Value of GB is ", GB)
	fmt.Println("Value of 1 << 10 is ", 1<<10)
	fmt.Println("Value of 1 << 20 is ", 1<<20)
	fmt.Println("Value of 1 << 30 is ", 1<<30)
}
