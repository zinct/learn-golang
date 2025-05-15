package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var c = a * b

	fmt.Println(c)

	// Augmented Assignment
	var i = 1
	i += 10

	// Unary Operator
	i++
	i--

	// Logical Operator
	var (
		ujian = 80
		absensi = 80
		nilaiAkhir = 80
	)

	var lulusUjian = ujian >= 80
	var lulusAbsensi = absensi >= 80
	var lulusNilaiAkhir = nilaiAkhir >= 80

	var lulus = lulusUjian && lulusAbsensi && lulusNilaiAkhir

	fmt.Println(lulus)
	fmt.Println(ujian >= 80)
	fmt.Println(absensi >= 80)
	fmt.Println(nilaiAkhir >= 80)

	// Comparison Operator
	var (
		name1 = "Indra"
		name2 = "Indra"
	)

	fmt.Println(name1 == name2)
	fmt.Println(name1 != name2)
	fmt.Println(name1 > name2)
	fmt.Println(name1 < name2)
	fmt.Println(name1 >= name2)
}