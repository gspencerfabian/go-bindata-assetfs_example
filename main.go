package main

import (
	"html/template"
	"net/http"
)

var err error

func main() {
	http.HandleFunc("/", indexHandle)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(assetFS())))
	http.ListenAndServe(":8080", nil)
}

func indexHandle(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "Spencer", Photo: "static/img/image.jpg"}
	tfile := MustAsset("data/templates/index.html")
	t := template.Must(template.New("index.html").Parse(string(tfile)))
	err = t.Execute(w, user)
	if err != nil {
		panic(err)
	}
}

type User struct {
	Name string
	Photo string
}
