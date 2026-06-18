package main

import "fmt"

func main() {

	//const tidak bisa diubah nilainya
	const nama string = "faris al farizi"
	const umur = 16

	//membuat constant yang banyak sekaligus
	const (
		firstName string = "faris"
		lastName         = "al farizi"
		age              = 17
	)

	fmt.Println(nama)
	fmt.Println(umur)
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)

}
