package main

/**
blank identifier adalah identifier yang digunakan untuk menangani nilai
yang tidak ingin digunakan atau disimpan. Blank identifier
direpresentasikan dengan underscore (_) sebagai namanya
*/

func main() {
	firstName, _ := blank()
	println(firstName)
}

func blank() (string, string) {
	return "saleh", "rashid"
}
