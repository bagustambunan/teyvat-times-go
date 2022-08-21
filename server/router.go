package server

import (
	"final-project-backend/dto"
	"final-project-backend/handlers"
	"final-project-backend/middlewares"
	"final-project-backend/services"
	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	AuthService         services.AuthService
	UserService         services.UserService
	PostService         services.PostService
	SubscriptionService services.SubscriptionService
	VoucherService      services.VoucherService
}

func NewRouter(conf *RouterConfig) *gin.Engine {
	router := gin.Default()

	h := handlers.New(&handlers.HandlerConfig{
		AuthService:         conf.AuthService,
		UserService:         conf.UserService,
		PostService:         conf.PostService,
		SubscriptionService: conf.SubscriptionService,
		VoucherService:      conf.VoucherService,
	})

	router.Use(middlewares.ErrorHandler)

	// AUTH
	router.POST(
		"/sign-up",
		middlewares.RequestValidator(&dto.SignUpReq{}),
		h.SignUp,
	)
	router.POST(
		"/sign-in",
		middlewares.RequestValidator(&dto.SignInReq{}),
		h.SignIn,
	)

	// ADMIN > POST
	router.GET(
		"/posts/",
		middlewares.AuthorizeInternal,
		h.GetPosts,
	)
	router.GET(
		"/posts/:postID",
		middlewares.AuthorizeInternal,
		h.GetPost,
	)
	router.POST(
		"/posts",
		middlewares.AuthorizeInternal,
		middlewares.RequestValidator(&dto.PostReq{}),
		h.AddPost,
	)

	// PUBLIC > POST
	router.GET(
		"/pub/posts/",
		middlewares.AuthorizePublic,
		h.GetPosts,
	)
	router.GET(
		"/pub/posts/:slug",
		middlewares.AuthorizePublic,
		h.PubReadPost,
	)
	router.POST(
		"/pub/posts/:postID/activities",
		middlewares.AuthorizePublic,
		middlewares.RequestValidator(&dto.ActivityReq{}),
		h.PubPostActivity,
	)
	router.POST(
		"/pub/posts/:postID/unlocks",
		middlewares.AuthorizePublic,
		h.PubPostUnlock,
	)

	// PUBLIC > SUBSCRIPTION
	router.POST(
		"/pub/transactions",
		middlewares.AuthorizePublic,
		middlewares.RequestValidator(&dto.TransactionReq{}),
		h.AddTransaction,
	)
	router.POST(
		"/pub/subscriptions",
		middlewares.AuthorizePublic,
		h.TestAddUserSubscription,
	)

	router.NoRoute(h.HandleNotFound)

	return router
}
