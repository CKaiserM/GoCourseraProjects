package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func IndexHttpHandler(w http.ResponseWriter, r *http.Request) {
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

func NewsHttpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received:", r.URL.Path)
	//	fmt.Fprintf(w, "<h1>Welcome to Night Owl</h1>")
	t, err := template.ParseFiles("templates/news.html")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%v Server error \n", http.StatusNotFound)
		fmt.Fprintf(w, "Description: %s \n", err)
		return
	}
	var date_now = time.Now().String()

	//fmt.Println(date_now)
	t.Execute(w, date_now)
}

func main() {
	http.HandleFunc("/", IndexHttpHandler)
	http.HandleFunc("/news", NewsHttpHandler)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	http.Handle("/manuals/", http.StripPrefix("/manuals/", http.FileServer(http.Dir("manuals"))))
	fmt.Println("Listening on port 3000...")
	http.ListenAndServe(":3000", nil)
}
