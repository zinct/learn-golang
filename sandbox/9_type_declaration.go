package main

import "fmt"

func main() {
	type NoKTP string

	var noKTPIndra NoKTP = "1234567890"

	var noKTPIndraString string = string(noKTPIndra)

	fmt.Println(noKTPIndra)
	fmt.Println(noKTPIndraString)
}
