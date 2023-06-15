package main

import "fmt"

func main() {
	type noKTP string
	type married bool

	var noKTPAdy noKTP = "33741234567297987"
	var marriedStatus married = true

	fmt.Println(noKTPAdy)
	fmt.Println(marriedStatus)
}
