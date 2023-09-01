package main

import (
	"fmt"
	"strconv"
)

func main() {
	var intToInt int32 = 4000
	var intToInt2 = int16(intToInt)
	println(intToInt2)

	floatToInt := 1000.9
	floatToInt2 := int(floatToInt)
	println(floatToInt2)

	var intToFloat = 20
	var intToFloat2 = float32(intToFloat)
	fmt.Printf("%.2f", intToFloat2)

	var intToString = 9
	println(strconv.Itoa(intToString))

	var stringToInt = "saleh"
	println(strconv.Atoi(stringToInt))
}
