package main

import "fmt"

func main() {
	var name = make([]string, 3, 5)

	name[0] = "John"
	name[1] = "Jane"
	name[2] = "Doe"
	name[3] = "Smith"

	fmt.Println(name)
}
