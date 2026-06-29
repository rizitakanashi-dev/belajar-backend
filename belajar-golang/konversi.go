package main

import "fmt"

func main() {
	var nilai32 int32 = 32566

	//konversi dari int32 ke int64
	var nilai64 int64 = int64(nilai32)

	//konversi dari int32 ke int16
	var nilai16 int16 = int16(nilai32)

	fmt.Println("nilai int16 : ", nilai16)
	fmt.Println("nilai int32 : ", nilai32)
	fmt.Println("nilai int64 : ", nilai64)

	var name = "faris"
	var f byte = name[0]
	var fString string = string(f)

	fmt.Println(name)
	fmt.Println(f)
	fmt.Println(fString)
}
