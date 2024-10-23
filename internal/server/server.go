package server

import (
	"crossword/cmd/api"
	"log"
	"net/http"
)

// StartServer starting  HTTP-server
func NewServer(addr string) *http.Server {
	router := api.SetupRoutes()

	// Логируем все запросы (middleware)
	loggedRouter := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Method, r.RequestURI, r.RemoteAddr)
		router.ServeHTTP(w, r)
	})

	return &http.Server{
		Addr:    addr,
		Handler: loggedRouter,
	}
}
