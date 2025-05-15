package main

import "fmt"

func sayHello() {
	message := recover()
	fmt.Println(message)
	fmt.Println("Hello")
}

func sayWorld() {
	defer sayHello()

	if true {
		panic(func() {
			fmt.Println("This is panic")
		})
	}

	fmt.Println("World")
}

func main() {
	sayWorld()
}
