# GoRice

[![GoDoc Widget]][GoDoc] 

goview support for go.rice

## Install
```bash

go get -u github.com/foolin/goview

go get -u github.com/foolin/goview/supports/gorice

```

### Example


```go

package main

import (
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/gorice"
	"net/http"
)

func main() {

	//static
	staticBox := rice.MustFindBox("static")
	staticFileServer := http.StripPrefix("/static/", http.FileServer(staticBox.HTTPBox()))
	http.Handle("/static/", staticFileServer)

	//new view engine
	gv := gorice.New(rice.MustFindBox("views"))
	//set engine for default instance
	goview.Use(gv)

	//render index use `index` without `.tpl` extension, that will render with master layout.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := goview.Render(w, http.StatusOK, "index", goview.M{
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
		err := goview.Render(w, http.StatusOK, "page.html", goview.M{"title": "Page file title!!"})
		if err != nil {
			fmt.Fprintf(w, "Render page.html error: %v!", err)
		}
	})

	fmt.Println("Listening and serving HTTP on :9090")
	http.ListenAndServe(":9090", nil)
}

```

Project structure:
```go
|-- app/views/
    |--- index.html          
    |--- page.html
    |-- layouts/
        |--- footer.html
        |--- master.html
|-- app/static/  
    |-- css/
        |--- bootstrap.css   	
    |-- img/
        |--- gopher.png

See in "examples/gorice" folder
```

[gorice example](https://github.com/foolin/goview/tree/master/_examples/gorice)


           
## More examples

See [_examples/](https://github.com/foolin/goview/blob/master/_examples/) for a variety of examples.

[GoDoc]: https://godoc.org/github.com/foolin/goview/supports/gorice
[GoDoc Widget]: https://godoc.org/github.com/foolin/goview/supports/gorice?status.svg
