package main

import "fmt"

func main() {
	name := "faris"

	switch name {
	case "faris":
		fmt.Println("hello faris")
	case "rizi":
		fmt.Println("hello rizi")
	default:
		fmt.Println("hi, boleh kenalan?")
	}

	//short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("nama sudah benar")
	}
}
