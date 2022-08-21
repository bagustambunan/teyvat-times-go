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

	router.NoRoute(h.HandleNotFound)

	return router
}
