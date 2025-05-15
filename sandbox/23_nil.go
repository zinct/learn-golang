package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string {
			"name": name,
		}
	}
}

func main() {
	map1 := newMap("")

	fmt.Println(map1)
}