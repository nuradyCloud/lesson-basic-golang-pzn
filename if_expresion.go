package main

import "fmt"

func main() {
	var name = "Nur"

	if name == "Ady" {
		fmt.Println("Hai Ady")
	} else if name == "Nur" {
		fmt.Println("Hi Nur")
	} else {
		fmt.Println("Boleh kenalan kakak?")
	}

	if length := len(name); length > 5 { // type short condition
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
