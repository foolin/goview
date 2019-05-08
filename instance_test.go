package goview

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRender(t *testing.T) {

	engine := New(Config{
		Root:      "_examples/test",
		Extension: ".tpl",
		Master:    "layouts/master",
		Partials:  []string{},
		Funcs: template.FuncMap{
			"echo": func(v string) string {
				return "$" + v
			},
		},
		DisableCache: true,
	})

	Use(engine)

	recorder := httptest.NewRecorder()
	expect := "<v>Index</v>"
	err := Render(recorder, http.StatusOK, "index", M{})
	if err != nil {
		t.Errorf("render error: %v", err)
		return
	}
	result := recorder.Result()
	if result.StatusCode != http.StatusOK {
		t.Errorf("actual: %v, expect: %v", result.StatusCode, http.StatusOK)
	}
	resultBytes, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Errorf("read result body error: %v", err)
		return
	}
	val := string(resultBytes)
	if val != expect {
		t.Errorf("actual: %v, expect: %v", val, expect)
	}
}
