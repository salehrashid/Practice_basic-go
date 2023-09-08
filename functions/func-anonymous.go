package main

func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}
	registerUser("saleh", blacklist)

	registerUser("admin", func(name string) bool {
		return name == "admin"
	})

	foo2("saleh", func(nama string) string {
		return nama
	})

	foo3("rashid", func(name string) {
		println(name)
	})

	foo := func(a, b int) int {
		return a + b
	}

	println(foo(3, 3))
}

type Blacklist func(string) bool
type Function func(string) string
type Function2 func(string)

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		println("you are blocked", name)
	} else {
		println("welcome", name)
	}
}

func foo2(name string, function Function) {
	println(function(name))
}

func foo3(name string, function Function2) {
	function(name)
}
