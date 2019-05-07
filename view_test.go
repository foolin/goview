package goview

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
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

func ExampleNew() {

	/*
		Project structure:

		|-- app/views/
		    |--- index.tpl
		    |--- page.tpl
		    |-- layouts/
		        |--- footer.tpl
		        |--- head.tpl
		        |--- master.tpl
		    |-- partials/
		        |--- ad.tpl
	*/

	//config
	gv := New(Config{
		Root:      "views",
		Extension: ".tpl",
		Master:    "layouts/master",
		Partials:  []string{"partials/ad"},
		Funcs: template.FuncMap{
			"sub": func(a, b int) int {
				return a - b
			},
			"copy": func() string {
				return time.Now().Format("2006")
			},
		},
		DisableCache: true,
	})

	//Set new instance
	Use(gv)

	//render index use `index` without `.tpl` extension, that will render with master layout.
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

	//render page use `page.tpl` with '.tpl' will only file template without master layout.
	http.HandleFunc("/page", func(w http.ResponseWriter, r *http.Request) {
		err := Render(w, http.StatusOK, "page.tpl", M{"title": "Page file title!!"})
		if err != nil {
			fmt.Fprintf(w, "Render page.html error: %v!", err)
		}
	})

	fmt.Println("Listening and serving HTTP on :9090")
	http.ListenAndServe(":9090", nil)
}
