package main

import "fmt"

func main() {
	var nama string
	nama = "faris al farizi"
	fmt.Println(nama)

	nama = "rizi takanashi" //isi variabel dapat dirubah asalkan masih tipe data yang sama
	fmt.Println(nama)

	var umur = 16 //jika langsung mengisi variabelnya, maka program akan mendetaksi tipe data nya secara otomatis
	fmt.Println(umur)

	//bisa juga ditambahkan tipe datanya walaupun langsung diisi
	var score int8 = 90 //tipe datanya akan menjadi int8
	fmt.Println(score)

	//tidak menggunakan var
	jurusan := "rekayasa perangkat lunak" //membuat variabel menggunakan tanda :=
	fmt.Println(jurusan)

	//membuat variabel yang banyak sekaligus
	var (
		negara    = "indonesia"
		provinsi  = "sumatera barat"
		kode_post = 26252
	)

	fmt.Println(negara)
	fmt.Println(provinsi)
	fmt.Println(kode_post)
}
