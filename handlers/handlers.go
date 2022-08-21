package handlers

import (
	"final-project-backend/httperror"
	"final-project-backend/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	authService         services.AuthService
	userService         services.UserService
	postService         services.PostService
	subscriptionService services.SubscriptionService
}

type HandlerConfig struct {
	AuthService         services.AuthService
	UserService         services.UserService
	PostService         services.PostService
	SubscriptionService services.SubscriptionService
}

func New(conf *HandlerConfig) *Handler {
	return &Handler{
		authService:         conf.AuthService,
		userService:         conf.UserService,
		postService:         conf.PostService,
		subscriptionService: conf.SubscriptionService,
	}
}

func (h *Handler) HandleNotFound(c *gin.Context) {
	_ = c.Error(httperror.NotFoundError())
}
