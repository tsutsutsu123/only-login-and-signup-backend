package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"github.com/stretchr/testify/assert"
)

func assertSimpleResponsePOST(t *testing.T, api,  data, exmsg string, st int) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", api, strings.NewReader(data))

	router.ServeHTTP(w, req)

	assert.Equal(t, st, w.Code)
	assert.JSONEq(t, exmsg, w.Body.String())
}

func TestFirstEndpointHandler_Normal(t *testing.T) {
	// 正常系
	assertSimpleResponsePOST(
		t,
		"/api/register",
		`{"username": "testuser", "password": "password123"}`,
		`{"data":"validated!"}`,
		http.StatusOK)

	// 準正常系
	assertSimpleResponsePOST(
		t,
		"/api/register",
		`{"username": "hoge"}`,
		`{"error":"Key: 'RegisterInput.Password' Error:Field validation for 'Password' failed on the 'required' tag"}`,
		http.StatusBadRequest)
	assertSimpleResponsePOST(
		t,
		"/api/register",
		`{"password": "passwd"}`,
		`{"error":"Key: 'RegisterInput.Username' Error:Field validation for 'Username' failed on the 'required' tag"}`,
		http.StatusBadRequest)
}
