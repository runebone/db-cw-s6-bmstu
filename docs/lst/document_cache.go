type DocumentHandler struct {
	service *s.DocumentService
	cache   *keydb.Client
	store   *sessions.CookieStore
}

// ...

func (h *DocumentHandler) GetDocumentText(c echo.Context) error {
	documentID := c.Param("id")
	useCache := true

	if useCache {
		cachedText, err := h.cache.Get(ctx, documentID).Result()
		if err == keydb.Nil {
			uid, err := uuid.Parse(documentID)
			if err != nil {
				return c.String(http.StatusBadRequest, "Invalid document ID - parse uuid")
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
	} else {
		uid, err := uuid.Parse(documentID)
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid document ID - parse uuid")
		}

		text, err := h.service.GetDocumentText(uid)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.String(http.StatusOK, text)
	}
}
