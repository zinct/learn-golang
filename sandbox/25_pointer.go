package main

import "fmt"

func main() {
	var age int = 20
	var agePointer *int = &age

	fmt.Println(age)
	fmt.Println(*agePointer)
}

func getAdultYears(age *int) int {
	return *age + 20
}
