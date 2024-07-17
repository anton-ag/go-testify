package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRequestStatusOK(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)
	resp := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(resp, req)

	require.Equal(t, resp.Code, http.StatusOK)
	assert.NotEmpty(t, resp.Body)

}

func TestMainHandlerWrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=notmoscow", nil)
	resp := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(resp, req)

	require.Equal(t, resp.Code, http.StatusBadRequest)
	assert.Equal(t, resp.Body.String(), "wrong city value")
}

func TestMainHandlerTooManyCafes(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=20&city=moscow", nil)
	resp := httptest.NewRecorder()
	totaCount := 4

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(resp, req)
	body := resp.Body.String()

	require.Equal(t, resp.Code, http.StatusOK)
	require.NotEmpty(t, body)

	list := strings.Split(body, ",")
	assert.Len(t, list, totaCount)
}
