package main

import (
	"net/http"
)

// in order for a function to respond to a request from a web browser it needs two parameters
func Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "This is the home page")

	//-----
	renderTemplate(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	// sum := addValues(2, 2)
	// // fmt.Printf(w, "This is the about page")
	// _, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 is %d", sum))

	// -----------
	renderTemplate(w, "about.page.html")
}
