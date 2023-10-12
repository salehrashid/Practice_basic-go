package main

import (
	"encoding/json"
	"fmt"
)

type User2 struct {
	FullName string
	Age      int
}

func main() {
	user := User2{
		FullName: "Rashid",
		Age:      63,
	}
	jsonData, err := json.Marshal(user)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(jsonData))
}
