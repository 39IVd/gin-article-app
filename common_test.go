package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var tmpArticleList []article

// TestMain : test function들이 실행되기 전 Gin을 setup하는 함수
func TestMain(m *testing.M) {
	// gin을 TestMode로 Set
	gin.SetMode(gin.TestMode)

	// 나머지 test함수들 실행
	os.Exit(m.Run())
}

// getRouter : test가 진행되는 동안 router를 생성, 반환하는 Helper Function
func getRouter(withTemplates bool) *gin.Engine {
	// 1. create router
	r := gin.Default()
	if withTemplates {
		// 2. Load all the templates
		// /templates/ 안의 모든 파일을 load
		// 한번 로드되면 다시 재로드하지 않음 -> 성능 향상
		r.LoadHTMLGlob("templates/*")
	}
	return r
}

// testHTTPResponse : request를 처리하고 response를 test하는 Helper Function
// test가 성공적인지를 반환하는 함수 f를 실행
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {

	// response recorder 생성
	w := httptest.NewRecorder()

	// service 생성 및 request 처리
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

/*
HTTP Code와 반환된 HTML을 확인하는 과정
1. 새 router 생성
2. main app이 사용하는 handler를 사용할 route 정의 (showIndexPage)
3. 해당 route에 접근할 request 생성
4. HTTP code와 HTML을 test하기 위해, response를 처리하는 함수 생성
5. 이 함수를 인자로 갖는 testHTTPResponse()를 호출하여 test 종료
*/

// saveLists : main article list를 임시 article list에 저장
func saveLists() {
	tmpArticleList = articleList
}

// restoreLists : 임시 article list를 main article list로 복구
func restoreLists() {
	articleList = tmpArticleList
}
