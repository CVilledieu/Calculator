package main

import (
	"fmt"
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("site/*html")),
	}
}

type Calculator struct {
	CurrentValue input
	SecondValue  input
	Operator     string
}

type input struct {
	number float64
}

func main() {
	fmt.Printf("hello world")
}
