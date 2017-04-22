package main

import (
	"html/template"
	"net/http"
)

var err error

func main() {
	http.Handle("/static/", http.FileServer(assetFS()))
	http.HandleFunc("/", userHandle)
	http.ListenAndServe(":8080", nil)
}

func userHandle(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "Spencer", Photo: "/img/image.jpg"}
	t := template.Must(template.New("index.html").ParseFiles("./data/index.html"))
	err = t.Execute(w, user)
	if err != nil {
		panic(err)
	}
}

type User struct {
	Name string
	Photo string
}
