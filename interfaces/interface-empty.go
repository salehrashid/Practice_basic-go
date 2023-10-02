package main

import "fmt"

/**
https://dasarpemrogramangolang.novalagung.com/A-interface-kosong.html
*/

/**
Interface kosong atau empty interface yang dinotasikan dengan interface{}
atau any, merupakan tipe data yang sangat spesial. Variabel bertipe ini
bisa menampung segala jenis data, bahkan array, pointer, apapun. Tipe
data dengan konsep ini biasa disebut dengan dynamic typing.
*/

func main() {
	var dynmcData interface{}

	dynmcData = 1                         //integer
	dynmcData = 1.0                       //float
	dynmcData = nil                       //nil
	dynmcData = true                      //boolean
	dynmcData = "string"                  //string
	dynmcData = [1]int{9}                 //array
	dynmcData = []string{"apalah"}        //slice
	dynmcData = map[string]int{"data": 1} //map

	fmt.Println(dynmcData)

	var data map[string]interface{}

	data = map[string]interface{}{
		"nama":   "saleh",
		"alamat": "indo",
		"umur":   19,
		"animeFavorit": []string{
			"majo no tabitabi",
			"grand blue",
			"zom100",
		},
	}

	for _, value := range data {
		fmt.Println(value)
	}

	/**
	Tipe any merupakan alias dari interface{}, keduanya adalah sama.
	*/

	var data2 map[string]any

	data2 = map[string]any{
		"nama":   "saleh",
		"alamat": "indo",
		"umur":   19,
		"filmFavorit": []string{
			"top gun maverick",
			"tenet",
			"interstellar",
		},
	}

	for _, value := range data2 {
		fmt.Println(value)
	}
}
