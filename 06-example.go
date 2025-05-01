package main

import "fmt"

func greet(name string) {
	fmt.Println("Hello,", name)
}

func square(n int) int {
	return n * n
}

func divide(a, b int) (int, int) {
	return a / b, a % b
}

func add(a, b int) (sum int) {
	sum = a + b
	return
}

func sum(nums ...int) int {
	total := 0
	for _, val := range nums {
		total += val
	}
	return total
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	greet("Ashish")
	fmt.Println("Square of 5 is:", square(5))
	quotient, remainder := divide(10, 3)
	fmt.Println("Quotient:", quotient)
	fmt.Println("Remainder:", remainder)
	fmt.Println("Sum of 5 and 3 is:", add(5, 3))
	fmt.Println(sum(1, 2, 3, 4)) // 10
	fmt.Println("Factorial of 5 is:", factorial(5))
	fmt.Println("Factorial of 0 is:", factorial(0))
}
