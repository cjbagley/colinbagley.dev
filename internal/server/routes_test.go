package server_test

import (
	"net/http"
	"testing"

	"github.com/cjbagley/colinbagley.dev/internal/server"
)

func TestRoutes(t *testing.T) {
	t.Run("All pages load", func(t *testing.T) {
		routes := server.GetRoutes()

		for _, r := range routes {
			t.Logf("Testing route: %s %s", r.Method, r.Path)
			res := server.MockServerRequest(r.Method, r.Path, nil)

			expectedStatus := http.StatusOK

			if res.Code != expectedStatus {
				t.Errorf("Status Code: got %d, want %d", res.Code, expectedStatus)
			}

			if res.Body.String() == "" {
				t.Errorf("Page for '%s' has no body content", r.Path)
			}
		}
	})
	t.Run("404 on unmatched route", func(t *testing.T) {
		res := server.MockServerRequest(http.MethodGet, "/this-page-does-not-exist-asdf", nil)
		expectedStatus := http.StatusNotFound

		if res.Code != expectedStatus {
			t.Errorf("Status Code: got %d, want %d", res.Code, expectedStatus)
		}
	})
}
