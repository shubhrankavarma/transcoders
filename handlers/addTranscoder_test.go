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
		body := `
			{
				"input_type": "mp4",
				"output_type": "dash",
				"template_command":"comming soon",
				"status": "active",
				"updated_by": "me"
			}
		`
		req := httptest.NewRequest(http.MethodPost, requestEndPoint, strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", jwtToken)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.AddTranscoder(c)
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.NoError(t, err)
	})
	t.Run(successfulStatus, func(t *testing.T) {
		e := echo.New()
		body := `
			{
				"input_type": "mp4",
				"output_type": "hls",
				"template_command":"comming soon",
				"status": "active",
				"updated_by": "me"
			}
		`
		req := httptest.NewRequest(http.MethodPost, requestEndPoint, strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", jwtToken)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.AddTranscoder(c)
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.NoError(t, err)
	})
	t.Run(successfulStatus, func(t *testing.T) {
		e := echo.New()
		body := `
			{
				"input_type": "dash",
				"output_type": "hls",
				"template_command":"comming soon",
				"status": "active",
				"updated_by": "me"
			}
		`
		req := httptest.NewRequest(http.MethodPost, requestEndPoint, strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", jwtToken)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.AddTranscoder(c)
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.NoError(t, err)
	})
	t.Run("Transcoder adding should fail - Invalid Key", func(t *testing.T) {
		e := echo.New()
		body := `
			{
				"input_type": "dash",
				"output_type": "mp4",
				"template_commnd":"comming soon",
				"status": "active",
				"updated_by": "me"
			}
		`
		req := httptest.NewRequest(http.MethodPost, requestEndPoint, strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", jwtToken)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.AddTranscoder(c)
		assert.NoError(t, err)
	})
	t.Run("Transcoder adding should fail - Invalid input type", func(t *testing.T) {
		e := echo.New()
		body := `
			{
				"input_type": "mp5",
				"output_type": "dash",
				"template_command":"comming soon",
				"status": "active",
				"updated_by": "me"
			}
		`
		req := httptest.NewRequest(http.MethodPost, requestEndPoint, strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", jwtToken)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.AddTranscoder(c)

		// Should give 400 error code
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.NoError(t, err)
	})
	t.Run("Transcoder adding should fail - Invalid output type", func(t *testing.T) {
		e := echo.New()
		body := `
			{
				"input_type": "mp4",
				"output_type": "mp5",
				"template_command":"comming soon",
				"status": "active",
				"updated_by": "me"
			}
		`
		req := httptest.NewRequest(http.MethodPost, requestEndPoint, strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", jwtToken)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.AddTranscoder(c)

		// Should give 400 error code
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.NoError(t, err)
	})
	t.Run("Transcoder adding should fail - already present", func(t *testing.T) {
		e := echo.New()
		body := `
			{
				"input_type": "mp4",
				"output_type": "dash",
				"template_command":"comming soon",
				"status": "active",
				"updated_by": "me"
			}
		`
		req := httptest.NewRequest(http.MethodPost, requestEndPoint, strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", jwtToken)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.AddTranscoder(c)

		// Should give 400 error code
		assert.Equal(t, http.StatusConflict, rec.Code)
		assert.NoError(t, err)
	})
	t.Run("Transcoder adding should fail - input type and output type should not be same", func(t *testing.T) {
		e := echo.New()
		body := `
			{
				"input_type": "mp4",
				"output_type": "mp4",
				"template_command":"comming soon",
				"status": "active",
				"updated_by": "me"
			}
		`
		req := httptest.NewRequest(http.MethodPost, requestEndPoint, strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", jwtToken)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.AddTranscoder(c)

		// Should give 400 error code
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.NoError(t, err)
	})
}
