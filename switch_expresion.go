package main

import "fmt"

func main() {
	name := "Pamungkas"

	switch name {
	case "Nur":
		fmt.Println("Hi Nur")
		fmt.Println("Hello Nur")
	case "Ady":
		fmt.Println("Hi Ady")
		fmt.Println("Hello Ady")
	default:
		fmt.Println("Boleh kenalan?")
		fmt.Println("Kenalan yux?")
	}

	//switch length := name; len(length) > 5 { //short statement
	//case true:
	//	fmt.Println("Nama terlalu panjang")
	//case false:
	//	fmt.Println("Nama sudah sesuai")
	//}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
