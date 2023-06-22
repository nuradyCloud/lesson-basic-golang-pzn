package main

import (
	"fmt"
)

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	namedFiltered := filter(name)
	fmt.Println("Hello", namedFiltered)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "****"
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Ady", spamFilter)
	sayHelloWithFilter("anjing", spamFilter)
}
