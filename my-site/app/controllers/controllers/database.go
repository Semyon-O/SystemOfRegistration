package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func NewUser(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	phone := r.FormValue("phone")

	db, err := sql.Open("sqlite3", "./app/site.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	query := fmt.Sprintf("INSERT INTO Users (email, phone, password) VALUES ('%s', '%s', '%s')", email, phone, password)
	result, err := db.Exec(query)

	if err != nil {
		panic(err)
	}
	fmt.Println("Email: %s, password: %s, phone: %s", email, password, phone)
	fmt.Println(result.RowsAffected())
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Authorization(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	fmt.Println(email, password)

	db, err := sql.Open("sqlite3", "./app/site.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := fmt.Sprintf("select * from Users where email = '%s' AND password = '%s'", email, password)
	row := db.QueryRow(query)
	user := User{}
	err = row.Scan(&user.id, &user.email, &user.phone, &user.password)
	if err != nil {
		fmt.Print(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		fmt.Println(user.id, user.phone, user.email, user.password)
		fmt.Println(row)
		http.Redirect(w, r, "/news", http.StatusSeeOther)
	}
}
