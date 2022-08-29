package handlers

import (
	"git.garena.com/sea-labs-id/batch-01/bagus-tambunan/final-project-backend/httperror"
	"git.garena.com/sea-labs-id/batch-01/bagus-tambunan/final-project-backend/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	authService         services.AuthService
	userService         services.UserService
	postService         services.PostService
	subscriptionService services.SubscriptionService
	transactionService  services.TransactionService
	voucherService      services.VoucherService
	giftService         services.GiftService
}

type HandlerConfig struct {
	AuthService         services.AuthService
	UserService         services.UserService
	PostService         services.PostService
	SubscriptionService services.SubscriptionService
	TransactionService  services.TransactionService
	VoucherService      services.VoucherService
	GiftService         services.GiftService
}

func New(conf *HandlerConfig) *Handler {
	return &Handler{
		authService:         conf.AuthService,
		userService:         conf.UserService,
		postService:         conf.PostService,
		subscriptionService: conf.SubscriptionService,
		transactionService:  conf.TransactionService,
		voucherService:      conf.VoucherService,
		giftService:         conf.GiftService,
	}
}

func (h *Handler) HandleNotFound(c *gin.Context) {
	_ = c.Error(httperror.NotFoundError())
}
