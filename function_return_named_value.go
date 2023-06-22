package main

import "fmt"

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func getCompleteName() (firstName, middleName, lastname string) {
	firstName = "Nur"
	middleName = "Ady"
	lastname = "Pamungkas"
	return
}
