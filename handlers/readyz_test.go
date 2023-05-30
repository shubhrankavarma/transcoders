package handlers

import (
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestReadyz(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()

	// Create a new request and recorder
	req := httptest.NewRequest(http.MethodGet, "/readyz", nil)
	rec := httptest.NewRecorder()

	// Create a new TranscoderHandler instance
	h := &TranscoderHandler{}

	// Test readiness check when IsReady is nil
	h.Readyz(e.NewContext(req, rec))
	// assert.Error(t, err)
	assert.Equal(t, http.StatusServiceUnavailable, rec.Code)

	// Test readiness check when IsReady is false
	h.IsReady = &atomic.Value{}
	h.IsReady.Store(false)
	h.Readyz(e.NewContext(req, rec))
	// assert.Error(t, err)
	assert.Equal(t, http.StatusServiceUnavailable, rec.Code)

	// Test readiness check when IsReady is true
	h.IsReady.Store(true)
	rec = httptest.NewRecorder()
	h.Readyz(e.NewContext(req, rec))
	// assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

}
