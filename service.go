package main

import (
	"log"
	"net/http"

	"github.com/LBCC-Technology-Development-Club/site/blog"
	"github.com/LBCC-Technology-Development-Club/site/uAuth"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	// Blog endpoints
	router.Route("/", func(r chi.Router) {
		r.Mount("/api/blog", blog.Routes())
	})

	// Login endpoints
	router.Route("/", func(r chi.Router) {
		r.Mount("/api/login", uAuth.Routes())
	})

	return router
}

func main() {
	router := Routes()

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error())
	}

	log.Fatal(http.ListenAndServe(":8080", router))
}
