package main

import (
	"log"
	"net/http"

	"github.com/LBCC-Technology-Development-Club/site/pkg/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// Routes defines all routes/endpoints for the API. Split into /blog and /login
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
	router.Route("/blog", func(r chi.Router) {
		r.Mount("/", service.BlogRoutes())
	})

	// Login endpoints
	router.Route("/login", func(r chi.Router) {
		r.Mount("/", service.LoginRoutes())
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
