package main

import "fmt"

func main() {
	fmt.Println("satu =", 1)
	fmt.Println("satu koma dua =", 1.2)

	println("==============================")

	//2 karakter
	var angka1 uint8 = 90

	//4 karakter
	var angka2 uint16 = 9000

	//9 karakter
	var angka3 uint32 = 900000000

	//19 karakter
	var angka4 uint64 = 9000000000000000000

	//19 katakter
	var angka5 uint = 9000000000000000000

	println("uint8 =", angka1)
	println("uint16 =", angka2)
	println("uint32 =", angka3)
	println("uint64 =", angka4)
	println("uint =", angka5)

	println("==============================")

	//2 karakter
	var number int8 = 90

	//4 karakter
	var number2 int16 = 9000

	//9 karakter
	var number3 int32 = 900000000

	//19 karakter
	var number4 int64 = 9000000000000000000

	println("int8 =", number)
	println("int16 =", number2)
	println("int32 =", number3)
	println("int64 =", number4)
}
