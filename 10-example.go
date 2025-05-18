package main

import (
	"fmt"
	"strings"
)

func main() {
	message := "Hello, भारत"
	fmt.Println(message)                            // Output: Hello, भारत
	fmt.Println("Length of message:", len(message)) // Output: Length of message: 19
	first := "Go"
	second := "Lang"
	result := first + second                    // Output: GoLang
	fmt.Println("Concatenated string:", result) // Output: Concatenated string: GoLang
	fmt.Println(result[0])                      // Output: 71 (byte value)
	fmt.Println(string(result[0]))              // Output: "G"

	fmt.Println(strings.ToUpper("go"))                        // "GO"
	fmt.Println(strings.ToLower("GoLang"))                    // "golang"
	fmt.Println(strings.Contains("go", "o"))                  // true
	fmt.Println(strings.HasPrefix("Go", "G"))                 // true
	fmt.Println(strings.HasSuffix("Go", "o"))                 // true
	fmt.Println(strings.Split("a,b,c", ","))                  // ["a", "b", "c"]
	fmt.Println(strings.Replace("go go go", "go", "stop", 2)) // "stop stop go"

	for i, r := range "Hello, भारत" {
		fmt.Printf("Index: %d, Rune: %c\n", i, r)
	}

	s := "Go"
	bytes := []byte(s) // [71 111]
	fmt.Println(bytes) // Output: [71 111]
	str := string(bytes)
	fmt.Println(str) // Output: "Go"
	runes := []rune("भारत")
	fmt.Println(runes) // [2349 2366 2352 2340]
}
