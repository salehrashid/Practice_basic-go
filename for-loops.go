package main

import "fmt"

func main() {

	println("================ conventional for loops ================")

	for i := 1; i <= 25; i++ {
		println(i, ":")
	}

	println("================ print data array with enhanced for loops (foreach) ================")

	namaMotorArray := [3]string{
		"MV Agusta",
		"Aprilia",
		"Kawasaki",
	}

	for _, value := range namaMotorArray {
		fmt.Println(value)
	}

	println("================ print data slice with enhanced for loops (foreach) ================")

	namaMobilSlice := []string{
		"Brabus",
		"Mazda",
		"Supra",
	}

	for _, value := range namaMobilSlice {
		fmt.Println(value)
	}

	println("================ print data map with enhanced for loops (foreach) ================")

	namaPesawatMap := map[string]string{
		"pesawat1": "Sukhoi Su-35S",
		"pesawat2": "Dassault Rafale",
		"pesawat3": "Lockheed Martin F-22 Raptor",
	}

	for key, value := range namaPesawatMap {
		fmt.Println(key, "=", value)
	}

	println("================ angka genap ================")

	for i := 0; i <= 50; i += 2 {
		println(i)
	}

	println("================ angka ganjil ================")

	for i := 0; i <= 50; i++ {
		if i%2 != 0 {
			println(i)
		}
	}

	println("================ angka ganjil dan genap ================")

	for i := 0; i <= 50; i++ {
		if i%2 == 0 {
			println(i, "genap")
		} else {
			println(i, "ganjil")
		}
	}
}
