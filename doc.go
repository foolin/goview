//Package goview a lightweight, minimalist and idiomatic template library
//based on golang html/template for building Go web application.
//
//Example:
//
//	package main
//
//	import (
//		"fmt"
//		"github.com/foolin/goview"
//		"net/http"
//	)
//
//	func main() {
//
//		//render index use `index` without `.html` extension, that will render with master layout.
//		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//			err := goview.Render(w, http.StatusOK, "index", goview.M{
//				"title": "Index title!",
//				"add": func(a int, b int) int {
//					return a + b
//				},
//			})
//			if err != nil {
//				fmt.Fprintf(w, "Render index error: %v!", err)
//			}
//
//		})
//
//		//render page use `page.html` with '.html' will only file template without master layout.
//		http.HandleFunc("/page", func(w http.ResponseWriter, r *http.Request) {
//			err := goview.Render(w, http.StatusOK, "page.html", goview.M{"title": "Page file title!!"})
//			if err != nil {
//				fmt.Fprintf(w, "Render page.html error: %v!", err)
//			}
//		})
//
//		fmt.Println("Listening and serving HTTP on :9090")
//		http.ListenAndServe(":9090", nil)
//
//	}
//
//
//Project structure:
//
//   |-- app/views/
//	   |--- index.html
//	   |--- page.html
//	   |-- layouts/
//		   |--- footer.html
//		   |--- master.html
//
//Learn more at https://github.com/foolin/goview
//
//
//
//================== Supports ==================
//
//Ginview for Gin framework:
//https://godoc.org/github.com/foolin/goview/supports/ginview
//
//Echoview for Echo framework:
//https://godoc.org/github.com/foolin/goview/supports/echoview
//
//Gorice for Go.rice:
//https://godoc.org/github.com/foolin/goview/supports/gorice
//
//Examples:
//https://github.com/foolin/goview/_examples
package goview
