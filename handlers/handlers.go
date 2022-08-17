package handlers

import (
	"final-project-backend/httperror"
	"final-project-backend/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	authService services.AuthService
}

type HandlerConfig struct {
	AuthService services.AuthService
}

func New(conf *HandlerConfig) *Handler {
	return &Handler{
		authService: conf.AuthService,
	}
}

func (h *Handler) HandleNotFound(c *gin.Context) {
	_ = c.Error(httperror.NotFoundError())
}
