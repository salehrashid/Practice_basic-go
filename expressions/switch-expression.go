package main

func main() {
	name := "saleh"

	println("================ conventional switch expression ================")

	switch name {
	case "saleh":
		println("ini saleh")
	case "umar":
		println("ini umar")
	case "rizky":
		println("ini rizky")

	/** sebuah else kalo di if **/
	default:
		println("gak kenal")
	}

	println("================ short switch expression ================")

	switch value := 100; value <= 80 {
	case true:
		println("kau lulus")
	case false:
		println("kau remidi")
	}

	println("================ switch without expression ================")

	length := len(name)
	switch {
	case length > 10:
		println("nama panjang")
	case length >= 5:
		println("nama mayan pendek")
	default:
		println("nama pendek")
	}
}
