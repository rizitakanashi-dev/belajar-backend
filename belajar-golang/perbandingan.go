package main

import "fmt"

func main() {
	var name1 = "faris"
	var name2 = "rizi"
	var result bool = name1 == name2
	fmt.Println(result)

	value1 := 100
	value2 := 250
	fmt.Println(value1 <= value2)
	fmt.Println(value1 >= value2)
	fmt.Println(value1 != value2)
	fmt.Println(value1 == value2)
}
