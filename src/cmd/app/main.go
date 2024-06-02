package main

import (
	"database/sql"
	"html/template"
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	db, err := sql.Open("postgres", "port=8001 user=admin password=admin dbname=db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	t := &Template{
		templates: template.Must(template.ParseGlob("web/templates/*.html")),
	}
	e.Renderer = t

	e.GET("/register", handlers.ShowRegisterForm)
	e.POST("/register", handlers.RegisterUser(db))

	e.GET("/login", handlers.ShowLoginForm)
	e.POST("/login", handlers.LoginUser(db))

	e.GET("/profile", handlers.ShowProfile)

	e.Logger.Fatal(e.Start(":8080"))
}
