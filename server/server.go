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

	router := NewRouter(&RouterConfig{
		AuthService: authService,
	})

	err := router.Run(fmt.Sprintf(":%d", config.Config.AppPort))
	if err != nil {
		fmt.Println("server error: ", err)
	}
}
