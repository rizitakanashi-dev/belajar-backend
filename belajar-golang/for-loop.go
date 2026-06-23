package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("perulangan ke-", counter)
		counter++
	}

	//for statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("perulangan ke-", counter)
	}

	//for range
	names := []string{"faris", "al", "farizi"}
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}
}
