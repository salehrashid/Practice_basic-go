package main

/**
https://dasarpemrogramangolang.novalagung.com/A-interface.html
*/

/**
Interface adalah kumpulan definisi method yang tidak memiliki isi
(hanya definisi saja), yang dibungkus dengan nama tertentu.

Interface merupakan tipe data. Nilai objek bertipe interface zero
value-nya adalah nil. Interface mulai bisa digunakan jika sudah
ada isinya, yaitu objek konkret yang memiliki definisi method
minimal sama dengan yang ada di interface-nya.
*/

func main() {
	vehicleBraker()
	vehicleStarter()
}

type Vehicle interface {
	brakeSystem() string
	startSystem() string
}

type Train struct {
	Brake string
	Start string
}

type Motorcycle struct {
	Brake string
	Start string
}

type Car struct {
	Brake string
	Start string
}

// implement interface

func (train Train) brakeSystem() string {
	return "rem kereta dengan " + train.Brake
}

func (car Car) brakeSystem() string {
	return "rem mobil dengan " + car.Brake
}

func (motorcycle Motorcycle) brakeSystem() string {
	return "rem motor dengan " + motorcycle.Brake
}

func (train Train) startSystem() string {
	return "nyalakan kereta dengan " + train.Start
}

func (car Car) startSystem() string {
	return "nyalakan mobil dengan " + car.Start
}

func (motorcycle Motorcycle) startSystem() string {
	return "nyalakan motor dengan " + motorcycle.Start
}

func printBrake(braker []Vehicle) {
	for _, b := range braker {
		println(b.brakeSystem())
	}
}

func printStarter(starter []Vehicle) {
	for _, s := range starter {
		println(s.startSystem())
	}
}

func vehicleBraker() {
	kereta := Train{Brake: "tarik tuas rem"}
	mobil := Car{Brake: "injek pedal rem"}
	motor := Motorcycle{Brake: "injek tuas rem"}

	printBrake([]Vehicle{kereta, mobil, motor})
}

func vehicleStarter() {
	kereta := Train{Start: "putar tuas dan blabla"}
	mobil := Car{Start: "putar kunci"}
	motor := Motorcycle{Start: "pencet tombol"}

	printStarter([]Vehicle{kereta, mobil, motor})
}
