package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Ady",
		"address": "Semarang",
	}
	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-Lang Dasar"
	book["author"] = "Ady"
	book["ops"] = "kena deh"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ops")
	fmt.Println(book)
	fmt.Println(len(book))
}
