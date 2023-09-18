package main

import "fmt"

/**
https://dasarpemrogramangolang.novalagung.com/A-struct.html
*/

type Animes struct {
	name string
	eps  int
}

func main() {
	listAnime := []Animes{
		{name: "kanojo okarishimasu", eps: 12},
		{name: "majo no tabitabi", eps: 12},
		{name: "jujutsu kaisen", eps: 24},
	}

	println("================ struct slice biasa ================")

	for _, allAnime := range listAnime {
		println(allAnime.name, allAnime.eps)
	}

	//anonymous struct dengan slice
	listMovie := []struct {
		name     string
		duration float32
	}{
		{name: "dunkirk", duration: 1.46},
		{name: "tenet", duration: 2.30},
		{name: "interstellar", duration: 2.49},
	}

	println("================ struct slice anonymous ================")

	for _, allMovie := range listMovie {
		fmt.Println(allMovie.name, allMovie.duration)
	}
}
