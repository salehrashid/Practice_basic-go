package main

func main() {
	name, age, status, marriage := selfIdentity()
	nameIntro, ageIntro, statusIntro, marriageIntro := introWords()
	println(nameIntro, name, ageIntro, age, statusIntro, status, marriageIntro, marriage)
}

func selfIdentity() (string, int, string, bool) {
	return "saleh rashid", 19, "pelajar", false
}

func introWords() (string, string, string, string) {
	return "perkenalkan nama saya", "umur saya", "dan saya masih", "dan apakah kamu sudah menikah?"
}
