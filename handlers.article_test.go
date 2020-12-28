package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Test 항목
// 1. handler가 statusCode 200을 반환하는지
// 2. 반환된 HTML이 title tag, Home Page text를 포함하는지

// home page에 대한 GET request가 인증되지 않은 사용자에게 statusCode 200을 반환하는지 test
func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := getRouter(true)

	r.GET("/", showIndexPage)

	// "/" route에 request 요청
	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK

		// Body 내용을 읽어옴
		p, err := ioutil.ReadAll(w.Body)

		// Test that the page title is "Home Page"
		// 에러가 없고, Body의 내용(p) 안에 "<title>Home Page</title>"이 포함되어 있으면 true 반환
		pageOK := err == nil && strings.Index(string(p),
			"<title>Home Page</title>") > 0

		return statusOK && pageOK
	})

}
