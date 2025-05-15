package main

import "fmt"

func main() {
	i := 1

	for i <= 10 {
		fmt.Println("Perulangan ke", i)
		i++
	}

	hari := [...]string {"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	for i, hari := range hari {
		fmt.Println("Hari ke", i, "adalah", hari)
	}
}