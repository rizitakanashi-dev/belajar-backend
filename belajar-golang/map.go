package main

import "fmt"

func main() {
	person := map[string]string{
		"nama":    "faris",
		"jurusan": "rekayasa perangkat lunak",
	}

	person["provinsi"] = "sumatera barat" // untuk nambah/ngerubah

	fmt.Println(person)
	fmt.Println(person["nama"])
	fmt.Println(person["jurusan"])

	buku := make(map[string]string)
	buku["title"] = "10 dosa besar soeharto"
	buku["author"] = "rizi"
	buku["salah"] = "ups"

	delete(buku, "salah")
	fmt.Println(buku)
	fmt.Println(len(buku))
}
