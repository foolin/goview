/*
 * Copyright 2018 Foolin.  All rights reserved.
 *
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

import (
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/echoview"
	"html/template"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//new template engine
	e.Renderer = echoview.New(goview.Config{
		Root:      "views/frontend",
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

	e.GET("/", func(ctx echo.Context) error {
		// `HTML()` is a helper func to deal with multiple TemplateEngine's.
		// It detects the suitable TemplateEngine for each path automatically.
		return echoview.Render(ctx, http.StatusOK, "index", echo.Map{
			"title": "Frontend title!",
		})
	})

	//=========== Backend ===========//

	//new middleware
	mw := echoview.NewMiddleware(goview.Config{
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
	backendGroup := e.Group("/admin", mw)

	backendGroup.GET("/", func(ctx echo.Context) error {
		// With the middleware, `HTML()` can detect the valid TemplateEngine.
		return echoview.Render(ctx, http.StatusOK, "index", echo.Map{
			"title": "Backend title!",
		})
	})

	// Start server
	e.Logger.Fatal(e.Start(":9090"))
}
