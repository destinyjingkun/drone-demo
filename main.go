package main

import "fmt"
import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func main() {
	fmt.Println("server start")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":20080", nil)
}
