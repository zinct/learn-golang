package main

import "fmt"

type Person struct {
	Name string
	Age int
	// Cant access the struct fields | the func can be changed during runtime
	SayHello func()
}

// Can access the struct fields | the func cant be changed
func (person Person) SayHello2() {
	fmt.Println("Hello, my name is", person.Name)
}

func main() {
	person := Person{Name: "Indra Mahesa", Age: 20, SayHello: func() {
		fmt.Println("Hello, my name is" )
	}}

	fmt.Println(person)
	person.SayHello()
}