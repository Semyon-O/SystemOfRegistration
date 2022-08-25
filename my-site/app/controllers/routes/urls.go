package routes

import (
	"fmt"
	"html/template"
	"net/http"
)

func Index(p http.ResponseWriter, r *http.Request) {
	page, err := template.ParseFiles("./app/templates/index.html")
	if err != nil {
		fmt.Println(err)
	}
	page.Execute(p, nil)
}

func News(p http.ResponseWriter, r *http.Request) {
	page, err := template.ParseFiles("./app/templates/news.html")
	if err != nil {
		fmt.Println(err)
	}

	page.Execute(p, nil)
}

func Registration(p http.ResponseWriter, r *http.Request) {
	page, err := template.ParseFiles("./app/templates/registration.html")
	if err != nil {
		fmt.Println(err)
	}

	page.Execute(p, nil)
}
