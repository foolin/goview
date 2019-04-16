package gorice

import (
	"github.com/GeertJohan/go.rice"
	"github.com/foolin/goview"
)

/**
New gin template engine, default views root.
*/
func New(viewsRootBox *rice.Box) *goview.ViewEngine {
	return NewWithConfig(viewsRootBox, goview.DefaultConfig)
}

/**
New gin template engine
Important!!! The viewsRootBox's name and config.Root must be consistent.
*/
func NewWithConfig(viewsRootBox *rice.Box, config goview.Config) *goview.ViewEngine {
	config.Root = viewsRootBox.Name()
	engine := goview.New(config)
	engine.SetFileHandler(FileHandler(viewsRootBox))
	return engine
}

/**
Support go.rice file handler
*/
func FileHandler(viewsRootBox *rice.Box) goview.FileHandler {
	return func(config goview.Config, tplFile string) (content string, err error) {
		// get file contents as string
		return viewsRootBox.String(tplFile + config.Extension)
	}
}