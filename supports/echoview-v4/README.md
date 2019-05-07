# EchoView

[![GoDoc Widget]][GoDoc] 

goview support for echo template v4 version.

## Install
```bash

go get -u github.com/foolin/goview-v4

go get -u github.com/foolin/goview/supports/echoview-v4

```

### Example

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

[Echo example](https://github.com/foolin/goview/tree/master/_examples/echo-v4)


           
## More examples

See [_examples/](https://github.com/foolin/goview/blob/master/_examples/) for a variety of examples.

[GoDoc]: https://godoc.org/github.com/foolin/goview/supports/echoview-v4
[GoDoc Widget]: https://godoc.org/github.com/foolin/goview/supports/echoview-v4?status.svg
