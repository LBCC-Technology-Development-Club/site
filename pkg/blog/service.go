package blog

import (
	"strconv"

	"github.com/LBCC-Technology-Development-Club/site/pkg/core"
)

// RawPost describes the structure of raw posts that we get from the client
type RawPost struct {
	PostID    int64  `json:"pid"`
	Title     string `json:"title"`
	Summary   string `json:"summary"`
	Body      string `json:"body"`
	Timestamp string `json:"timestamp"`
}

// Service stores current user ID and blog
type Service struct {
	userID string
	blog   core.Blog
}

// GetPosts returns all posts
func (s Service) GetPosts() ([]*core.Post, error) {
	return s.blog.FindAllPosts()
}

// CreatePost creates a new post
func (s Service) CreatePost(rawPost RawPost) (*core.Post, error) {
	post := s.rawPostToPost(rawPost)
	err := s.blog.CreatePost(post)
	if err != nil {
		return nil, err
	}
	return post, nil
}

// rawPostToPost translates rawposts to posts
func (s Service) rawPostToPost(rawPost RawPost) *core.Post {
	return &core.Post{
		UserID:    s.userID,
		PostID:    strconv.Itoa(int(rawPost.PostID)),
		Title:     rawPost.Title,
		Summary:   rawPost.Summary,
		Body:      rawPost.Body,
		Timestamp: rawPost.Timestamp,
	}
}

// NewService provides a new service
func NewService(blog core.Blog, userID string) Service {
	return Service{
		blog:   blog,
		userID: userID,
	}
}
