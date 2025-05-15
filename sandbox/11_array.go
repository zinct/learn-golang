package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Indra"
	names[1] = "Mahesa"

	fmt.Println(names)

	// Panjang Array walaupun data tidak diisi
	fmt.Println(len(names))

	// Array Literal
	var values = [...]int{
		90,
		80,
	}

	fmt.Println(values)
	fmt.Println(len(values))
	
	
}