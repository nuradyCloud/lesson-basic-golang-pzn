package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	var address1 Address = Address{"Jakarta Barat", "DKI Jakarta", "Indonesia"}
	//address2 := address1 // before pointer pass by value
	var address2 *Address = &address1 //after pass by reference
	var address3 *Address = &address1

	address2.City = "Semarang"

	*address2 = Address{"Bandung", "Jawa Barat", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *Address = new(Address)
	address4.City = "Malang"
	fmt.Println(address4)

	//using pointer in function
	var alamat = Address{
		City:     "Salatiga",
		Province: "Jawa Tengah",
		Country:  "",
	}
	var alamatPointer *Address = &alamat
	ChangeCountryToIndonesia(alamatPointer)
	fmt.Println(alamat)
}
