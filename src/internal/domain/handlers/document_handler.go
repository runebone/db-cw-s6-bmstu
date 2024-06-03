package handlers

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	s "github.com/runebone/db-cw-s6-bmstu/internal/domain/services"
)

var ctx = context.Background()

type DocumentHandler struct {
	DocumentService *s.DocumentService
	Cache           *redis.Client
}

func NewDocumentHandler(service *s.DocumentService, cache *redis.Client) *DocumentHandler {
	return &DocumentHandler{
		DocumentService: service,
		Cache:           cache,
	}
}

func (h *DocumentHandler) GetDocumentText(c echo.Context) error {
	documentID := c.Param("id")

	cachedText, err := h.Cache.Get(ctx, documentID).Result()
	if err == redis.Nil {
		uid, err := uuid.Parse(documentID)
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid document ID")
		}

		text, err := h.DocumentService.GetDocumentText(uid)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		err = h.Cache.Set(ctx, documentID, text, 10*time.Minute).Err()
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

	uuids, err := h.DocumentService.DocumentRepo.GetDocumentsByContent(text)
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
