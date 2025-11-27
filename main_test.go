package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tsutsutsu123/only-login-and-signup/models" // modelsパッケージを追加
)

// TestMain は全てのテスト実行前後に呼ばれるセットアップ/クリーンアップ関数
func TestMain(m *testing.M) {
	// 1. テスト開始前のセットアップ
	// データベース接続を確立し、DB変数を初期化する
	log.Println("Setting up Test Database connection...")
	models.ConnectDataBase()

	if models.DB != nil {
		models.DB.DropTable(&models.User{})
	}
	models.DB.AutoMigrate(&models.User{})

	// 2. テストの実行
	// runTests は os.Exit() を呼び出して終了コードを返す
	code := m.Run()

	// 3. テスト終了後のクリーンアップ（オプション）
	// 必要に応じてDB接続を閉じる、テストデータを削除するなど
	log.Println("Tearing down Test environment.")

	if models.DB != nil {
		models.DB.Close()
	}

	// 実行結果コードを返す
	os.Exit(code)
}

func assertSimpleResponsePOST(t *testing.T, api, data, exmsg string, st int) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", api, strings.NewReader(data))

	router.ServeHTTP(w, req)

	assert.Equal(t, st, w.Code)
	assert.JSONEq(t, exmsg, w.Body.String())
}

func TestFirstEndpointHandler_Normal(t *testing.T) {
	// 正常系
	const expectedSuccessJSON = `{"data": {"username": "testuser_normal", "password": ""}}`
	assertSimpleResponsePOST(
		t,
		"/api/register",
		`{"username": "testuser_normal", "password": "password123"}`,
		expectedSuccessJSON,
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

	// DBのユニーク制約違反テスト (前のテストで作成したユーザーで再試行)
	const expectedDuplicateErrorJSON = `{"error": "Error 1062: Duplicate entry 'testuser_normal' for key 'users.username'"}`
	assertSimpleResponsePOST(
		t,
		"/api/register",
		`{"username": "testuser_normal", "password": "password123"}`,
		expectedDuplicateErrorJSON,
		http.StatusBadRequest)
}
