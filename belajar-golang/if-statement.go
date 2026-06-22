package main

import "fmt"

func main() {

	nilai := 85

	if nilai >= 75 {
		fmt.Println("anda lulus")
	} else {
		fmt.Println("anda belum lulus")
	}

	nama := "faris"

	length := len(nama)

	if length > 5 {
		fmt.Println("terlalu panjang")
	} else {
		fmt.Println("nama sudah pas")
	}

	/*
		bisa juga menggunakan cara ini:

		if length := len(nama); length > 5 {
			fmt.Println("terlalu panjang")
		} else {
			fmt.Println("nama sudah pas")
		}

	*/
}
