package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func sayHello(hasName HasName) {
	fmt.Println("Hello, my name is", hasName.GetName())
}

func main() {
	person := Person{Name: "Indra Mahesa"}
	sayHello(person)
}