package main

import "fmt"

func main() {
	total := sumAll(10, 20, 30, 40, 50)
	fmt.Println(total)

	sliceNumber := []int{60, 70, 80, 90, 100}
	total = sumAll(sliceNumber...)
	fmt.Println(total)
}

func sumAll(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}

	return total
}
