package main

import (
	"net/http"
	"fmt"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header.Get("User-Agent"))
	fmt.Println()
	fmt.Println("request method = "+r.Method)
	fmt.Println()
	fmt.Println(r.GetBody)
	fmt.Fprintln(w, "hello world")
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe("127.0.0.1:8000", nil)
}


