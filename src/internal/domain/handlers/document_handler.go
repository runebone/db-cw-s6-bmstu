package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	m "github.com/runebone/db-cw-s6-bmstu/internal/domain/models"
	s "github.com/runebone/db-cw-s6-bmstu/internal/domain/services"
)

var ctx = context.Background()

type DocumentHandler struct {
	service *s.DocumentService
	cache   *redis.Client
	store   *sessions.CookieStore
}

func NewDocumentHandler(service *s.DocumentService, cache *redis.Client, store *sessions.CookieStore) *DocumentHandler {
	return &DocumentHandler{
		service: service,
		cache:   cache,
		store:   store,
	}
}

func (h *DocumentHandler) GetDocumentText(c echo.Context) error {
	documentID := c.Param("id")

	cachedText, err := h.cache.Get(ctx, documentID).Result()
	if err == redis.Nil {
		uid, err := uuid.Parse(documentID)
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid document ID")
		}

		text, err := h.service.GetDocumentText(uid)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		err = h.cache.Set(ctx, documentID, text, 10*time.Minute).Err()
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to cache document text")
		}

		return c.String(http.StatusOK, text)
	} else if err != nil {
		return c.String(http.StatusBadRequest, "Invalid document ID")
	}

	return c.String(http.StatusOK, cachedText)
}

func (h *DocumentHandler) GetDocumentsByContent(c echo.Context) error {
	text := c.FormValue("content")

	uuids, err := h.service.GetDocumentsByContent(text)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// XXX meeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeh
	strs := make([]string, len(uuids))
	for i, u := range uuids {
		strs[i] = `<a href="/d/` + u.String() + `" hx-get="/d/` + u.String() + `" hx-target="#contents">` + u.String() + "</a>"
	}

	return c.String(http.StatusOK, strings.Join(strs, "<br>"))
}

func (h *DocumentHandler) UploadDocument(c echo.Context) error {
	title := c.FormValue("title")
	lang := c.FormValue("lang")
	orig := c.FormValue("orig")

	sess, _ := h.store.Get(c.Request(), "session")
	userID, ok := sess.Values["user_id"].(string)
	fmt.Printf("\nuserID (from doc): %s\n\n", userID)
	if !ok {
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	uid, err := uuid.Parse(userID)
	if err != nil {
		return err
	}

	origuid, err := uuid.Parse(orig)
	if err != nil {
		return err
	}

	doc := m.NewDocument(
		"uploaded.docs",
		title,
		m.Lang(lang),
		origuid,
		uid,
	)

	err = h.service.CreateDocument(doc)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if c.Request().Header.Get("HX-Request") != "" {
		return c.HTML(http.StatusOK, "<p>Huge success!</p>")
	}

	return c.Redirect(http.StatusSeeOther, "/login")
}
