package main

import (
	"bytes"
	"database/sql"
	"html/template"
	"io"
	"log"
	"net/http"

	keydb "github.com/go-redis/redis/v8"
	"github.com/gorilla/sessions"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	h "github.com/runebone/db-cw-s6-bmstu/internal/domain/handlers"
	r "github.com/runebone/db-cw-s6-bmstu/internal/domain/repositories"
	s "github.com/runebone/db-cw-s6-bmstu/internal/domain/services"
)

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

	cache := keydb.NewClient(&keydb.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	store := sessions.NewCookieStore([]byte("secret"))

	userRepo := r.NewUserRepository(db)
	userService := s.NewUserService(userRepo)
	userHandler := h.NewUserHandler(userService, store)

	documentRepo := r.NewDocumentRepository(db)
	documentService := s.NewDocumentService(documentRepo)
	documentHandler := h.NewDocumentHandler(documentService, cache, store)

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

	e.GET("/d/:id", documentHandler.GetDocumentText)
	e.GET("/d", func(c echo.Context) error {
		return c.Render(http.StatusOK, "upload-document.html", nil)
	})
	e.POST("/d", documentHandler.UploadDocument)

	e.GET("/search", func(c echo.Context) error {
		return c.Render(http.StatusOK, "search.html", nil)
	})
	e.POST("/search", documentHandler.GetDocumentsByContent)

	e.Logger.Fatal(e.Start(":8080"))
}
