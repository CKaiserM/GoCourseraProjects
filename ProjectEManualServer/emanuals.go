package main

import (
	"fmt"
	"net/http"
)

func httpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received:", r.URL.Path)
	fmt.Fprintf(w, "<h1>Welcome to Night Owl</h1>")
}

func main() {
	http.HandleFunc("/", httpHandler)
	fmt.Println("Listening on port 3000...")
	http.ListenAndServe(":3000", nil)
}
