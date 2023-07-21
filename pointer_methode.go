package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	ady := Man{"Ady"}
	ady.Married()

	fmt.Println(ady.Name)
}
