package view

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
)

type Renderer struct {
	Templates *template.Template
}

func (t *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}
