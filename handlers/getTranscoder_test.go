package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/amagimedia/transcoders/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetTranscoder(t *testing.T) {
	t.Run("Transcoder should be fetched successfully", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, RequestEndPoint, nil)
		rec := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.GetTranscoders(c)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NoError(t, err)
	})

	t.Run("Transcoder should be fetched successfully - by page_size 1", func(t *testing.T) {
		BeforeEach()
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, RequestEndPoint+"?page_size=1", nil)
		rec := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.GetTranscoders(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		// Check size of the returned array is one
		var results []models.Transcoder
		err = json.Unmarshal(rec.Body.Bytes(), &results)
		assert.NoError(t, err)
		assert.Len(t, results, 1)
	})
	t.Run("Transcoder should be fetched successfully - by limit 2", func(t *testing.T) {
		BeforeEach()
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, RequestEndPoint+"?page_size=2", nil)
		rec := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.GetTranscoders(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		// Check size of the returned array is two
		var results []models.Transcoder
		err = json.Unmarshal(rec.Body.Bytes(), &results)
		assert.NoError(t, err)
		assert.Len(t, results, 2)
	})
	t.Run("Transcoder should be fetched successfully - by input_type", func(t *testing.T) {
		BeforeEach()
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, RequestEndPoint+"?input_type=dash", nil)
		rec := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.GetTranscoders(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		// Check the input type of the returned array is dash
		var results []models.Transcoder
		err = json.Unmarshal(rec.Body.Bytes(), &results)
		assert.NoError(t, err)
		assert.Equal(t, results[0].InputType, "dash")
	})
	t.Run("Transcoder should be fetched successfully - by output_type", func(t *testing.T) {
		BeforeEach()
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, RequestEndPoint+"?output_type=mp4", nil)
		rec := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.GetTranscoders(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		// Check the output type of the returned array is mp4
		var results []models.Transcoder
		err = json.Unmarshal(rec.Body.Bytes(), &results)
		assert.NoError(t, err)
		assert.Equal(t, results[0].OutputType, "mp4")
	})
	t.Run("Transcoder should be fetched successfully - by output_type and input_type", func(t *testing.T) {
		BeforeEach()
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, RequestEndPoint+"?input_type=dash&output_type=mp4", nil)
		rec := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.GetTranscoders(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		// Check the output type of the returned array is mp4 and input type is dash
		var results []models.Transcoder
		err = json.Unmarshal(rec.Body.Bytes(), &results)
		assert.NoError(t, err)
		assert.Equal(t, results[0].OutputType, "mp4")
		assert.Equal(t, results[0].InputType, "dash")
	})

	t.Run("Transcoder should be fetched successfully - by operation", func(t *testing.T) {
		BeforeEach()
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, RequestEndPoint+"?operation=media_analysis", nil)
		rec := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.GetTranscoders(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		// Check the output type of the returned array is hls and input type is dash
		var results []models.Transcoder
		err = json.Unmarshal(rec.Body.Bytes(), &results)
		assert.NoError(t, err)
		assert.Equal(t, results[0].Operation, "media_analysis")
	})

	t.Run("Transcoder should be fetched successfully - by page = 2", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, RequestEndPoint+"?page=2", nil)
		rec := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.GetTranscoders(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
	t.Run("Transcoder fetching failed - wrong page", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, RequestEndPoint+"?page=a", nil)
		rec := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.GetTranscoders(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
	t.Run("Transcoder fetching failed - wrong limit", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, RequestEndPoint+"?page_size=a", nil)
		rec := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.GetTranscoders(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
	t.Run("Transcoder fetching failed - Mongo DB error", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, RequestEndPoint, nil)
		rec := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = wrongCol
		err := h.GetTranscoders(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})

}
