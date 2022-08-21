package server

import (
	"final-project-backend/config"
	"final-project-backend/db"
	"final-project-backend/repositories"
	"final-project-backend/services"
	"fmt"
)

func Init() {
	userRepository := repositories.NewUserRepository(&repositories.URConfig{DB: db.Get()})
	authService := services.NewAuthService(&services.AuthSConfig{
		UserRepository: userRepository,
		AppConfig:      config.Config,
	})
	userService := services.NewUserService(&services.USConfig{
		UserRepository: userRepository,
	})

	postRepository := repositories.NewPostRepository(&repositories.PRConfig{DB: db.Get()})
	postService := services.NewPostService(&services.PSConfig{
		PostRepository: postRepository,
	})

	router := NewRouter(&RouterConfig{
		AuthService: authService,
		UserService: userService,
		PostService: postService,
	})

	err := router.Run(fmt.Sprintf(":%d", config.Config.AppPort))
	if err != nil {
		fmt.Println("server error: ", err)
	}
}
