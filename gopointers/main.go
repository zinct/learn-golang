package main

import "fmt"

func main() {
	age := 21 // reguler variable

	getAdultYear(&age)
	getAdultYear(&age)
	getAdultYear(&age)

	fmt.Println(age)
}

func getAdultYear(age *int) {
	*age = *age + 1
}
