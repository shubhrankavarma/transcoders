package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestAddTranscoder(t *testing.T) {

	var successfulStatus string = "Transcoder should be added successfully"
	var requestEndPoint string = "/transcoders"

	t.Run(successfulStatus, func(t *testing.T) {
		e := echo.New()
		// Convert the body to string
		body, err := GetDummyData(map[string]any{}, map[string]string{})
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, requestEndPoint, strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", jwtToken)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		h.Cfg = cfg
		err = h.AddTranscoder(c)
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.NoError(t, err)

	})

	t.Run(successfulStatus, func(t *testing.T) {
		e := echo.New()
		body, err := GetDummyData(map[string]any{"asset_type": "video"}, map[string]string{})
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, requestEndPoint, strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", jwtToken)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err = h.AddTranscoder(c)
		assert.Equal(t, http.StatusConflict, rec.Code)
		assert.NoError(t, err)
	})

	t.Run(successfulStatus, func(t *testing.T) {
		e := echo.New()
		body, err := GetDummyData(map[string]any{"input_type": "hls"}, map[string]string{})
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, requestEndPoint, strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", jwtToken)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err = h.AddTranscoder(c)
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.NoError(t, err)
	})

	t.Run("Transcoder adding should fail - Invalid Key - Template Command", func(t *testing.T) {
		e := echo.New()
		body, err := GetDummyData(map[string]any{}, map[string]string{"template_commnd": "tempte_command"})
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, requestEndPoint, strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", jwtToken)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err = h.AddTranscoder(c)
		assert.NoError(t, err)
	})
	t.Run("Transcoder adding should fail - already present", func(t *testing.T) {
		e := echo.New()
		body, err := GetDummyData(map[string]any{}, map[string]string{})
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, requestEndPoint, strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", jwtToken)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err = h.AddTranscoder(c)

		// Should give 409 error code
		assert.Equal(t, http.StatusConflict, rec.Code)
		assert.NoError(t, err)
	})

	t.Run("Transcoder adding should fail - Mongo DB Error", func(t *testing.T) {
		e := echo.New()

		body, err := GetDummyData(map[string]any{"input_type": "mp4", "output_type": "hls"}, map[string]string{})
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, requestEndPoint, strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", jwtToken)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = wrongCol
		err = h.AddTranscoder(c)

		assert.NoError(t, err)

		// Should give 500 error code
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})

	t.Run("Transcoder adding should fail - Binding Error", func(t *testing.T) {
		e := echo.New()

		body := "A test string"

		req := httptest.NewRequest(http.MethodPost, requestEndPoint, strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", jwtToken)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.AddTranscoder(c)

		assert.NoError(t, err)

		// Should give 422 error code
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	})
}
