package main

import (
	"errors"
	"fmt"
	"net/http"
	"text/template"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {

	renderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 is %d", sum))
}

func Divide(w http.ResponseWriter, r *http.Request) {
	var x float32
	x = 100.0
	var y float32
	y = 10.0
	f, err := divideValues(x, y)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Cannot divide by 0")
		return
	} else {
		fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", x, y, f))
	}

}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("Cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

func addValues(x, y int) int {
	return x + y
}

func renderTemplate(w http.ResponseWriter, tmpl string) {

	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error while parsing the template", err)
	}
}
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
