package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLoginPost(t *testing.T) {
	r:= SetupRouter()
	req:=httptest.NewRequest("POST",
		"/login",
		 strings.NewReader(`{"name": "liwenzhou"}`))
	w:=httptest.NewRecorder()
	r.ServeHTTP(w,req)
	assert.Equal(t, http.StatusOK,w.Code)
	// 解析并检验响应内容是否复合预期
	var resp map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &resp)
	assert.Nil(t, err)
		assert.Equal(t,"Liwenzhou", resp["message"])
}
