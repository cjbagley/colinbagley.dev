package server_test

import (
	"net/http"
	"testing"

	"github.com/cjbagley/colinbagley.dev/internal/middleware"
	"github.com/cjbagley/colinbagley.dev/internal/server"
)

func TestServer(t *testing.T) {
	t.Run("applies middleware", func(t *testing.T) {
		res := server.MockServerRequest(http.MethodGet, "/", nil)
		expectedStatus := http.StatusOK

		if res.Code != expectedStatus {
			t.Errorf("Status Code: got %d, want %d", res.Code, expectedStatus)
		}
		if res.Header().Get(middleware.SiteHeader) != middleware.SiteValue {
			t.Errorf("Site header not found, got '%s'", res.Header().Get(middleware.SiteHeader))
		}
		if res.Header().Get("Content-Type") != "text/html" {
			t.Errorf("Content Type header expected 'text/html', got '%s'", res.Header().Get("Content-Type"))
		}
	})
}
