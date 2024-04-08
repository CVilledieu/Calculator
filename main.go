package main

import (
	"fmt"
	"io"
	"strconv"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("site/*.html")),
	}
}

type Value struct {
	Name string
}

type Page struct {
	Value Value
}

func newValue() Value {
	return Value{
		Name: "0",
	}
}

func newPage() Page {
	return Page{
		Value: newValue(),
	}
}

func doMath(firstNum, secondNum, operator string) string {
	answer := ""
	switch operator {
	case "+":
		f, _ := strconv.ParseFloat(firstNum, 64)
		s, _ := strconv.ParseFloat(secondNum, 64)
		answerFloat := f + s
		answer = strconv.FormatFloat(answerFloat, 'f', -1, 64)
	}

	return answer
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	page := newPage()
	e.Renderer = newTemplate()

	storedValue := ""
	operator := ""

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", page)
	})

	e.POST("/update", func(c echo.Context) error {
		value := c.FormValue("numpad")
		fmt.Println(value)
		newValue := page.Value.Name + value
		page.Value.Name = newValue
		return c.Render(200, "display", page)
	})

	e.POST("/next", func(c echo.Context) error {
		storedValue = page.Value.Name
		page.Value.Name = "0"
		operator = c.FormValue("op")
		return c.Render(200, "display", page)
	})

	e.GET("/return", func(c echo.Context) error {
		page.Value.Name = doMath(storedValue, page.Value.Name, operator)
		return c.Render(200, "display", page)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
