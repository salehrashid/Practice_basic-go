package main

import "fmt"

/**
Embedded struct adalah mekanisme untuk menempelkan
sebuah struct sebagai properti struct lain.
*/

type Anime struct {
	Name   string
	Genre  string
	Rating float32
}

type AnimeMovie struct {
	Name     string
	Genre    string
	Rating   float32
	Duration float32
	Anime
}

func main() {
	anime := AnimeMovie{}
	anime.Name = "majo no tabitabi."
	anime.Genre = "magic, adventure, fantasy."
	anime.Rating = 7.55

	anime.Anime.Name = "suzume no tojimari."
	anime.Anime.Genre = "adventure, fantasy."
	anime.Anime.Rating = 8.37
	anime.Duration = 2.1

	fmt.Println("anime season:", anime.Name, anime.Genre, anime.Rating)
	fmt.Println("anime movie:", anime.Anime.Name, anime.Anime.Genre, anime.Anime.Rating)
}
