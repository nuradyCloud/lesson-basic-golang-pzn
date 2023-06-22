package main

import "fmt"

func main() {
	firstName, _, _ := getFullName()
	fmt.Println(firstName)
	//fmt.Println(middleName)
	//fmt.Println(lastName)
}

func getFullName() (string, string, string) {
	return "Nur", "Ady", "Pamungkas"
}
