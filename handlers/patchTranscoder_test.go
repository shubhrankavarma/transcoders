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

	const requestQuery string = "?operation=media_analysis&asset_type=video"
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
	t.Run("Transcoder patching should failed, no asset_type parameter", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPatch, RequestEndPoint+"?operation=media_analysis", nil)
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
	t.Run("Transcoder patching should failed, no operation parameter", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPatch, RequestEndPoint+"?asset_type=video", nil)
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
		BeforeEach()
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
		BeforeEach()
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
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NoError(t, err)
	})

}
