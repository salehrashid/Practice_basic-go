package main

func main() {
	name1 := "saleh"

	printName := func() {
		name2 := "rizky"
		println(name2)

		printName2 := func() {
			name3 := "umar"
			println(name3)
		}

		printName2()
	}

	println(name1)
	printName()
}
