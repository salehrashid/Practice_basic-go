package main

/**
https://dasarpemrogramangolang.novalagung.com/A-pointer.html
*/

/**
Pointer adalah reference atau alamat memori. Variabel pointer berarti
variabel yang berisi alamat memori suatu nilai. Sebagai contoh sebuah
variabel bertipe integer memiliki nilai 4, maka yang dimaksud pointer
adalah alamat memori di mana nilai 4 disimpan, bukan nilai 4 itu
sendiri.
*/

func main() {
	/**
	Variabel bertipe pointer ditandai dengan adanya tanda asterisk (*)
	tepat sebelum penulisan tipe data ketika deklarasi.
	*/
	var string1 = "saleh"
	var string2 = &string1

	println("value nya  :", string1)
	println("alamat nya :", &string2) //akan disimpan dialamt memori tersebut 0xc00004a718
}
