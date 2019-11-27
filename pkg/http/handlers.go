package http

import (
	"encoding/json"
	"net/http"

	"github.com/LBCC-Technology-Development-Club/site/pkg/blog"
	"github.com/LBCC-Technology-Development-Club/site/pkg/core"
	"github.com/julienschmidt/httprouter"
)

// Service defines a service
type Service struct {
	blog   core.Blog
	Router http.Handler
}

// New provides a new service
func New(blog core.Blog) Service {
	service := Service{
		blog: blog,
	}

	router := httprouter.New()
	router.GET("/blog", service.AllPosts)

	service.Router = UseMiddlewares(router)

	return service
}

// AllPosts endpoint
func (s Service) AllPosts(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	service := blog.NewService(s.blog, r.Context().Value("userId").(string))
	posts, err := service.GetPosts()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}
