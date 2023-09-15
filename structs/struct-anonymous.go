package main

import "fmt"

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
	name   string
	engine string
}

func main() {
	//anonymous struct
	car := struct {
		name   string
		engine string
	}{
		name:   "dodge",
		engine: "Hemi V8",
	}
	fmt.Println("anonymous struct:", car)

	//struct biasa
	car2 := Car{
		name:   "porsche",
		engine: "flat 6",
	}
	fmt.Println("struct biasa:", car2)
}
