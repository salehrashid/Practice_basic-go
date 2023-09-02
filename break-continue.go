package main

func main() {

	println("================ break for loops ================")

	for i := 1; i <= 10; i++ {
		if i == 5 {
			break
		}
		/** akan berhenti diangka 4 **/
		println("angka", i)
	}

	println("================ continue for loops ================")

	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue
		}
		/** angka 5 diignore **/
		println("angka", i)
	}
}
