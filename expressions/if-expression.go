package main

func main() {

	println("================ conventional if else statement ================")

	var nilai = 100
	if nilai <= 60 {
		println("kamu lulus ujian")
	} else if nilai <= 70 {
		println("kamu remidial")
	} else {
		println("harus ikut ujaian ulang")
	}

	println("================ short if else statement ================")

	if value := 10; value < 5 {
		println("panjang")
	} else {
		println("pendek")
	}
}
