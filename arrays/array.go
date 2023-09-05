package main

import "fmt"

func main() {
	/** membuat array dengan panjang kapasitas 3 element saja dan bertipe string **/

	println("================ membuat array ================")

	var nama [3]string

	nama[0] = "saleh"
	nama[1] = "rashid"
	nama[2] = "babsel"

	println(nama[0])
	println(nama[1])
	println(nama[2])

	/** [...] = untuk membuat array dengan panjang kapasitas element yang tak terhingga **/
	var value = [...]int{
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
	}

	/** untuk mengubah value dari element ke 2 dengan index 1 **/
	value[1] = 5

	println("================ merubah value array menjadi angka 5 ================")

	fmt.Println(value)

	var carName = [3]string{
		"brabus",
		"chevy",
		"shelby",
	}

	println("================ print data array ================")

	fmt.Println(carName)

	println("================ print panjang data array ================")

	fmt.Println(len(carName))

	var smartphoneBrand = [...]string{
		"realme",
		"xiaomi",
		"oppo",
		"samsung",
		"iPhone",
		"vivo",
	}

	println("================ print data array dengan for loops ================")

	for i := 0; i < len(smartphoneBrand); i++ {
		fmt.Println(smartphoneBrand[i])
	}

	println("================ print data array dengan enhanced for loops (foreach) ================")

	//foreach, jika tidak mau print index maka cukup di ignore saja menggunakan underscore "_" atau blank identifier
	for index, value := range smartphoneBrand {
		fmt.Println("index", index, "=", "value", value)
	}
}
