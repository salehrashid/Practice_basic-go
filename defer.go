package main

func main() {
	defer println(hello())

	div := 0
	println(1 / div)

	//println("print pertama")
}

func hello() string {
	return "print kedua"
}
