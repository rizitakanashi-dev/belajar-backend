package main

import "fmt"

func main() {
	nama := [...]string{
		"faris",
		"rizi",
		"zen",
		"ray",
		"kyano",
		"rifqi",
	}
	slice := nama[4:6]

	fmt.Println(slice)
	fmt.Println(slice[0])
	fmt.Println(slice[1])

	days := [...]string{
		"senin",
		"selasa",
		"rabu",
		"kamis",
		"jumat",
		"sabtu",
		"minggu",
	}
	daySlice1 := days[4:]
	fmt.Println(daySlice1)

	daySlice2 := days[:6]
	fmt.Println(daySlice2)

	daySlice3 := append(daySlice2, "kamis baru")
	fmt.Println(daySlice3)

	newSlice := make([]string, 2, 5) // make([]tipe data, length slice, capasitas slice)
	newSlice[0] = "faris"
	newSlice[1] = "rizi"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice) // copy(toSlice, fromSlice)
	fmt.Println(copySlice)

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	// bedanya ada pada [5] dan []/[...]

	fmt.Println("ini array: ", iniArray)
	fmt.Println("ini slice:", iniSlice)
}
