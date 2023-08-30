package main

func main() {
	benar := true
	salah := false

	/** && operator harus dua duanya benar **/
	println("========== && operator ==========")
	println(benar && benar)
	println(salah && benar)
	println(benar && salah)
	println(salah && salah)

	/** || (OR, atau) operator, salah satu dari mereka harus benar **/
	println("========== || operator ==========")
	println(benar || benar)
	println(salah || benar)
	println(benar || salah)
	println(salah || salah)

	/** ! operator yang salah jadi benar (dibalik) **/
	println("========== ! operator ==========")
	println(!benar)
	println(!salah)

	println(80 > 80)
}
