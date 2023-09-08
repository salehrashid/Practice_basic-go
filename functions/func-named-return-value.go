package main

func main() {
	namaPertama, namaTengah, namaAkhir := getCompleteName()
	umur, money, motorcycle, marriage := harta()
	println("nama saya adalah", namaPertama, namaTengah, namaAkhir, "umur saya", umur, "tahun", "dan tabungan saya sisa", money, "dan emang saya punya motor ya? =", motorcycle, "emang aku udah nikah? =", marriage)
}

func getCompleteName() (firstName, middleName, lastname string) {
	firstName = "saleh"
	middleName = "rashid"
	lastname = "babsel"

	return
}

func harta() (umur, duit int, punyaMotor, udahNikah bool) {
	umur = 19
	duit = 500000
	punyaMotor = true
	udahNikah = false

	return
}
