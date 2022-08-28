package server

import (
	"final-project-backend/dto"
	"final-project-backend/handlers"
	"final-project-backend/middlewares"
	"final-project-backend/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	AuthService         services.AuthService
	UserService         services.UserService
	PostService         services.PostService
	SubscriptionService services.SubscriptionService
	TransactionService  services.TransactionService
	VoucherService      services.VoucherService
	GiftService         services.GiftService
}

func NewRouter(conf *RouterConfig) *gin.Engine {
	router := gin.Default()
	//router.Use(cors.New(cors.Config{
	//	AllowOrigins: []string{"*"},
	//	AllowMethods: []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
	//	AllowHeaders: []string{"Origin", "Content-Type", "X-Auth-Token", "Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Authorization"},
	//	MaxAge:       12 * time.Hour,
	//}))
	//router.Use(cors.Default())
	config := cors.DefaultConfig()
	//config.AllowOrigins = []string{"*"}
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Authorization", "Content-Type"}
	router.Use(cors.New(config))

	h := handlers.New(&handlers.HandlerConfig{
		AuthService:         conf.AuthService,
		UserService:         conf.UserService,
		PostService:         conf.PostService,
		SubscriptionService: conf.SubscriptionService,
		TransactionService:  conf.TransactionService,
		VoucherService:      conf.VoucherService,
		GiftService:         conf.GiftService,
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

	// ADMIN > CATEGORY
	router.GET(
		"/categories/:postCategoryID",
		middlewares.AuthorizeInternal,
		h.GetCategory,
	)
	router.POST(
		"/categories",
		middlewares.AuthorizeInternal,
		middlewares.RequestValidator(&dto.CategoryReq{}),
		h.AddCategory,
	)
	router.PUT(
		"/categories/:postCategoryID",
		middlewares.AuthorizeInternal,
		middlewares.RequestValidator(&dto.CategoryReq{}),
		h.UpdateCategory,
	)

	// ADMIN > POST
	router.GET(
		"/tiers/:postTierID",
		middlewares.AuthorizeInternal,
		h.GetTier,
	)
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
	router.PUT(
		"/posts/:postID",
		middlewares.AuthorizeInternal,
		middlewares.RequestValidator(&dto.PostReq{}),
		h.UpdatePost,
	)
	router.DELETE(
		"/posts/:postID",
		middlewares.AuthorizeInternal,
		h.DeletePost,
	)
	router.POST(
		"/posts",
		middlewares.AuthorizeInternal,
		middlewares.RequestValidator(&dto.PostReq{}),
		h.AddPost,
	)

	// ADMIN > TRANSACTION
	router.GET(
		"transactions/",
		middlewares.AuthorizeInternal,
		h.GetTransactions,
	)
	router.GET(
		"transaction-statuses/",
		middlewares.AuthorizeInternal,
		h.GetTransactionStatuses,
	)

	// ADMIN > VOUCHER
	router.GET(
		"/vouchers",
		middlewares.AuthorizeInternal,
		h.GetVouchers,
	)

	// PUBLIC > USER
	router.GET(
		"/pub/users/:userID",
		middlewares.AuthorizePublic,
		h.GetUser,
	)
	router.GET(
		"/pub/users/:userID/downlines",
		middlewares.AuthorizePublic,
		h.GetUserDownLines,
	)

	// PUBLIC > POST
	router.GET(
		"/pub/tiers/",
		middlewares.AuthorizePublic,
		h.GetTiers,
	)
	router.GET(
		"/pub/categories/",
		middlewares.AuthorizePublic,
		h.GetCategories,
	)
	router.GET(
		"/pub/posts/",
		middlewares.AuthorizePublic,
		h.GetPosts,
	)
	router.GET(
		"/pub/posts/slug/:slug",
		middlewares.AuthorizePublic,
		h.PubReadPost,
	)
	router.GET(
		"/pub/posts/:postID/activities",
		middlewares.AuthorizePublic,
		h.PubGetActivity,
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
	router.GET(
		"/pub/subscriptions/",
		middlewares.AuthorizePublic,
		h.GetSubscriptions,
	)
	router.GET(
		"/pub/subscriptions/:subscriptionID",
		middlewares.AuthorizePublic,
		h.GetSubscription,
	)
	router.POST(
		"/pub/user_subscriptions",
		middlewares.AuthorizePublic,
		h.TestAddUserSubscription,
	)

	// PUBLIC > TRANSACTION
	router.GET(
		"/pub/user_subscriptions/date/new",
		middlewares.AuthorizePublic,
		h.GetUserNewSubscriptionDate,
	)
	router.POST(
		"/pub/transactions",
		middlewares.AuthorizePublic,
		middlewares.RequestValidator(&dto.TransactionReq{}),
		h.AddTransaction,
	)
	router.GET(
		"/pub/transactions/",
		middlewares.AuthorizePublic,
		h.GetUserTransactions,
	)

	// PUBLIC > VOUCHER
	router.GET(
		"/pub/user_voucher/:code",
		middlewares.AuthorizePublic,
		h.GetUserVoucherFromCode,
	)

	// PUBLIC > GIFT
	router.GET(
		"/pub/gifts/",
		middlewares.AuthorizePublic,
		h.GetGifts,
	)
	router.GET(
		"/pub/gifts/:giftID",
		middlewares.AuthorizePublic,
		h.GetGift,
	)

	router.NoRoute(h.HandleNotFound)

	return router
}
