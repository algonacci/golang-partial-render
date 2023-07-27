package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type M map[string]interface{}

func main() {
	// var template, err = template.ParseGlob("views/*")
	// if err != nil {
	// 	panic(err.Error())
	// 	// return
	// }

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"name": "Batman"}
		var template = template.Must(template.ParseFiles(
			"views/index.html",
			"views/_header.html",
			"views/_message.html",
		))

		var err = template.ExecuteTemplate(w, "index", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"name": "Batman"}
		var template = template.Must(template.ParseFiles(
			"views/about.html",
			"views/_header.html",
			"views/_message.html",
		))

		var err = template.ExecuteTemplate(w, "about", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	var port = ":9000"
	fmt.Printf("🚀 [SERVER] is running on port http://localhost%s\n", port)
	http.ListenAndServe(":9000", nil)

}
