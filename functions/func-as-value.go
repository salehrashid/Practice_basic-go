package main

func main() {
	result := getSayonara("saleh")
	println(result)
	println(getSayonara("saleh"))
}

func getSayonara(name string) string {
	return "Sayonara " + name
}
