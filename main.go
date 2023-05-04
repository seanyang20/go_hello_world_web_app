package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var portNumber = ":8080"

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

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

// // if your function name starts with a capital letter, it means you can cann it from outside the package execpt
// // except the main package
// // having it lowercase makes the function private
// func addValues(x, y int) int {
// 	return x + y
// }

// func Divide(w http.ResponseWriter, r *http.Request) {
// 	f, err := divideValues(100.0, 0.0)
// 	if err != nil {
// 		fmt.Fprintf(w, "Cannot divide by 0")
// 		return
// 	}

// 	fmt.Fprintf(w, fmt.Sprintf("%f divded by %f is %f", 100.0, 0.0, f))
// }

// func divideValues(x, y float32) (float32, error) {
// 	if y <= 0 {
// 		err := errors.New("cannot divide by zero")
// 		return 0, err
// 	}
// 	result := x / y
// 	return result, nil
// }

func main() {

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	n, err := fmt.Fprintf(w, "Hello, world!")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	// })

	// ---------
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	// http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

}
