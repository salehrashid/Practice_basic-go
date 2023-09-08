package main

import "strings"

func main() {
	filter := spamFilter
	sayHelloWithFilter("saleh", filter)
	sayHelloWithFilter("Anjing", filter)

	formatCase("HAI", lowerCase)

	paragraph("Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.", desc)

	println(apply(add, 5, 3, 4))
	println(apply(multiply, 5, 3, 4))

	println(foo(identity, "saleh ", "XII RPL B"))
	println(foo(identity, "umar ", "XII DMM"))
	println(foo(identity2, "rizky ", "XII TKJ"))
}

type Filter func(string) string
type Formatter func(string) string
type Description func(string) string
type FuncParam func(int, int, int) int
type Callback func(string, string) string

func sayHelloWithFilter(name string, filter Filter) {
	println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func formatCase(txt string, formatter Formatter) {
	println(formatter(txt))
}

func lowerCase(txt string) string {
	toLowerCase := strings.ToLower(txt)
	return toLowerCase
}

func paragraph(text string, tulisan Description) {
	println(tulisan(text))
}

func desc(txt string) string {
	return txt
}

func apply(funcParam FuncParam, a, b, c int) int {
	return funcParam(a, b, c)
}

func add(x, y, z int) int {
	return x + y + z
}

func multiply(x, y, z int) int {
	return x * y * z
}

func foo(callback Callback, nama, kelas string) string {
	return callback(nama, kelas)
}

func identity(name, class string) string {
	return name + class
}

func identity2(name, class string) string {
	return name + class
}
