package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var transcoder Transcoder

func TestPutTranscoder(t *testing.T) {
	t.Run("Transcoder putting should fail - No input_type  present", func(t *testing.T) {
		e := echo.New()
		body, err := GetDummyData(map[string]any{}, map[string]string{"input_type": "inp"})
		assert.NoError(t, err)
		fmt.Println(body)
		req := httptest.NewRequest(http.MethodPut, "/transcoders", strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", jwtToken)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err = h.PutTranscoder(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
	t.Run("Transcoder putting should fail - Invalid request data", func(t *testing.T) {
		e := echo.New()
		body, err := GetDummyData(map[string]any{}, map[string]string{"output_type": "ou_put"})
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, "/transcoders", strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", jwtToken)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err = h.PutTranscoder(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
	t.Run("Transcoder putting should fail - Same input and output type", func(t *testing.T) {
		e := echo.New()
		body, err := GetDummyData(map[string]any{"output_type": "mp4", "input_type": "mp4"}, map[string]string{})
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, "/transcoders?id="+transcoder.ID.Hex(), strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", jwtToken)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err = h.PutTranscoder(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
	t.Run("Transcoder putting should fail - Transocder not Found", func(t *testing.T) {
		e := echo.New()
		body, err := GetDummyData(map[string]any{"input_type": "drm"}, map[string]string{})
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, "/transcoders", strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", jwtToken)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err = h.PutTranscoder(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, rec.Code)
	})

	t.Run("Transcoder putting should pass", func(t *testing.T) {
		e := echo.New()
		body, err := GetDummyData(map[string]any{}, map[string]string{})
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, "/transcoders", strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", jwtToken)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err = h.PutTranscoder(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}
