package main

func main() {
	runApp(true)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("application force closed")
	}
	println("running application")
}

func endApp() {
	println("application compiled")
}
