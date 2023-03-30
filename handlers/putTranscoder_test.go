package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var transcoder Transcoder

func TestPutTranscoder(t *testing.T) {
	t.Run("Transcoder should be added successfully", func(t *testing.T) {
		e := echo.New()
		body := `
			{
				"input_type": "hls",
				"output_type": "dash",
				"template_command":"comming soon",
				"status": "active",
				"updated_by": "me"
			}
		`
		req := httptest.NewRequest(http.MethodPost, "/transcoders", strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.AddTranscoder(c)
		assert.Equal(t, http.StatusCreated, rec.Code)
		h.Col.FindOne(req.Context(), map[string]interface{}{"input_type": "hls", "output_type": "dash"}).Decode(&transcoder)
		assert.NoError(t, err)
	})
	t.Run("Transcoder putting should fail - No Id present", func(t *testing.T) {
		e := echo.New()
		body := `
			{
				"input_type": "dash",
				"output_type": "mp4",
				"template_command":"comming soon",
				"status": "active",
				"updated_by": "me"
			}
		`
		req := httptest.NewRequest(http.MethodPut, "/transcoders", strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.PutTranscoder(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
	t.Run("Transcoder putting should fail - Invalid request data", func(t *testing.T) {
		e := echo.New()
		body := `
			{
				"input_type": "dash",
				"outpu_type": "mp4",
				"template_command":"comming soon",
				"status": "active",
				"updated_by": "me"
			}
		`
		req := httptest.NewRequest(http.MethodPut, "/transcoders/id="+transcoder.ID.Hex(), strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.PutTranscoder(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
	t.Run("Transcoder putting should fail - Same input and output type", func(t *testing.T) {
		e := echo.New()
		body := `
			{
				"input_type": "mp4",
				"outpu_type": "mp4",
				"template_command":"comming soon",
				"status": "active",
				"updated_by": "me"
			}
		`
		req := httptest.NewRequest(http.MethodPut, "/transcoders?id="+transcoder.ID.Hex(), strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.PutTranscoder(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
	t.Run("Transcoder putting should pass", func(t *testing.T) {
		e := echo.New()
		body := `
			{
				"input_type": "ts",
				"output_type": "mp4",
				"template_command":"comming soon",
				"status": "active",
				"updated_by": "me"
			}
		`
		req := httptest.NewRequest(http.MethodPut, "/transcoders?id="+transcoder.ID.Hex(), strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.PutTranscoder(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}
