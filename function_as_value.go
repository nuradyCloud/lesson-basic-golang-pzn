package main

import "fmt"

func main() {
	sayGoodBye := getGooBye
	fmt.Println(sayGoodBye("Ady"))
}

func getGooBye(name string) string {
	return "Good Bye " + name
}
