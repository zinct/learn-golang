package main

import "fmt"

func random() interface{} {
		return [...]int{1, 2, 3, 4, 5}
}

func handleError() {
	message := recover()
	fmt.Println("Error with message: ", message)
}

func main() {
	var result interface{} = random()

	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown type")
	}

	var resultString string = result.(string)

	defer handleError()
	resultInt := result.(int)

	fmt.Println(resultString)
	fmt.Println(resultInt)
}