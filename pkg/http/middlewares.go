package http

import (
	"log"
	"net/http"

	"github.com/rs/cors"
)

// JSONApi sets the content type for JSON
func JSONApi(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		h.ServeHTTP(w, r)
	})
}

// Cors sets up proper cors
func Cors(h http.Handler) http.Handler {
	corsConfig := cors.New(cors.Options{
		AllowedHeaders: []string{"Origin", "Accept", "Content-Type", "X-Requested-With", "Authorization"},
		AllowedMethods: []string{"POST", "PUT", "GET", "PATCH", "OPTIONS", "HEAD", "DELETE"},
		Debug:          true,
	})
	return corsConfig.Handler(h)
}

// AccessLog logs requests
func AccessLog(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s: %s\n", r.Method, r.RequestURI)
	})
}

//UseMiddlewares defines which middlewares to use
func UseMiddlewares(h http.Handler) http.Handler {
	h = JSONApi(h)
	h = Cors(h)
	return AccessLog(h)
}
