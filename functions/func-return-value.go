package main

func main() {
	println(getHello("saleh"))
	println(getHello(""))
	println(age(11))
}

func getHello(name string) string {
	if name == "" {
		return "kenalan dong"
	} else {
		return "hai " + name
	}
}

func age(umur int) string {
	if umur <= 10 {
		return "kamu masih anak kecil:)"
	} else {
		return "kamu udah gede"
	}
}
