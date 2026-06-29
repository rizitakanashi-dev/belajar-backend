package main

import "fmt"

func main() {
	type NoKTP string

	var ktpFaris NoKTP = "11111111111"
	fmt.Println(ktpFaris)
	fmt.Println(NoKTP("22222222222"))
}
