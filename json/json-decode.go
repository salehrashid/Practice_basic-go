package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName string `json:"name"`
	Age      int
}

func main() {
	jsonString := `{"Name": "Saleh","Age": 19}`
	jsonData := []byte(jsonString)

	var data User

	//Struct to JSON
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("user :", data.FullName)
	fmt.Println("age  :", data.Age)
}
