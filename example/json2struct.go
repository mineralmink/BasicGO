package main

import (
	"encoding/json"
	"fmt"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

func main() {
	data := []byte(`{
		"id": 1,
		"title": "pay credit card",
		"status": "active"
	}`)

	t := Todo{}
	err := json.Unmarshal(data, &t) //change json to struct (where)

	fmt.Printf("%#v\n", t)
	fmt.Println(err)
}
