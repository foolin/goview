package goview

import (
	"fmt"
	"net/http"
)

func ExampleDefault() {

	/*
	   Project structure:

	   |-- app/views/
	       |--- index.html
	       |--- page.html
	       |-- layouts/
	           |--- footer.html
	           |--- master.html

	*/

	//render index use `index` without `.html` extension, that will render with master layout.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := Render(w, http.StatusOK, "index", M{
			"title": "Index title!",
			"add": func(a int, b int) int {
				return a + b
			},
		})
		if err != nil {
			fmt.Fprintf(w, "Render index error: %v!", err)
		}

	})

	//render page use `page.html` with '.html' will only file template without master layout.
	http.HandleFunc("/page", func(w http.ResponseWriter, r *http.Request) {
		err := Render(w, http.StatusOK, "page.html", M{"title": "Page file title!!"})
		if err != nil {
			fmt.Fprintf(w, "Render page.html error: %v!", err)
		}
	})

	fmt.Println("Listening and serving HTTP on :9090")
	http.ListenAndServe(":9090", nil)
}
