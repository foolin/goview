# goview

[![GoDoc Widget]][GoDoc] [![Travis Widget]][Travis] [![GoReportCard Widget]][GoReportCard] 

Goview is a lightweight, minimalist and idiomatic template library based on golang [html/template](https://golang.org/pkg/html/template/) for building Go web application.

## Contents

- [Install](#install)
- [Features](#features)
- [Docs](#docs)
- [Supports](#supports)
    - [Gin Framework](https://github.com/foolin/goview/tree/master/supports/ginview)
    - [Echo Framework](https://github.com/foolin/goview/tree/master/supports/echoview)
    - [Go.Rice](https://github.com/foolin/goview/tree/master/supports/gorice)
- [Usage](#usage)
    - [Overview](#overview)
    - [Config](#config)
    - [Include syntax](#include-syntax)
    - [Render name](#render-name)
- [Examples](#examples)
    - [Basic example](#basic-example)
    - [Gin example](#gin-example)
    - [Echo example](#echo-example)
    - [Go-chi example](#go-chi-example)
    - [Advance example](#advance-example)
    - [Multiple example](#multiple-example)
    - [go.rice example](#gorice-example)
    - [more examples](#more-examples)


## Install
```bash
go get github.com/foolin/goview
```


## Features

* **Lightweight** - use golang html/template syntax.
* **Easy** - easy use for your web application.
* **Fast** - Support configure cache template.
* **Include syntax** - Support include file.
* **Master layout** - Support configure master layout file.
* **Extension** - Support configure template file extension.
* **Easy** - Support configure templates directory.
* **Auto reload** - Support dynamic reload template(disable cache mode).
* **Multiple Engine** - Support multiple templates for frontend and backend.
* **No external dependencies** - plain ol' Go html/template.
* **Gorice** - Support gorice for package resources.
* **Gin/Echo/Chi** - Support gin framework,echo framework, go-chi framework.


## Docs
See <https://www.godoc.org/github.com/foolin/goview>


## Supports
- **[ginview](https://github.com/foolin/goview/tree/master/supports/ginview)** goview for gin framework
- **[echoview](https://github.com/foolin/goview/tree/master/supports/echoview)** goview for echo framework
- **[gorice](https://github.com/foolin/goview/tree/master/supports/gorice)** goview for go.rice


## Usage

### Overview

Project structure:

```go
|-- app/views/
    |--- index.html          
    |--- page.html
    |-- layouts/
        |--- footer.html
        |--- master.html
   
```

Use default instance:

```go
    //write http.ResponseWriter
    //"index" -> index.html
    goview.Render(writer, http.StatusOK, "index", goview.M{})
```

Use new instance with config:

```go

    gv := goview.New(goview.Config{
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
    goview.Use(gv)
    
    //write http.ResponseWriter
    goview.Render(writer, http.StatusOK, "index", goview.M{})

```


Use multiple instance with config:

```go
    //============== Frontend ============== //
    gvFrontend := goview.New(goview.Config{
        Root:      "views/frontend",
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
    
    //write http.ResponseWriter
    gvFrontend.Render(writer, http.StatusOK, "index", goview.M{})
    
    //============== Backend ============== //
    gvBackend := goview.New(goview.Config{
        Root:      "views/backend",
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
    
    //write http.ResponseWriter
    gvBackend.Render(writer, http.StatusOK, "index", goview.M{})

```

### Config

```go
goview.Config{
    Root:      "views", //template root path
    Extension: ".tpl", //file extension
    Master:    "layouts/master", //master layout file
    Partials:  []string{"partials/head"}, //partial files
    Funcs: template.FuncMap{
        "sub": func(a, b int) int {
            return a - b
        },
        // more funcs
    },
    DisableCache: false, //if disable cache, auto reload template file for debug.
}
```

### Include syntax

```go
//template file
{{include "layouts/footer"}}
```

### Render name: 

Render name use `index` without `.html` extension, that will render with master layout.

- **"index"** - Render with master layout.
- **"index.html"** - Not render with master layout.

```
Notice: `.html` is default template extension, you can change with config
```


Render with master

```go
//use name without extension `.html`
goview.Render(w, http.StatusOK, "index", goview.M{})
```

The `w` is instance of  `http.ResponseWriter`

Render only file(not use master layout)
```go
//use full name with extension `.html`
goview.Render(w, http.StatusOK, "page.html", goview.M{})
```





## Examples

See [_examples/](https://github.com/foolin/goview/blob/master/_examples/) for a variety of examples.


### Basic example
```go

package main

import (
	"fmt"
	"github.com/foolin/goview"
	"net/http"
)

func main() {

	//render index use `index` without `.html` extension, that will render with master layout.
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

	//render page use `page.tpl` with '.html' will only file template without master layout.
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
    

See in "examples/basic" folder
```

[Basic example](https://github.com/foolin/goview/tree/master/_examples/basic)


### Gin example

```bash
go get github.com/foolin/goview/supports/ginview
```

```go

package main

import (
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	//new template engine
	router.HTMLRender = ginview.Default()

	router.GET("/", func(ctx *gin.Context) {
		//render with master
		ctx.HTML(http.StatusOK, "index", gin.H{
			"title": "Index title!",
			"add": func(a int, b int) int {
				return a + b
			},
		})
	})

	router.GET("/page", func(ctx *gin.Context) {
		//render only file, must full name with extension
		ctx.HTML(http.StatusOK, "page.html", gin.H{"title": "Page file title!!"})
	})

	router.Run(":9090")
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
    

See in "examples/basic" folder
```

[Gin example](https://github.com/foolin/goview/tree/master/_examples/gin)



### Echo example

Echo <=v3 version:
```bash
go get github.com/foolin/goview/supports/echoview
```

Echo v4 version:

```bash
go get github.com/foolin/goview/supports/echoview-v4
```


```go

package main

import (
	"github.com/foolin/goview/supports/echoview"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main() {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//Set Renderer
	e.Renderer = echoview.Default()

	// Routes
	e.GET("/", func(c echo.Context) error {
		//render with master
		return c.Render(http.StatusOK, "index", echo.Map{
			"title": "Index title!",
			"add": func(a int, b int) int {
				return a + b
			},
		})
	})

	e.GET("/page", func(c echo.Context) error {
		//render only file, must full name with extension
		return c.Render(http.StatusOK, "page.html", echo.Map{"title": "Page file title!!"})
	})

	// Start server
	e.Logger.Fatal(e.Start(":9090"))
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
    

See in "examples/basic" folder
```

[Echo example](https://github.com/foolin/goview/tree/master/_examples/echo)
[Echo v4 example](https://github.com/foolin/goview/tree/master/_examples/echo-v4)


### Go-chi example
```go

package main

import (
	"fmt"
	"github.com/foolin/goview"
	"github.com/go-chi/chi"
	"net/http"
)

func main() {

	r := chi.NewRouter()

	//render index use `index` without `.html` extension, that will render with master layout.
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
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

	//render page use `page.tpl` with '.html' will only file template without master layout.
	r.Get("/page", func(w http.ResponseWriter, r *http.Request) {
		err := goview.Render(w, http.StatusOK, "page.html", goview.M{"title": "Page file title!!"})
		if err != nil {
			fmt.Fprintf(w, "Render page.html error: %v!", err)
		}
	})

	fmt.Println("Listening and serving HTTP on :9090")
	http.ListenAndServe(":9090", r)

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
    

See in "examples/basic" folder
```

[Chi example](https://github.com/foolin/goview/tree/master/_examples/go-chi)



### Advance example
```go

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
	goview.Use(gv)

	//render index use `index` without `.html` extension, that will render with master layout.
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

	//render page use `page.tpl` with '.html' will only file template without master layout.
	http.HandleFunc("/page", func(w http.ResponseWriter, r *http.Request) {
		err := goview.Render(w, http.StatusOK, "page.tpl", goview.M{"title": "Page file title!!"})
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
    |--- index.tpl          
    |--- page.tpl
    |-- layouts/
        |--- footer.tpl
        |--- head.tpl
        |--- master.tpl
    |-- partials/
        |--- ad.tpl
    

See in "examples/advance" folder
```

[Advance example](https://github.com/foolin/goview/tree/master/_examples/advance)

### Multiple example
```go

package main

import (
	"html/template"
	"net/http"
	"time"

	"github.com/foolin/goview"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//new template engine
	router.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:      "views/fontend",
		Extension: ".html",
		Master:    "layouts/master",
		Partials:  []string{"partials/ad"},
		Funcs: template.FuncMap{
			"copy": func() string {
				return time.Now().Format("2006")
			},
		},
		DisableCache: true,
	})

	router.GET("/", func(ctx *gin.Context) {
		// `HTML()` is a helper func to deal with multiple TemplateEngine's.
		// It detects the suitable TemplateEngine for each path automatically.
		gintemplate.HTML(ctx, http.StatusOK, "index", gin.H{
			"title": "Fontend title!",
		})
	})

	//=========== Backend ===========//

	//new middleware
	mw := gintemplate.NewMiddleware(gintemplate.TemplateConfig{
		Root:      "views/backend",
		Extension: ".html",
		Master:    "layouts/master",
		Partials:  []string{},
		Funcs: template.FuncMap{
			"copy": func() string {
				return time.Now().Format("2006")
			},
		},
		DisableCache: true,
	})

	// You should use helper func `Middleware()` to set the supplied
	// TemplateEngine and make `HTML()` work validly.
	backendGroup := router.Group("/admin", mw)

	backendGroup.GET("/", func(ctx *gin.Context) {
		// With the middleware, `HTML()` can detect the valid TemplateEngine.
		gintemplate.HTML(ctx, http.StatusOK, "index", gin.H{
			"title": "Backend title!",
		})
	})

	router.Run(":9090")
}


```

Project structure:
```go
|-- app/views/
    |-- fontend/
        |--- index.html
        |-- layouts/
            |--- footer.html
            |--- head.html
            |--- master.html
        |-- partials/
     	   |--- ad.html
    |-- backend/
        |--- index.html
        |-- layouts/
            |--- footer.html
            |--- head.html
            |--- master.html
        
See in "examples/multiple" folder
```

[Multiple example](https://github.com/foolin/goview/tree/master/_examples/multiple)


### go.rice example

```bash
go get github.com/foolin/goview/supports/gorice
```

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

	//render index use `index` without `.html` extension, that will render with master layout.
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

	//render page use `page.tpl` with '.html' will only file template without master layout.
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

### More examples

See [_examples/](https://github.com/foolin/goview/blob/master/_examples/) for a variety of examples.


[GoDoc]: https://godoc.org/github.com/foolin/goview
[GoDoc Widget]: https://godoc.org/github.com/foolin/goview?status.svg
[Travis]: https://travis-ci.org/foolin/goview
[Travis Widget]: https://travis-ci.org/foolin/goview.svg?branch=master
[GoReportCard]: https://goreportcard.com/report/github.com/foolin/goview
[GoReportCard Widget]: https://goreportcard.com/badge/github.com/foolin/goview
[GoCover]: https://goreportcard.com/report/github.com/foolin/goview
[GoCover Widget]: https://goreportcard.com/badge/github.com/foolin/goview


### Todo
 [ ] Add Partials support directory or glob
 [ ] Add functions support.
 