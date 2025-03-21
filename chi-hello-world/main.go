package main

import (
    "net/http"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)


func main() {
    s := CreateNewServer()
    s.MountHandlers()
    http.ListenAndServe(":3000", s.Router)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello World!"))
}

type Server struct {
    Router *chi.Mux
}

func CreateNewServer() *Server {
    s := &Server{}
    s.Router = chi.NewRouter()
    return s
}

func (s *Server) MountHandlers() {
    s.Router.Use(middleware.Logger)

    s.Router.Get("/", HelloWorld)
}
