package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var todos []Todo

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

func todoHandler(w http.ResponseWriter, req *http.Request) {
	//method := "GET"
	method := req.Method
	if method == "POST" {
		body, err := ioutil.ReadAll(req.Body) //now body is a byte array -> we need to convert them (to struct)

		if err != nil {
			fmt.Fprintf(w, "error: %v", err)
			return
		}
		t := Todo{}
		err = json.Unmarshal(body, &t)
		if err != nil {
			fmt.Fprintf(w, "error: %v", err)
		}
		todos = append(todos, t)

		fmt.Printf("body: %#v\n", t)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, "{\"method\": \"%s\"}", method)
		//fmt.Fprintf(w, "Hello %s created todos", method)
		return
	}
	//fmt.Fprintf(w, "Hello %s todos", method)
	w.Header().Set("content-type", "application/json")
	fmt.Fprintf(w, "{\"method\": \"%s\"}", method)
}

func main() {
	fmt.Println("start...")
	http.HandleFunc("/todos", todoHandler)

	http.ListenAndServe(":1234", nil)
	fmt.Println("end...")
}
