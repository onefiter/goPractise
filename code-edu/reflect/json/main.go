package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
	Age  int
	Sex  byte `json:"gender"`
}

func main() {
	user := User{
		Name: "钱钟书",
		Age:  57,
		Sex:  1,
	}

	json, _ := json.Marshal(user)
	fmt.Printf("%v", string(json))
}
