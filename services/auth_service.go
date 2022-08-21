package services

import (
	"final-project-backend/config"
	"final-project-backend/dto"
	"final-project-backend/httperror"
	"final-project-backend/models"
	"final-project-backend/repositories"
	"github.com/golang-jwt/jwt/v4"
	"math/rand"
	"time"
)

type AuthService interface {
	SignIn(*dto.SignInReq) (*dto.TokenRes, error)
	GetUser(u *models.User) (*dto.GetUserRes, error)
	CheckReferrerCode(refCode string) error
	GetUserByReferralCode(refCode string) (*models.User, error)
	AddUser(u *models.User) (*dto.SignUpRes, error)
	AddUserReferral(userRef *models.UserReferral) error
}

type authService struct {
	userRepository repositories.UserRepository
	appConfig      config.AppConfig
}

type AuthSConfig struct {
	UserRepository repositories.UserRepository
	AppConfig      config.AppConfig
}

func NewAuthService(c *AuthSConfig) AuthService {
	return &authService{
		userRepository: c.UserRepository,
		appConfig:      c.AppConfig,
	}
}

type idTokenClaims struct {
	jwt.RegisteredClaims
	User *models.User `json:"user"`
}

func (serv *authService) generateReferralCode(size int) string {
	alpha := "WXYZGHJKLQRSTABCDEFMNPUV"
	buf := make([]byte, size)
	for i := 0; i < size; i++ {
		buf[i] = alpha[rand.Intn(len(alpha))]
	}
	return string(buf)
}

func (serv *authService) generateJWTToken(user *models.User) (*dto.TokenRes, error) {
	claims := &idTokenClaims{
		jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Minute * time.Duration(serv.appConfig.JWTExpiryInMinutes))),
			Issuer:    serv.appConfig.AppName,
		},
		user,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(serv.appConfig.JWTSecretKey)
	if err != nil {
		return nil, err
	}

	dtoToken := dto.TokenRes{
		UserID: user.ID,
		Token:  tokenString,
	}
	return &dtoToken, nil
}

func (serv *authService) SignIn(req *dto.SignInReq) (*dto.TokenRes, error) {
	user, noAuthErr := serv.userRepository.MatchingCredential(req.Email, req.Password)
	if noAuthErr != nil || user == nil {
		return nil, httperror.UnauthorizedError()
	}
	token, err := serv.generateJWTToken(user)
	return token, err
}

func (serv *authService) GetUser(u *models.User) (*dto.GetUserRes, error) {
	user, err := serv.userRepository.FindUser(u)
	if err != nil {
		return nil, err
	}
	return new(dto.GetUserRes).FromUser(user), nil
}

func (serv *authService) CheckReferrerCode(refCode string) error {
	_, err := serv.userRepository.FindUserByReferralCode(refCode)
	return err
}

func (serv *authService) GetUserByReferralCode(refCode string) (*models.User, error) {
	user, err := serv.userRepository.FindUserByReferralCode(refCode)
	return user, err
}

func (serv *authService) AddUser(u *models.User) (*dto.SignUpRes, error) {
	for {
		newRefCode := serv.generateReferralCode(6)
		if _, err := serv.userRepository.FindUserByReferralCode(newRefCode); err != nil {
			u.ReferralCode = newRefCode
			break
		}
	}

	user, rowsAffected, err := serv.userRepository.Save(u)
	if err == nil && rowsAffected == 0 {
		return nil, httperror.BadRequestError("Duplicate email", "DUPLICATE_EMAIL")
	}
	return new(dto.SignUpRes).FromUser(user), nil
}

func (serv *authService) AddUserReferral(userRef *models.UserReferral) error {
	err := serv.userRepository.SaveUserReferral(userRef)
	return err
}
