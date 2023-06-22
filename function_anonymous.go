package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("ady", blacklist)

	registerUser("root", func(name string) bool {
		return name == "root"
	})
	registerUser("ady", func(name string) bool {
		return name == "root"
	})
}
