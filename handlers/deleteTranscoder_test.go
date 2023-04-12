package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// TOD0: Deleting now will chage the status to inactive, but will not delete the transcoder from the database
func TestDeleteTranscoder(t *testing.T) {
	t.Run("Transcoder should be deleted successfully", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodDelete, "/transcoders?output_type=dash&input_type=mp4&codec=h264&descriptor=media_analysis", nil)
		rec := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", jwtToken)
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
		req.Header.Set("Authorization", jwtToken)
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
		req := httptest.NewRequest(http.MethodDelete, "/transcoders?output_type=dash&input_type=mp4&codec=h264&descriptor=media_analysis", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", jwtToken)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.DeleteTranscoder(c)
		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.NoError(t, err)
	})
}
