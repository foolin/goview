/*
 * Copyright 2018 Foolin.  All rights reserved.
 *
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

import (
	"fmt"
	"github.com/foolin/goview"
	"html/template"
	"net/http"
	"time"
)

func main() {

	gv := goview.New(goview.Config{
		Root:      "views",
		Extension: ".tpl",
		Master:    "layouts/master",
		Partials:  []string{"partials/ad"},
		Funcs: template.FuncMap{
			"safeHTML": func(v string) template.HTML {
				return template.HTML(v)
			},
			"copy": func() string {
				return time.Now().Format("2006")
			},
		},
		DisableCache: true,
	})

	//Set new instance
	goview.Use(gv)

	rawContent := `This is <b>HTML</b> content! Posted on <time datetime="2019-05-16 01:02:03">May 16</time> by foolin.`

	//render index use `index` without `.tpl` extension, that will render with master layout.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := goview.Render(w, http.StatusOK, "index", goview.M{
			"Title":       "Index title!",
			"HtmlContent": template.HTML(rawContent),
			"RawContent":  rawContent,
			"tempConvertHTML": func(v string) template.HTML {
				return template.HTML(v)
			},
		})
		if err != nil {
			fmt.Fprintf(w, "Render index error: %v!", err)
		}

	})

	//render page use `page.tpl` with '.tpl' will only file template without master layout.
	http.HandleFunc("/page", func(w http.ResponseWriter, r *http.Request) {
		err := goview.Render(w, http.StatusOK, "page.tpl", goview.M{"Title": "Page file title!!"})
		if err != nil {
			fmt.Fprintf(w, "Render page.html error: %v!", err)
		}
	})

	fmt.Println("Listening and serving HTTP on :9090")
	http.ListenAndServe(":9090", nil)
}
