package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder() // Responsible to collect the response
	req, _ := http.NewRequest("GET", "/hello", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "World", w.Body.String())

}

func TestPostUser(t *testing.T) {
	router := setupRouter()
	router = postRouter(router)

	w := httptest.NewRecorder()

	exampleUser := User{
		Username: "Antonio",
		Gender: "Male",
	}
	
	userJson, _ := json.Marshal(exampleUser)
	req, _ := http.NewRequest("POST", "/user/add", strings.NewReader(string(userJson)))

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(userJson), w.Body.String())

}
