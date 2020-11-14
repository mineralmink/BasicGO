package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func todoHandler(w http.ResponseWriter, req *http.Request) {
	//method := "GET"
	method := req.Method
	if method == "POST" {
		body, err := ioutil.ReadAll(req.Body) //now body is a byte array -> we need to convert them (to struct)
		if err != nil {
			fmt.Fprintf(w, "error: %v", err)
			return
		}
		fmt.Printf("body: %s\n", body)

		fmt.Fprintf(w, "Hello %s created todos", method)
		return

	}
	fmt.Fprintf(w, "Hello %s todos", method)
}

func main() {
	fmt.Println("start...")
	http.HandleFunc("/todos", todoHandler)

	http.ListenAndServe(":1234", nil)
	fmt.Println("end...")
}
