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

func TestPutTranscoder(t *testing.T) {
	t.Run("Transcoder putting should fail - No asset_type  present", func(t *testing.T) {
		e := echo.New()
		body, err := GetDummyData(map[string]any{}, map[string]string{"asset_type": "asset_typ"})
		assert.NoError(t, err)
		fmt.Println(body)
		req := httptest.NewRequest(http.MethodPut, RequestEndPoint, strings.NewReader(body))
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
		body, err := GetDummyData(map[string]any{}, map[string]string{"asset_type": "asse"})
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, RequestEndPoint, strings.NewReader(body))
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
		body, err := GetDummyData(map[string]any{"operation": "encoding"}, map[string]string{})
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, RequestEndPoint, strings.NewReader(body))
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

	t.Run("Transcoder putting should fail - Mongo DB error", func(t *testing.T) {
		e := echo.New()
		body, err := GetDummyData(map[string]any{"input_type": "drm"}, map[string]string{})
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, RequestEndPoint, strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", jwtToken)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = wrongCol
		err = h.PutTranscoder(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})

	t.Run("Transcoder putting should fail - Binding Error", func(t *testing.T) {
		e := echo.New()
		body := "some random string"

		req := httptest.NewRequest(http.MethodPut, RequestEndPoint, strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", jwtToken)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.PutTranscoder(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	})

	t.Run("Transcoder putting should pass", func(t *testing.T) {
		BeforeEach()
		e := echo.New()
		body, err := GetDummyData(map[string]any{}, map[string]string{})
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, RequestEndPoint, strings.NewReader(body))
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
