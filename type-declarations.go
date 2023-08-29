package main

import (
	"fmt"
	"strconv"
)

type tipeStr string

func (val tipeStr) ConvToInt() int {
	intVal, _ := strconv.Atoi(string(val))
	return intVal
}

type tipeInt int

func (value tipeInt) ConvToStr() string {
	return strconv.Itoa(int(value))
}

func main() {
	var contoh string
	contoh = "123"
	newContoh := tipeStr(contoh)
	fmt.Println(newContoh.ConvToInt(), "from string to int")

	var example int
	example = 123
	newExample := tipeInt(example)
	fmt.Println(newExample.ConvToStr(), "from int to string")

	typeDeclaration()
}

func typeDeclaration() {
	type Tulisan string
	type Boolean bool
	type Angka int

	var alamat Tulisan = "bogor"       //var alamat string = "bogor"
	var udahNikahBelom Boolean = false //var udahNikahBelom bool = false
	var umur Angka = 19                //var umur int = 19

	println(udahNikahBelom)
	println(alamat)
	println(umur)
}
