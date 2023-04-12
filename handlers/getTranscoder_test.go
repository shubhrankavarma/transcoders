package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetTranscoder(t *testing.T) {
	t.Run("Transcoder should be fetched successfully", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/transcoders", nil)
		rec := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", jwtToken)
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.GetTranscoders(c)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NoError(t, err)
	})
	t.Run("Transcoder should be fetched successfully - by page_size 1", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/transcoders?page_size=1", nil)
		rec := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", jwtToken)
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.GetTranscoders(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		// Check size of the returned array is one
		var results []Transcoder
		err = json.Unmarshal(rec.Body.Bytes(), &results)
		assert.NoError(t, err)
		assert.Len(t, results, 1)
	})
	t.Run("Transcoder should be fetched successfully - by limit 2", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/transcoders?page_size=2", nil)
		rec := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", jwtToken)
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.GetTranscoders(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		// Check size of the returned array is two
		var results []Transcoder
		err = json.Unmarshal(rec.Body.Bytes(), &results)
		assert.NoError(t, err)
		assert.Len(t, results, 2)
	})
	t.Run("Transcoder should be fetched successfully - by input_type", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/transcoders?input_type=dash", nil)
		rec := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", jwtToken)
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.GetTranscoders(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		// Check the input type of the returned array is dash
		var results []Transcoder
		err = json.Unmarshal(rec.Body.Bytes(), &results)
		assert.NoError(t, err)
		assert.Equal(t, results[0].InputType, "dash")
	})
	t.Run("Transcoder should be fetched successfully - by output_type", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/transcoders?output_type=hls", nil)
		rec := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", jwtToken)
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.GetTranscoders(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		// Check the output type of the returned array is hls
		var results []Transcoder
		err = json.Unmarshal(rec.Body.Bytes(), &results)
		assert.NoError(t, err)
		assert.Equal(t, results[0].OutputType, "hls")
	})
	t.Run("Transcoder should be fetched successfully - by output_type and input_type", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/transcoders?input_type=dash&output_type=hls", nil)
		rec := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", jwtToken)
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.GetTranscoders(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		// Check the output type of the returned array is hls and input type is dash
		var results []Transcoder
		err = json.Unmarshal(rec.Body.Bytes(), &results)
		assert.NoError(t, err)
		assert.Equal(t, results[0].OutputType, "hls")
		assert.Equal(t, results[0].InputType, "dash")
	})

	t.Run("Transcoder should be fetched successfully - by codec", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/transcoders?codec=h264", nil)
		rec := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", jwtToken)
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.GetTranscoders(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		// Check the output type of the returned array is hls and input type is dash
		var results []Transcoder
		err = json.Unmarshal(rec.Body.Bytes(), &results)
		assert.NoError(t, err)
		assert.Equal(t, results[0].Codec, "h264")
	})

	t.Run("Transcoder should be fetched successfully - by descriptor", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/transcoders?descriptor=media_analysis", nil)
		rec := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", jwtToken)
		c := e.NewContext(req, rec)
		h := &TranscoderHandler{}
		h.Col = transcoderCol
		err := h.GetTranscoders(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		// Check the output type of the returned array is hls and input type is dash
		var results []Transcoder
		err = json.Unmarshal(rec.Body.Bytes(), &results)
		assert.NoError(t, err)
		assert.Equal(t, results[0].Descriptor, "media_analysis")
	})

}
