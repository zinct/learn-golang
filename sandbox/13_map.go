package main

import "fmt"

func main() {
	person := map[string]string {
		"name": "Indra",
		"address": "Jl. Imam Bonjol",
		"age": "20",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["age"])

	delete(person, "age")

	fmt.Println(person)
}
