package hello_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/iotassss/saitamarental/internal/hello"
	"github.com/stretchr/testify/assert"
)

func TestHelloHandler(t *testing.T) {
	// Gin のテスト用のコンテキストを作成
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/hello", hello.HelloHandler)

	// HTTPリクエストを作成
	req, _ := http.NewRequest("GET", "/hello", nil)
	w := httptest.NewRecorder()

	// リクエストを実行
	r.ServeHTTP(w, req)

	// レスポンスの確認
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message": "Hello, World!"}`, w.Body.String())
}
