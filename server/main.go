package main

import (
	"contexttest/server/auth"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {
	router := chi.NewRouter()

	router.Use(auth.MiddlewareTest())

	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test route"))
	})
	log.Println("listening on port 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
