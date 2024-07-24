package server_test

import (
	"net/http"
	"testing"

	"github.com/cjbagley/colinbagley.dev/internal/server"
)

func TestRoutes(t *testing.T) {
	t.Run("Homepage loads", func(t *testing.T) {
		res := server.MockServerRequest(http.MethodGet, "/", nil)
		expectedStatus := http.StatusOK

		if res.Code != expectedStatus {
			t.Errorf("Status Code: got %d, want %d", res.Code, expectedStatus)
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
