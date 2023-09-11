package main

import "fmt"

type FavoriteAnime struct {
	Name, Genre, Waifu string
	Eps, Season        int
	Favorite           bool
}

type FavoriteMovie struct {
	Name, Genre string
	Duration    float32
}

func (anime FavoriteAnime) favoriteAnime() {
	println("My favorite anime is called", anime.Name, "or wandering witch: the journey of elaina.", "This anime have an", anime.Eps, "have season", anime.Season, ".", "And my waifu is", anime.Waifu, ".", "Someone ask like this: \"is this your favorite anime?\" and i said", anime.Favorite)
}

func (movie FavoriteMovie) favoriteMovie() {
	fmt.Println("My favorite movie is", movie.Name, ".", "The duration of this movie is", movie.Duration, ".", "And have a genre", movie.Genre, ".")
}

func main() {
	var animes FavoriteAnime
	animes.Name = "Majo no tabi tabi"
	animes.Genre = "Magic, Adventure"
	animes.Waifu = "Elaina"
	animes.Eps = 12
	animes.Season = 1
	animes.Favorite = true
	animes.favoriteAnime()

	var movies FavoriteMovie
	movies.Name = "dunkirk"
	movies.Genre = "Drama, Adventure, War"
	movies.Duration = 1.46
	movies.favoriteMovie()
}
