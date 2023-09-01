package main

import "fmt"

func main() {
	namaMotor := map[string]string{
		"motor1": "ducati",
		"motor2": "BMW",
		"motor3": "aprilia",
		"motor4": "kawasaki",
	}
	/** mengganti data map yang telah ditentukan **/
	namaMotor["motor2"] = "suzuki"

	/** menambahkan data map **/
	namaMotor["motor5"] = "MV agusta"

	println("================ print data map ================")

	fmt.Println(namaMotor)
	fmt.Println(namaMotor["motor1"])

	println("================ mengetahui panjang data map ================")

	fmt.Println(len(namaMotor))

	println("================ membuat data map baru ================")

	makeMap := make(map[string]string)
	makeMap["motor6"] = "honda"
	fmt.Println(makeMap)

	println("================ menghapus data map yang telah ditentukan ================")

	delete(namaMotor, "motor4")
	fmt.Println(namaMotor)
}
