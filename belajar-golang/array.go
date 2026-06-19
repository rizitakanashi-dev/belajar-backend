package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "faris"
	names[1] = "rizi"
	names[2] = "takanashi"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var nilai = [3]int{
		98,
		76,
		8,
	}

	fmt.Println(nilai)
	fmt.Println(nilai[0])
	fmt.Println(nilai[1])
	fmt.Println(nilai[2])

	fmt.Println(len(names))
	fmt.Println(len(nilai))
}
