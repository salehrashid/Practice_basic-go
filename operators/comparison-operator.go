package main

func main() {
	name := "saleh"
	name2 := "saleh"

	value1 := 10
	value2 := 20

	var resultValue1 = value1 < value2
	var resultValue2 = value1 > value2
	var resultValue3 = value1 == value2
	var resultValue4 = value1 != value2

	var resultString1 = name != name2
	var resultString2 = name == name2

	println(resultString1, resultString2)
	println(resultValue1, resultValue2, resultValue3, resultValue4)
	println(value1 >= 10 && value2 < 21)
	println(value1 >= 10, value1 <= 10)
}
