package main

import (
	"fmt"
	"reflect"
)

type ValidationError struct {
	Message string
}

func (v *ValidationError) Error() string {
	return v.Message
}
func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, &ValidationError{Message: "pembagi tidak boleh 0"}
	}
	return nilai / pembagi, nil
}

func main() {
	hasil, err := pembagian(10, 0)
	

	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println(reflect.TypeOf(err))
		} else {
		fmt.Println("Hasil:", hasil)
	}
}
