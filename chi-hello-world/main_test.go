package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func executeRequest(req *http.Request, s *Server) *httptest.ResponseRecorder {
    rr := httptest.NewRecorder()

    s.Router.ServeHTTP(rr, req)
    
    return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
    if expected != actual {
        t.Errorf("Got %d but want %d", actual, expected)
    }
}

func TestHelloWorld(t *testing.T) {
    s := CreateNewServer()

    s.MountHandlers()

    req, _ := http.NewRequest("GET", "/", nil)

    response := executeRequest(req, s)

    checkResponseCode(t, http.StatusOK, response.Code)
    
    if "Hello World!" != response.Body.String() {
        t.Errorf("Got %v esnt %v", response.Body.String(), "Hello World!")
    }
}
