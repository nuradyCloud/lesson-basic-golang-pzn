package main

import "fmt"

func main() {
	name := "Ady"
	counter := 0

	increment := func() {
		name := "Nur"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)

}
