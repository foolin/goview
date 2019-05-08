package goview

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDefault(t *testing.T) {
	recorder := httptest.NewRecorder()
	err := Render(recorder, http.StatusOK, "index", M{})
	//expect like this info:
	// "error: ViewEngine render read name:layouts/master,  path:/foolin/goview/views/layouts/master.html, error: open /foolin/goview/views/layouts/master.html: The system cannot find the path specified"
	if err == nil {
		t.Error("render is ok?")
	} else {
		t.Logf("expect error: %v", err)
	}
}

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
	assertRecorder(t, recorder, http.StatusOK, expect)
}

func assertRecorder(t *testing.T, recorder *httptest.ResponseRecorder, expectStatusCode int, expectOut string) {
	result := recorder.Result()
	if result.StatusCode != expectStatusCode {
		t.Errorf("actual: %v, expect: %v", result.StatusCode, expectStatusCode)
	}
	resultBytes, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Errorf("read result body error: %v", err)
		return
	}
	val := string(resultBytes)
	if val != expectOut {
		t.Errorf("actual: %v, expect: %v", val, expectOut)
	}
}
