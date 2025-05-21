package main

import "fmt"

func main() {
	var numbers [3]int = [3]int{1, 2, 3}
	fmt.Println(numbers)

	fmt.Println(numbers[0])
	fmt.Println(numbers[1])
	fmt.Println(numbers[2])
	fmt.Println(len(numbers))

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	for i, v := range numbers {
		fmt.Println(i, v)
	}

	nums := []int{10, 20, 30}
	fmt.Println(nums)

	slice := make([]int, 5) // Creates a slice of length 5
	slice = append(slice, 40)
	fmt.Println(slice)

	arr := [5]int{1, 2, 3, 4, 5}
	sub := arr[1:4] // Output: [2, 3, 4]
	fmt.Println(sub)

	m := make(map[string]int)
	m["age"] = 25
	m["score"] = 90
	fmt.Println(m)
	fmt.Println(m["age"])
	fmt.Println(m["score"])
	fmt.Println(len(m))
	for key, value := range m {
		fmt.Println(key, value)
	}
	delete(m, "age")
	fmt.Println(m)
	fmt.Println(m["age"])   // Output: 0 (default value for int)
	fmt.Println(m["score"]) // Output: 90

	val, ok := m["score"]
	fmt.Println(val, ok) // Output: 0 false (0 is the default value for int)
	if ok {
		fmt.Println("Key exists")
	} else {
		fmt.Println("Key does not exist")
	}
}
