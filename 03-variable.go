package main

import (
	"fmt"
)

func main() {
	var name string = "Ashish"
	var age int = 25

	lang := "Go"
	year := 2025

	const pi = 3.14
	const language = "Go"

	fmt.Println("name:", name)
	fmt.Println("age:", age)

	fmt.Println("lang:", lang)
	fmt.Println("year:", year)
	fmt.Printf("name: %s, age: %d, lang: %T, year: %T\n", name, age, lang, year)
	fmt.Println("pi:", pi)
	fmt.Println("language:", language)
}
