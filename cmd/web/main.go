package main

import (
	"io"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/troptropcontent/what_the_tide/database"
	"github.com/troptropcontent/what_the_tide/internal/models"
	calendar_handlers "github.com/troptropcontent/what_the_tide/internal/modules/calendar/handlers"
)

func healthCheckHandler(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func createSubscriptionHandler(c echo.Context) error {
	return c.File("internal/views/subscription/created.html")
}

func rootHandler(c echo.Context) error {
	ports := []models.Port{}
	database.DB.Order("name asc").Find(&ports)
	return c.Render(http.StatusOK, "root", ports)
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	// Initialize the database
	database.MustInit()

	t := &Template{
		templates: template.Must(template.ParseGlob("internal/views/*.html")),
	}
	e := echo.New()
	e.Renderer = t
	e.Use(middleware.Logger())
	e.GET("/", rootHandler)
	e.GET("/up", healthCheckHandler)
	e.Static("/public", "public")

	agendaRoutes := e.Group("/agenda")
	agendaRoutes.POST("/subscription", calendar_handlers.CreateSubscription)

	e.Logger.Fatal(e.Start(":3001"))
}
