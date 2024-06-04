package handlers

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	m "github.com/runebone/db-cw-s6-bmstu/internal/domain/models"
	s "github.com/runebone/db-cw-s6-bmstu/internal/domain/services"
)

type UserHandler struct {
	service *s.UserService
	store   *sessions.CookieStore
}

func NewUserHandler(service *s.UserService, store *sessions.CookieStore) *UserHandler {
	return &UserHandler{
		service: service,
		store:   store,
	}
}

func (h *UserHandler) ShowRegisterForm(c echo.Context) error {
	return c.Render(http.StatusOK, "register.html", nil)
}

func (h *UserHandler) RegisterUser(c echo.Context) error {
	username := m.Username(c.FormValue("username"))
	email := m.Email(c.FormValue("email"))
	password := m.Password(c.FormValue("password"))

	err := h.service.RegisterUser(username, email, password)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if c.Request().Header.Get("HX-Request") != "" {
		return c.HTML(http.StatusOK, "<p>Registration successful! Please <a href='/login' hx-get='/login' hx-target='#main'>login</a>.</p>")
	}

	return c.Redirect(http.StatusSeeOther, "/login")
}

func (h *UserHandler) ShowLoginForm(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", nil)
}

func (h *UserHandler) LoginUser(c echo.Context) error {
	username := m.Username(c.FormValue("username"))
	password := m.Password(c.FormValue("password"))

	user, err := h.service.AuthenticateUser(username, password)
	if err != nil {
		return c.String(http.StatusUnauthorized, "Invalid credentials")
	}

	sess, _ := h.store.Get(c.Request(), "session")
	sess.Values["user_id"] = user.ID.String()
	sess.Save(c.Request(), c.Response())

	if c.Request().Header.Get("HX-Request") != "" {
		return c.HTML(http.StatusOK, "<p>Login successful! Go to your <a href='/profile' hx-get='/profile' hx-target='#main'>profile</a>.</p>")
	}

	return c.Redirect(http.StatusSeeOther, "/profile")
}

func (h *UserHandler) ShowProfile(c echo.Context) error {
	sess, _ := h.store.Get(c.Request(), "session")
	userID, ok := sess.Values["user_id"].(string)
	fmt.Printf("\nuserID: %s\n\n", userID)
	if !ok {
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	uid, err := uuid.Parse(userID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Invalid user ID")
	}

	user, err := h.service.UserRepo.GetUserByID(uid)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.Render(http.StatusOK, "profile.html", map[string]interface{}{
		"user": user,
	})
}
