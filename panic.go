package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}
	fmt.Println("Application selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Application ERORR")
	}
	fmt.Println("Application is running")
}

func main() {
	runApp(true)
	fmt.Println("Ady")
}
