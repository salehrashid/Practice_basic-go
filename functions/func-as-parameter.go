package main

import "strings"

func main() {
	filter := spamFilter
	sayHelloWithFilter("saleh", filter)
	sayHelloWithFilter("Anjing", filter)

	formatCase("HAI", lowerCase)

	paragraph("sasasasiashaish", desc)

	println(apply(add, 5, 3, 4))
	println(apply(multiply, 5, 3, 4))

	println(foo(identity, "saleh ", "XII RPL B"))
}

type Filter func(string) string

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

type Formatter func(string) string

func formatCase(txt string, formatter Formatter) {
	println(formatter(txt))
}
func lowerCase(txt string) string {
	toLowerCase := strings.ToLower(txt)
	return toLowerCase
}

type Tulisan func(string) string

func paragraph(text string, tulisan Tulisan) {
	println(tulisan(text))
}

func desc(txt string) string {
	return txt
}

type FuncParam func(int, int, int) int

func apply(funcParam FuncParam, a, b, c int) int {
	return funcParam(a, b, c)
}

func add(x, y, z int) int {
	return x + y + z
}

func multiply(x, y, z int) int {
	return x * y * z
}

type Callback func(string, string) string

func foo(callback Callback, nama, kelas string) string {
	return callback(nama, kelas)
}

func identity(name, class string) string {
	return name + class
}
