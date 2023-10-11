package main

func main() {
	var name string
	var age int
	var benar bool
	//:= itu pengganti var agar lebih singkat
	hp := "realme c2"
	//bisa juga dengan membuat sebuah variable yang sekaligus banyak
	var (
		firstName = "saleh"
		lastName  = "rashid"
	)
	println(firstName, lastName)

	hp = "redmi note 10"
	println(hp)

	name = "saleh"
	println(name)
	name = "saleh rashid"
	println(name)

	age = 19
	println(age)

	benar = true
	println(benar)

	//bisa mendeklarasi dan inisialisasi beberapa variable sekaligus
	var1, var2, var3 := 22, "w", true
	println(var1, var2, var3)
}
