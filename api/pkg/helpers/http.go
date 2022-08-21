package helpers

import (
	"io"
	"net/http"
	"net/http/httptest"
)

func NewTestRequest(method, target string, body io.Reader) *http.Request {
	req := httptest.NewRequest(method, target, body)
	req.Header.Set("Content-Type", "application/json")
	return req
}
