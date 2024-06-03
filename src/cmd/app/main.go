package main

import (
	"bytes"
	"database/sql"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"github.com/runebone/db-cw-s6-bmstu/internal/domain/handlers"
	"github.com/runebone/db-cw-s6-bmstu/internal/domain/repositories"
	"github.com/runebone/db-cw-s6-bmstu/internal/domain/services"
)

// var store = sessions.NewCookieStore([]byte("secret"))

type Template struct {
	tmpl *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	var buf bytes.Buffer
	if err := t.tmpl.ExecuteTemplate(&buf, name, data); err != nil {
		return err
	}
	_, err := buf.WriteTo(w)
	return err
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	db, err := sql.Open("postgres", "port=8001 user=admin password=admin dbname=db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	t := &Template{
		tmpl: template.Must(template.ParseGlob("web/templates/*.html")),
	}
	e.Renderer = t

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// documentRepo := repositories.NewDocumentRepository(db)
	// documentService := services.NewDocumentService(documentRepo)
	// documentHandler := handlers.NewDocumentHandler(documentService)

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", nil)
	})

	e.GET("/hello", func(c echo.Context) error {
		return c.Render(http.StatusOK, "hello.html", nil)
	})

	e.GET("/register", userHandler.ShowRegisterForm)
	e.POST("/register", userHandler.RegisterUser)

	e.GET("/login", userHandler.ShowLoginForm)
	e.POST("/login", userHandler.LoginUser)

	e.GET("/profile", userHandler.ShowProfile)

	// e.GET("/documents/:id", documentHandler.GetDocumentText)

	e.Logger.Fatal(e.Start(":8080"))
}
