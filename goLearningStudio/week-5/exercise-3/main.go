package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type MyPage struct {
	Title string
	Body  string
}

func main() {
	http.HandleFunc("/", renderTemplate)
	http.ListenAndServe(":80", nil)
}

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("A page visitor has arrived!")

	pageinput := MyPage{
		Title: "My Go Language Webpage!",
		Body:  "Hello there, Gophers! Welcome to my landing page.",
	}

	parsedTemplate, _ := template.ParseFiles("index.html")
	err := parsedTemplate.Execute(w, pageinput)
	if err != nil {
		fmt.Println("Error loading template: ", err)
		return
	}
}
