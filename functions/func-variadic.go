package main

import "fmt"

func main() {
	total := sumAll(1, 4, 5, 3, 13, 3)
	fmt.Println(total)

	//kalau sudah terlanjur buat array(slice)
	slice := []int{1, 3, 4, 5, 6, 7, 8}
	total = sumAll(slice...)
	fmt.Println(total)

	println(dataString("data 1", "data 2", "data 3"))
	dataInt(1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	dataBoolean(true, false)
}

func sumAll(numbers ...int) int {
	total := 0
	for _, valueAngka := range numbers {
		total += valueAngka
	}
	return total
}

func dataInt(numbers ...int) {
	for _, valueNumber := range numbers {
		println(valueNumber)
	}
}

func dataString(text ...string) string {
	for _, valueString := range text {
		println(valueString)
	}
	return ""
}

func dataBoolean(booleans ...bool) {
	for _, valueBool := range booleans {
		println(valueBool)
	}
}
