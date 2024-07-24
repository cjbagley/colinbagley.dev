package server

import (
	"io"
	"net/http/httptest"
)

func MockServerRequest(method string, url string, body io.Reader) *httptest.ResponseRecorder {
	server := NewServer()

	res := httptest.NewRecorder()
	req := httptest.NewRequest(method, url, body)
	server.Handler.ServeHTTP(res, req)

	return res
}
