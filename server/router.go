package server

import (
	"final-project-backend/dto"
	"final-project-backend/handlers"
	"final-project-backend/middlewares"
	"final-project-backend/services"
	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	AuthService services.AuthService
	PostService services.PostService
}

func NewRouter(conf *RouterConfig) *gin.Engine {
	router := gin.Default()

	h := handlers.New(&handlers.HandlerConfig{
		AuthService: conf.AuthService,
		PostService: conf.PostService,
	})

	router.Use(middlewares.ErrorHandler)

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

	router.GET(
		"/posts/",
		middlewares.AuthorizePublic,
		h.GetPosts,
	)
	router.GET(
		"/posts/:slug",
		middlewares.AuthorizePublic,
		h.GetPost,
	)

	router.NoRoute(h.HandleNotFound)

	return router
}
