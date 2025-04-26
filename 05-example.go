package main

import "fmt"

func main() {

	age := 18

	if age >= 18 {
		fmt.Println("You can vote.")
	} else if age >= 13 {
		fmt.Println("You're a teenager.")
	} else {
		fmt.Println("You're a child.")
	}

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	i := 6
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for {
		fmt.Println("This loop runs forever")
		break // avoid actual infinite loop
	}

	numbers := []int{10, 20, 30}
	for index, value := range numbers {
		fmt.Println("Index:", index, "Value:", value)
	}

	user := map[string]string{"name": "Ashish", "city": "Delhi"}
	for key, value := range user {
		fmt.Println(key, "=", value)
	}

	day := "Saturday"

	switch day {
	case "Monday":
		fmt.Println("Start of the week!")
	case "Friday":
		fmt.Println("Almost the weekend!")
	default:
		fmt.Println("Midweek grind!")
	}

	score := 85

	switch {
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 75:
		fmt.Println("Grade: B")
	default:
		fmt.Println("Grade: C")
	}
}
