package echoview

import (
	"github.com/foolin/goview"
	"github.com/labstack/echo/v4"
	"io"
)

const templateEngineKey = "foolin-goview-echoview"

type ViewEngine struct {
	*goview.ViewEngine
}

func New(config goview.Config) *ViewEngine {
	return &ViewEngine{
		ViewEngine: goview.New(config),
	}
}

func Default() *ViewEngine {
	return New(goview.DefaultConfig)
}

func (e *ViewEngine) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return e.RenderWriter(w, name, data)
}

// You should use helper func `Middleware()` to set the supplied
// TemplateEngine and make `Render()` work validly.
func Render(ctx echo.Context, code int, name string, data interface{}) error {
	if val := ctx.Get(templateEngineKey); val != nil {
		if e, ok := val.(*ViewEngine); ok {
			return e.Render(ctx.Response().Writer, name, data, ctx)
		}
	}
	return ctx.Render(code, name, data)
}

//New gin middleware for func `goview.Render()`
func NewMiddleware(config goview.Config) echo.MiddlewareFunc {
	return Middleware(New(config))
}

func Middleware(e *ViewEngine) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(templateEngineKey, e)
			return next(c)
		}
	}
}
