package main

import "fmt"

/**
https://dasarpemrogramangolang.novalagung.com/A-struct.html
*/

type ParentStruct struct {
	Property1 string
	Property2 string
}

type ChildStruct ParentStruct

func main() {
	properties := ChildStruct{
		Property1: "apalah",
		Property2: "ya begitulah",
	}

	fmt.Println(properties)
}
