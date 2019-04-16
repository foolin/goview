/*
Goview is a lightweight, minimalist and idiomatic template library based on golang html/template for building Go web application.

Example:

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

		//render page use `page.html` with '.html' will only file template without master layout.
		http.HandleFunc("/page", func(w http.ResponseWriter, r *http.Request) {
			err := goview.Render(w, http.StatusOK, "page.html", goview.M{"title": "Page file title!!"})
			if err != nil {
				fmt.Fprintf(w, "Render page.html error: %v!", err)
			}
		})

		fmt.Println("Listening and serving HTTP on :9090")
		http.ListenAndServe(":9090", nil)

	}


Project structure:

   |-- app/views/
	   |--- index.html
	   |--- page.html
	   |-- layouts/
		   |--- footer.html
		   |--- master.html

Learn more at https://github.com/foolin/goview

================== Supports ==================

Gin framework:
https://godoc.org/github.com/foolin/goview/supports/ginview

Gin framework:
https://godoc.org/github.com/foolin/goview/supports/ginview

Gin framework:
https://godoc.org/github.com/foolin/goview/supports/ginview

Examples:
https://godoc.org/github.com/foolin/goview/_examples

*/

package goview

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

//const templateEngineKey = "httpx_templateEngine"
var HtmlContentType = []string{"text/html; charset=utf-8"}

var DefaultConfig = Config{
	Root:         "views",
	Extension:    ".html",
	Master:       "layouts/master",
	Partials:     []string{},
	Funcs:        make(template.FuncMap),
	DisableCache: false,
	Delims:       Delims{Left: "{{", Right: "}}"},
}

type ViewEngine struct {
	config      Config
	tplMap      map[string]*template.Template
	tplMutex    sync.RWMutex
	fileHandler FileHandler
}

type Config struct {
	Root         string           //view root
	Extension    string           //template extension
	Master       string           //template master
	Partials     []string         //template partial, such as head, foot
	Funcs        template.FuncMap //template functions
	DisableCache bool             //disable cache, debug mode
	Delims       Delims           //delimeters
}

type M map[string]interface{}

type Delims struct {
	Left  string
	Right string
}

type FileHandler func(config Config, tplFile string) (content string, err error)

func New(config Config) *ViewEngine {
	return &ViewEngine{
		config:      config,
		tplMap:      make(map[string]*template.Template),
		tplMutex:    sync.RWMutex{},
		fileHandler: DefaultFileHandler(),
	}
}

func Default() *ViewEngine {
	return New(DefaultConfig)
}

func (e *ViewEngine) Render(w http.ResponseWriter, statusCode int, name string, data interface{}) error {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = HtmlContentType
	}
	w.WriteHeader(statusCode)
	return e.executeRender(w, name, data)
}

func (e *ViewEngine) RenderWriter(w io.Writer, name string, data interface{}) error {
	return e.executeRender(w, name, data)
}

func (e *ViewEngine) executeRender(out io.Writer, name string, data interface{}) error {
	useMaster := true
	if filepath.Ext(name) == e.config.Extension {
		useMaster = false
		name = strings.TrimSuffix(name, e.config.Extension)

	}
	return e.executeTemplate(out, name, data, useMaster)
}

func (e *ViewEngine) executeTemplate(out io.Writer, name string, data interface{}, useMaster bool) error {
	var tpl *template.Template
	var err error
	var ok bool

	allFuncs := make(template.FuncMap, 0)
	allFuncs["include"] = func(layout string) (template.HTML, error) {
		buf := new(bytes.Buffer)
		err := e.executeTemplate(buf, layout, data, false)
		return template.HTML(buf.String()), err
	}

	// Get the plugin collection
	for k, v := range e.config.Funcs {
		allFuncs[k] = v
	}

	e.tplMutex.RLock()
	tpl, ok = e.tplMap[name]
	e.tplMutex.RUnlock()

	exeName := name
	if useMaster && e.config.Master != "" {
		exeName = e.config.Master
	}

	if !ok || e.config.DisableCache {
		tplList := make([]string, 0)
		if useMaster {
			//render()
			if e.config.Master != "" {
				tplList = append(tplList, e.config.Master)
			}
		}
		tplList = append(tplList, name)
		tplList = append(tplList, e.config.Partials...)

		// Loop through each template and test the full path
		tpl = template.New(name).Funcs(allFuncs).Delims(e.config.Delims.Left, e.config.Delims.Right)
		for _, v := range tplList {
			var data string
			data, err = e.fileHandler(e.config, v)
			if err != nil {
				return err
			}
			var tmpl *template.Template
			if v == name {
				tmpl = tpl
			} else {
				tmpl = tpl.New(v)
			}
			_, err = tmpl.Parse(data)
			if err != nil {
				return fmt.Errorf("ViewEngine render parser name:%v, error: %v", v, err)
			}
		}
		e.tplMutex.Lock()
		e.tplMap[name] = tpl
		e.tplMutex.Unlock()
	}

	// Display the content to the screen
	err = tpl.Funcs(allFuncs).ExecuteTemplate(out, exeName, data)
	if err != nil {
		return fmt.Errorf("ViewEngine execute template error: %v", err)
	}

	return nil
}

func (e *ViewEngine) SetFileHandler(handle FileHandler) {
	if handle == nil {
		panic("FileHandler can't set nil!")
	}
	e.fileHandler = handle
}

func DefaultFileHandler() FileHandler {
	return func(config Config, tplFile string) (content string, err error) {
		// Get the absolute path of the root template
		path, err := filepath.Abs(config.Root + string(os.PathSeparator) + tplFile + config.Extension)
		if err != nil {
			return "", fmt.Errorf("ViewEngine path:%v error: %v", path, err)
		}
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return "", fmt.Errorf("ViewEngine render read name:%v, path:%v, error: %v", tplFile, path, err)
		}
		return string(data), nil
	}
}
