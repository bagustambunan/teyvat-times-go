package server

import (
	"final-project-backend/dto"
	"final-project-backend/handlers"
	"final-project-backend/middlewares"
	"final-project-backend/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
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

	corsConfig := cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Type", "X-Auth-Token", "Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Authorization"},
		MaxAge:          12 * time.Hour,
	})
	router.Use(corsConfig)

	h := handlers.New(&handlers.HandlerConfig{
		AuthService:         conf.AuthService,
		UserService:         conf.UserService,
		PostService:         conf.PostService,
		SubscriptionService: conf.SubscriptionService,
		TransactionService:  conf.TransactionService,
		VoucherService:      conf.VoucherService,
		GiftService:         conf.GiftService,
	})

	router.Static("docs", "swaggerui")
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
		"/posts",
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
		"/transactions",
		middlewares.AuthorizeInternal,
		h.GetTransactions,
	)
	router.GET(
		"/transaction-statuses",
		middlewares.AuthorizeInternal,
		h.GetTransactionStatuses,
	)
	router.POST(
		"/transactions/:transactionID/approve",
		middlewares.AuthorizeInternal,
		h.ApproveTransaction,
	)
	router.POST(
		"/transactions/:transactionID/reject",
		middlewares.AuthorizeInternal,
		h.RejectTransaction,
	)

	// ADMIN > VOUCHER
	router.GET(
		"/vouchers",
		middlewares.AuthorizeInternal,
		h.GetVouchers,
	)

	// ADMIN GIFT
	router.GET(
		"/gift-claim-statuses",
		middlewares.AuthorizeInternal,
		h.GetGiftClaimStatuses,
	)
	router.GET(
		"/gift-claims",
		middlewares.AuthorizeInternal,
		h.GetGiftClaims,
	)
	router.POST(
		"/gift-claims/:gcID/deliver",
		middlewares.AuthorizeInternal,
		h.DeliverGiftClaim,
	)
	router.POST(
		"/gift-claims/:gcID/reject",
		middlewares.AuthorizeInternal,
		h.RejectGiftClaim,
	)

	// PUBLIC > USER
	router.GET(
		"/pub/users/:userID",
		middlewares.AuthorizePublic,
		h.GetUser,
	)
	router.GET(
		"/pub/users/:userID/down-lines",
		middlewares.AuthorizePublic,
		h.GetUserDownLines,
	)

	// PUBLIC > POST
	router.GET(
		"/pub/tiers",
		middlewares.AuthorizePublic,
		h.GetTiers,
	)
	router.GET(
		"/pub/categories",
		middlewares.AuthorizePublic,
		h.GetCategories,
	)
	router.GET(
		"/pub/posts",
		middlewares.AuthorizePublic,
		h.GetPosts,
	)
	router.GET(
		"/pub/reading-history",
		middlewares.AuthorizePublic,
		h.PubGetReadingHistory,
	)
	router.GET(
		"/pub/posts/overview/:slug",
		middlewares.AuthorizePublic,
		h.PubGetOverviewPost,
	)
	router.GET(
		"/pub/posts/read/:slug",
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
	router.GET(
		"/pub/trending-posts",
		middlewares.AuthorizePublic,
		h.PubGetTrending,
	)

	// PUBLIC > SUBSCRIPTION
	router.GET(
		"/pub/subscriptions",
		middlewares.AuthorizePublic,
		h.GetSubscriptions,
	)
	router.GET(
		"/pub/subscriptions/:subscriptionID",
		middlewares.AuthorizePublic,
		h.GetSubscription,
	)
	router.GET(
		"/pub/user-subscriptions",
		middlewares.AuthorizePublic,
		h.GetUserSubscriptions,
	)
	router.GET(
		"/pub/user-subscriptions/active",
		middlewares.AuthorizePublic,
		h.GetUserActiveSubscription,
	)

	// PUBLIC > TRANSACTION
	router.POST(
		"/pub/transactions",
		middlewares.AuthorizePublic,
		middlewares.RequestValidator(&dto.TransactionReq{}),
		h.AddTransaction,
	)
	router.GET(
		"/pub/transactions",
		middlewares.AuthorizePublic,
		h.GetUserTransactions,
	)
	router.GET(
		"/pub/transactions/:transactionID",
		middlewares.AuthorizePublic,
		h.GetTransactionDetail,
	)
	router.GET(
		"/pub/user-spending",
		middlewares.AuthorizePublic,
		h.GetUserSpending,
	)

	// PUBLIC > VOUCHER
	router.GET(
		"/pub/user-vouchers",
		middlewares.AuthorizePublic,
		h.GetUserVouchers,
	)
	router.GET(
		"/pub/user-vouchers/:code",
		middlewares.AuthorizePublic,
		h.GetUserVoucherFromCode,
	)

	// PUBLIC > GIFT
	router.GET(
		"/pub/gifts",
		middlewares.AuthorizePublic,
		h.GetGifts,
	)
	router.GET(
		"/pub/gifts/:giftID",
		middlewares.AuthorizePublic,
		h.GetGift,
	)
	router.GET(
		"/pub/unclaimed-user-gifts",
		middlewares.AuthorizePublic,
		h.GetUnclaimedUserGifts,
	)
	router.POST(
		"/pub/gift-claims",
		middlewares.AuthorizePublic,
		h.SaveGiftClaim,
	)
	router.GET(
		"/pub/gift-claims",
		middlewares.AuthorizePublic,
		h.GetUserGiftClaims,
	)
	router.GET(
		"/pub/gift-claims/:gcID",
		middlewares.AuthorizePublic,
		h.GetGiftClaim,
	)
	router.POST(
		"/pub/gift-claims/:gcID/complete",
		middlewares.AuthorizePublic,
		h.PubCompleteGiftClaim,
	)

	// OPEN > PAYMENT
	router.POST(
		"/open/payment",
		middlewares.RequestValidator(&dto.PaymentReq{}),
		h.ProcessPayment,
	)

	router.NoRoute(h.HandleNotFound)

	return router
}
