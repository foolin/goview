package goview

import (
	"net/http"
)

var instance *ViewEngine

// Use setting default instance engine
func Use(engine *ViewEngine) {
	instance = engine
}

// Render render view template with default instance
func Render(w http.ResponseWriter, status int, name string, data interface{}) error {
	if instance == nil {
		instance = Default()
		//return fmt.Errorf("instance not yet initialized, please call Init() first before Render()")
	}
	return instance.Render(w, status, name, data)
}
