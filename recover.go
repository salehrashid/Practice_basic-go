package main

import "fmt"

func main() {
	runApplication(false)
	crash()
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

func run() {
	err := recover()
	fmt.Println("nangkap error:", err)
	println("habis")
}

func crash() {
	defer run()
	panic("error bang")
}
