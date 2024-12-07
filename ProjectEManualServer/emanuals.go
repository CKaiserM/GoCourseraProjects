package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func httpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received:", r.URL.Path)
	//	fmt.Fprintf(w, "<h1>Welcome to Night Owl</h1>")
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%v Server error \n", http.StatusNotFound)
		fmt.Fprintf(w, "Description: %s \n", err)
		return
	}
	pages, _ := scanDir("./manuals")
	fmt.Println(pages)
	t.Execute(w, pages)
}

func main() {
	http.HandleFunc("/", httpHandler)
	fmt.Println("Listening on port 3000...")
	http.ListenAndServe(":3000", nil)
}
