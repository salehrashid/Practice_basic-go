package main

import (
	"fmt"
)

/**
https://dasarpemrogramangolang.novalagung.com/A-struct.html
*/

/**
Nested struct adalah anonymous struct yang di-embed ke sebuah struct.
Deklarasinya langsung di dalam struct peng-embed.

Teknik ini biasa digunakan ketika decoding data json yang struktur
datanya cukup kompleks dengan proses decode hanya sekali.
*/

type Person struct {
	PersonalInfo struct {
		Name string
		Age  int
	}
	Grade   int
	Hobbies []string
}

type Streaming struct {
	Anime struct {
		Title    []string
		Favorite string
	}
	Movie struct {
		Title    []string
		Favorite string
	}

	SpentInternet float32
	SpentWeek     int
}

func main() {
	personalInformation()
	watch()
}

func personalInformation() {
	s := Person{
		PersonalInfo: struct {
			Name string
			Age  int
		}{
			Name: "saleh",
			Age:  19,
		},
		Grade:   3,
		Hobbies: []string{"music", "game", "badminton", "workout"},
	}

	fmt.Println(s)
}

func watch() {
	streaming := Streaming{
		Anime: struct {
			Title    []string
			Favorite string
		}{
			Title:    []string{"majo no tabitabi", "NieR Automata", "grand blue"},
			Favorite: "NieR Automata",
		},
		Movie: struct {
			Title    []string
			Favorite string
		}{
			Title:    []string{"dunkirk", "inception", "tenet", "interstellar"},
			Favorite: "dunkirk",
		},
		SpentInternet: 2.00,
		SpentWeek:     2,
	}

	fmt.Println(streaming)
}
