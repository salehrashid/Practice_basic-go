package main

import "fmt"

type Identity struct {
	FullName, Address string
	Age               int
	Kerjaan           bool
	devices           string
}

func main() {

	println("================ cara pertama ================")

	var identity Identity
	identity.FullName = "Saleh Rashid Babsel"
	identity.Age = 19
	identity.Address = "Bogor Kota"
	identity.devices = "Realme C2 " + "Lenovo ThinkBook"
	identity.Kerjaan = true

	fmt.Println(identity)

	println("================ cara kedua ================")

	identity2 := Identity{
		FullName: "Rizky",
		Address:  "Bogor",
		Age:      18,
		Kerjaan:  true,
		devices:  "Asus VivoBook",
	}

	fmt.Println(identity2)

	println("================ cara ketiga ================")

	/**
	disarankan menggunakan cara yang kedua agar lebih mudah dibandingkan
	menggunakan cara ketiga, dikarenakan cara ketiga lebih rentan error
	atau tidak bisa mengisi data secara acak atau tidak berdasarkan urutan
	tipe data yang sudah ditentukan di type declaration.
	*/
	identity3 := Identity{"Umar", "Kuwait", 18, true, "mac m1"}

	fmt.Println(identity3)
}
