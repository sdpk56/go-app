package main

import (
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	// Render the home html page from static folder
	http.ServeFile(w, r, "static/home.html")
}

func main() {

	http.HandleFunc("/home", homePage)
	err := http.ListenAndServe("0.0.0.0:1234", nil)
	if err != nil {
		log.Fatal(err)
	}
}
