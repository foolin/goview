package goview

import (
	"bytes"
	"html/template"
	"testing"
)

var cases = []struct {
	Name string
	Data M
	Out  string
}{
	{
		Name: "echo.tpl",
		Data: M{"name": "GoView"},
		Out:  "$GoView",
	},
	{
		Name: "include",
		Data: M{"name": "GoView"},
		Out:  "<v>IncGoView</v>",
	},
	{
		Name: "index",
		Data: M{},
		Out:  "<v>Index</v>",
	},
	{
		Name: "sum",
		Data: M{
			"sum": func(a int, b int) int {
				return a + b
			},
			"a": 1,
			"b": 2,
		},
		Out: "<v>3</v>",
	},
}

func TestViewEngine_RenderWriter(t *testing.T) {
	gv := New(Config{
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

	for _, v := range cases {
		buff := new(bytes.Buffer)
		err := gv.RenderWriter(buff, v.Name, v.Data)
		if err != nil {
			t.Errorf("name: %v, data: %v, error: %v", v.Name, v.Data, err)
			continue
		}
		val := string(buff.Bytes())
		if val != v.Out {
			t.Errorf("actual: %v, expect: %v", val, v.Out)
		}
	}
}
