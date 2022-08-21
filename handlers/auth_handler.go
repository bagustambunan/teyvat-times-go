package handlers

import (
	"final-project-backend/dto"
	"final-project-backend/helpers"
	"final-project-backend/httperror"
	"final-project-backend/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

func (h *Handler) SignUp(c *gin.Context) {
	payload, _ := c.Get("payload")
	signUpReq, _ := payload.(*dto.SignUpReq)

	userRef := &models.UserReferral{}
	if signUpReq.ReferrerCode != "" {
		referrerUser, errRefCode := h.authService.GetUserByReferralCode(signUpReq.ReferrerCode)
		if errRefCode != nil {
			_ = c.Error(httperror.BadRequestError("Referral code is not valid", "INVALID_REFERRAL_CODE"))
			return
		}
		userRef.ReferrerUserID = referrerUser.ID
	}

	address := &models.Address{
		Street:     signUpReq.Street,
		City:       signUpReq.City,
		State:      signUpReq.State,
		Country:    signUpReq.Country,
		PostalCode: signUpReq.PostalCode,
	}
	hashPw, _ := bcrypt.GenerateFromPassword([]byte(signUpReq.Password), bcrypt.DefaultCost)
	hashPwStr := string(hashPw)
	user := &models.User{
		Email:    signUpReq.Email,
		Name:     signUpReq.Name,
		Phone:    signUpReq.Phone,
		Address:  address,
		Password: hashPwStr,
	}

	insertedUser, saveUserErr := h.authService.AddUser(user)
	if saveUserErr != nil {
		_ = c.Error(saveUserErr)
		return
	}
	userRef.UserID = user.ID

	userRefErr := h.authService.AddUserReferral(userRef)
	if userRefErr != nil {
		_ = c.Error(userRefErr)
		return
	}

	helpers.StandardResponse(c, http.StatusCreated, insertedUser)
}

func (h *Handler) SignIn(c *gin.Context) {
	payload, _ := c.Get("payload")
	signInReq, _ := payload.(*dto.SignInReq)
	token, err := h.authService.SignIn(signInReq)
	if err != nil {
		_ = c.Error(err)
		return
	}
	helpers.StandardResponse(c, http.StatusOK, *token)
}

func (h *Handler) GetUserFromToken(c *gin.Context) *models.User {
	userPayload, _ := c.Get("user")
	user, isUser := userPayload.(*models.User)
	if !isUser {
		_ = c.Error(httperror.UnauthorizedError())
		return nil
	}
	return user
}

func (h *Handler) GetUserInfo(c *gin.Context) {
	userID, idErr := strconv.Atoi(c.Param("userID"))
	if idErr != nil {
		_ = c.Error(idErr)
		return
	}

	if userID != h.GetUserFromToken(c).ID {
		_ = c.Error(httperror.UnauthorizedError())
		return
	}

	userRes, findUserErr := h.authService.GetUser(&models.User{
		ID: userID,
	})
	if findUserErr != nil {
		_ = c.Error(findUserErr)
		return
	}
	helpers.StandardResponse(c, http.StatusOK, userRes)
}
