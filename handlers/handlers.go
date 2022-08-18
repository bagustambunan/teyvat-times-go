package handlers

import (
	"final-project-backend/httperror"
	"final-project-backend/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	authService services.AuthService
	postService services.PostService
}

type HandlerConfig struct {
	AuthService services.AuthService
	PostService services.PostService
}

func New(conf *HandlerConfig) *Handler {
	return &Handler{
		authService: conf.AuthService,
		postService: conf.PostService,
	}
}

func (h *Handler) HandleNotFound(c *gin.Context) {
	_ = c.Error(httperror.NotFoundError())
}
