package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestDeleteTranscoder(t *testing.T) {
	t.Run("Transcoder should be deleted successfully", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodDelete, "/transcoders?output_type=dash&input_type=mp4", nil)
		rec := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.DeleteTranscoder(c)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NoError(t, err)
	})

	t.Run("Transcoder deletion should fail - Invalid query param", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/transcoders", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.DeleteTranscoder(c)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.NoError(t, err)
	})
	t.Run("Transcoder deletion should fail - Transcoder not found", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/transcoders?output_type=dash&input_type=mp4", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.DeleteTranscoder(c)
		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.NoError(t, err)
	})
}