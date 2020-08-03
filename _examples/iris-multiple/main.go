/*
 * Copyright 2018 Foolin.  All rights reserved.
 *
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

import (
	"html/template"
	"time"

	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/irisview"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	// Register a new template engine.
	app.RegisterView(irisview.New(goview.Config{
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
	}))

	app.Get("/", func(ctx iris.Context) {
		ctx.View("index", iris.Map{
			"title": "Frontend title!",
		})
	})

	//=========== Backend ===========//

	// Assign a new template middleware.
	mw := irisview.NewMiddleware(goview.Config{
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

	backendGroup := app.Party("/admin", mw)

	backendGroup.Get("/", func(ctx iris.Context) {
		// Use the ctx.View as you used to. Zero changes to your codebase,
		// even if you use multiple templates.
		ctx.View("index", iris.Map{
			"title": "Backend title!",
		})
	})

	app.Listen(":9090")
}
