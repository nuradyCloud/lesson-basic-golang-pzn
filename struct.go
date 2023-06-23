package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func sayHi(customer Customer, name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func (customer Customer) sayHi2(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func (cu Customer) sayHu() {
	fmt.Println("Huuuuu from", cu.Name)
}
func main() {
	var ady Customer
	ady.Name = "Ady"
	ady.Address = "Jakarta"
	ady.Age = 30

	sayHi(ady, "Joko") // cara panggil pertama
	ady.sayHi2("Joko") // cara panggil kedua
	ady.sayHu()

	//fmt.Println(ady.Name)
	//fmt.Println(ady.Address)
	//fmt.Println(ady.Age)
	//
	//joko := Customer{
	//	Name:    "Joko",
	//	Address: "Semarang",
	//	Age:     25,
	//}
	//fmt.Println(joko)
	//
	//budi := Customer{"Budi", "Cirebon", 27}
	//fmt.Println(budi)
}
