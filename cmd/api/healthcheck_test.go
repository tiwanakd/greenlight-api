package main

import (
	"net/http"
	"testing"

	"github.com/tiwanakd/greenlight-api/internal/assert"
)

func TestHealthCheckHandler(t *testing.T) {
	app := newTestApplication()
	ts := newTestServer(app.routes())
	defer ts.Close()

	code, headers, body := ts.get(t, "/v1/healthcheck")

	t.Run("CheckResponseCode", func(t *testing.T) {
		assert.Equal(t, code, http.StatusOK)
	})

	t.Run("CheckJSONHeader", func(t *testing.T) {
		assert.Equal(t, headers.Get("Content-Type"), "application/json")
	})

	t.Run("CheckResponseBody", func(t *testing.T) {
		assert.StringContains(t, body, `"status": "available"`)
	})
}
