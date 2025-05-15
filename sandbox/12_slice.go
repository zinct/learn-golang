package main

import "fmt"

func main() {
	var names = [...]string{"Indra", "Mahesa", "Audri", "Mona", "Najogi"}

	slice1 := names[2:4]
	fmt.Println(len(slice1))
	fmt.Println(slice1)

	slice2 := names[:4]
	fmt.Println(slice2)

	slice3 := names[2:]
	fmt.Println(slice3)

	hari := []string {"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	hari[5] = "Sabtu Baru"

	hari2 := append(hari, "Libur")
	hari2[0] = "Senin Lama"

	fmt.Println(hari)	
	fmt.Println(hari2)
}
