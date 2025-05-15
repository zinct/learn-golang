package main

import "fmt"

func main() {
	sayHello("Indra", getGoodBye)
	sayHello("Indra", func(name string) string {
		return "HEHE" + name
	})
}

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func sayHello(name string, filter func(string) string) {
	fmt.Println("Hello", filter(name))
}