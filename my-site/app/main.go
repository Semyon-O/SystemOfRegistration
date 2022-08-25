package main

import (
	"log"
	"net/http"

	controller "main/app/controllers/controllers"
	url "main/app/controllers/routes"
)

func HandleFunc() {
	http.HandleFunc("/", url.Index)
	http.HandleFunc("/news/", url.News)
	http.HandleFunc("/registration/", url.Registration)
	http.HandleFunc("/user/", controller.NewUser)
	http.HandleFunc("/authorization/", controller.Authorization)
}

func main() {
	HandleFunc()
	log.Println("Запуск веб-сервера на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
