package routes

import (
    "net/http"
    "github.com/go-chi/chi/v5"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if err := h(w, r); err != nil {
        // handle returned err here
        w.WriteHeader(503)
        w.Write([]byte("Bad"))
    }
}


func ConfigureRoutes(router *chi.Mux) {
    // https://go-chi.io/#/pages/routing
    router.Method(http.MethodGet, "/", Handler(IndexHandler))
}



    