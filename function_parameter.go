package main

import "fmt"

func main() {
	firstName := "Nur Ady"
	sayHelloTo(firstName, "Pamungkas")
	sayHelloTo("Rudi", "Anggara")
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}
