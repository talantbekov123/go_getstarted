package main

import (
	"fmt"
	"net/http"
	"html/template"

	//local libraries
	"./models"
)

var posts []models.Post


func writeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/write.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "write", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", posts)
}

func savePostHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	content := r.FormValue("content")
	
	posts = append(posts, models.Post{title,content})

	http.Redirect(w, r, "/", 302)
}

func main() {
	fmt.Println("Working on localhost")

	http.HandleFunc("/write", writeHandler)
	http.HandleFunc("/savePost", savePostHandler)
	http.HandleFunc("/", indexHandler)

	http.ListenAndServe(":4011", nil)
}