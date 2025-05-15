package main

import "fmt"

func main() {
	fmt.Println(sumAll(10, 20, 30, 40, 50))

	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println(sumAllSlice(numbers...))
}

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func sumAllSlice(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}