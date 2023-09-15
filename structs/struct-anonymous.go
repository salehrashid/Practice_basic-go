package main

import "fmt"

/**
https://dasarpemrogramangolang.novalagung.com/A-struct.html
*/

/*
*
An anonymous struct is just like a normal struct,
but it is defined without a name and therefore
cannot be referenced elsewhere in the code.
*/

/*
*
Anonymous struct adalah struct yang tidak
dideklarasikan di awal sebagai tipe data baru,
melainkan langsung ketika pembuatan objek.
Teknik ini cukup efisien untuk pembuatan
variabel objek yang struct-nya hanya dipakai sekali.
*/

type Car struct {
	name       string
	engineType string
}

func main() {
	//anonymous struct
	car := struct {
		name       string
		engineType string
	}{
		name:       "dodge",
		engineType: "Hemi V8",
	}
	fmt.Println("anonymous struct:", car)

	//struct biasa
	car2 := Car{
		name:       "porsche",
		engineType: "flat 6",
	}
	fmt.Println("struct biasa:", car2)

	//anonymous struct dengan keyword var
	var car3 struct {
		name       string
		engineType string
	}
	car3.name = "toyota supra"
	car3.engineType = "inline 6"

	fmt.Println("anonymous struct dengan keyword var:", car3)

	//anonymous struct dengan keyword var sekaligus inisialisasi
	var car4 = struct {
		name       string
		engineType string
	}{
		name:       "skyline r34",
		engineType: "inline 6",
	}

	fmt.Println("anonymous struct dengan keyword var sekaligus inisialisasi: ", car4)
}
