package main

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
	for _, allAnime := range listAnime {
		println(allAnime.name, allAnime.eps)
	}
}
