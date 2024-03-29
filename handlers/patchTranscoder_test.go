package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestPatchTranscoder(t *testing.T) {

	const requestQuery string = "?input_type=hls&output_type=dash&codec=h264&descriptor=media_analysis"
	t.Run("Transcoder patching should failed, Invalid paramters", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPatch, RequestEndPoint, nil)
		rec := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", jwtToken)
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.PatchTranscoder(c)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.NoError(t, err)
	})
	t.Run("Transcoder patching should failed, no input parameter", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPatch, RequestEndPoint+"?output_type=hls", nil)
		rec := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", jwtToken)
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.PatchTranscoder(c)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.NoError(t, err)
	})
	t.Run("Transcoder patching should failed, no output parameter", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPatch, RequestEndPoint+"?input_type=hls", nil)
		rec := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", jwtToken)
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.PatchTranscoder(c)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.NoError(t, err)
	})
	t.Run("Transcoder patching should failed, invalid request payload", func(t *testing.T) {
		e := echo.New()

		body := `{
			"some_other_param": "some_other_value"
		}`

		req := httptest.NewRequest(http.MethodPatch, RequestEndPoint+requestQuery, strings.NewReader(body))
		rec := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", jwtToken)
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.PatchTranscoder(c)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.NoError(t, err)
	})
	t.Run("Transcoder patching should failed, no transcoder found", func(t *testing.T) {
		e := echo.New()

		body := `{
			"updated_by": "test_user"
		}`

		req := httptest.NewRequest(http.MethodPatch, RequestEndPoint+requestQuery, strings.NewReader(body))
		rec := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", jwtToken)
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.PatchTranscoder(c)
		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.NoError(t, err)
	})

	t.Run("Transcoder patching should failed, Unable to decode", func(t *testing.T) {
		e := echo.New()

		body := "some_invalid_body"

		req := httptest.NewRequest(http.MethodPatch, RequestEndPoint+requestQuery, strings.NewReader(body))
		rec := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", jwtToken)
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.PatchTranscoder(c)
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
		assert.NoError(t, err)
	})
	t.Run("Transcoder patching should failed, Mongo DB error", func(t *testing.T) {
		e := echo.New()

		body := `{
			"updated_by": "test_user"
		}`

		req := httptest.NewRequest(http.MethodPatch, RequestEndPoint+requestQuery, strings.NewReader(body))
		rec := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", jwtToken)
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = wrongCol
		err := h.PatchTranscoder(c)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.NoError(t, err)
	})

	t.Run("Transcoder patching should be done successfully", func(t *testing.T) {
		e := echo.New()

		body := `{
			"updated_by": "test_user"
		}`

		req := httptest.NewRequest(http.MethodPatch, RequestEndPoint+"?input_type=dash&output_type=mp4&codec=h264&descriptor=media_analysis", strings.NewReader(body))
		rec := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", jwtToken)
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.PatchTranscoder(c)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NoError(t, err)
	})

}
