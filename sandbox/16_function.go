package main

import "fmt"

func main() {
	sayHello("Indra")
	name, length, isTrue := getHello("Indra")
	fmt.Println(name, length, isTrue)
}

func sayHello(name string) {
	fmt.Println("Hello", name)
}

func getHello(name string) (string, int, bool) {
	return "Hello " + name, len(name), true
}