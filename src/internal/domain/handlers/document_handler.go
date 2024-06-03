package handlers

import (
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	s "github.com/runebone/db-cw-s6-bmstu/internal/domain/services"
)

type DocumentHandler struct {
	DocumentService *s.DocumentService
}

func NewDocumentHandler(service *s.DocumentService) *DocumentHandler {
	return &DocumentHandler{DocumentService: service}
}

func (h *DocumentHandler) GetDocumentText(c echo.Context) error {
	documentID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid document ID")
	}

	text, err := h.DocumentService.GetDocumentText(documentID)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, text)
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
