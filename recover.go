package main

import "fmt"

func main() {
	runApplication(false)
}

func runApplication(error bool) {
	defer endApplication()
	if error {
		panic("application turned off immediately")
	}
	fmt.Println("running app")
}

func endApplication() {
	err := recover()
	if err != nil {
		fmt.Println("something went wrong:", err)
	}
	fmt.Println("application compiled successfully")
}
