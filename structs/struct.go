package main

import "fmt"

/*
*
deklarasi struct Identity
*/
type Identity struct {
	FullName, Address string
	Age               int
	Work              bool
	Devices           string
}

/*
*
penerapan struct
*/
func main() {

	println("================ inisialisasi cara pertama ================")

	var identity Identity
	identity.FullName = "Saleh Rashid Babsel"
	identity.Age = 19
	identity.Address = "Bogor Kota"
	identity.Devices = "Realme C2 " + "Lenovo ThinkBook"
	identity.Work = true

	fmt.Println(identity)

	println("================ inisialisasi cara kedua ================")

	identity2 := Identity{
		FullName: "Rizky",
		Address:  "Bogor",
		Age:      18,
		Work:     true,
		Devices:  "Asus VivoBook",
	}

	fmt.Println(identity2)

	println("================ inisialisasi cara ketiga ================")

	/**
	disarankan menggunakan cara yang kedua agar lebih mudah dibandingkan
	menggunakan cara ketiga, dikarenakan cara ketiga lebih rentan error
	atau tidak bisa mengisi data secara acak atau tidak berdasarkan urutan
	tipe data yang sudah ditentukan di type declaration.
	*/
	identity3 := Identity{"Umar", "Kuwait", 18, true, "mac m1"}

	fmt.Println(identity3)

	println("================ inisialisasi cara keempat ================")

	var identity4 = Identity{}
	identity4.FullName = "Saleh Rashid Babsel"
	identity4.Age = 19
	identity4.Address = "Bogor Kota"
	identity4.Devices = "Realme C2 " + "Lenovo ThinkBook"
	identity4.Work = true

	fmt.Println(identity4)
}
