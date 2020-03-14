package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rkorkosz/events/pkg/event"
)

func main() {
	r := chi.NewRouter()
	store, err := event.NewBoltStore("events.db", nil)
	if err != nil {
		log.Fatal(err)
	}
	r.Use(middleware.DefaultLogger)
	r.Use(middleware.Recoverer)
	r.Use(event.Middleware(store))
	r.Route("/v1", func(r chi.Router) {
		r.Put("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("w"))
		})
	})
	log.Fatal(http.ListenAndServe(":8000", r))
}
