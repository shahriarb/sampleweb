package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world! it is ShaB again")
	//http.NotFound(w, r)
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":9000", nil)
}
