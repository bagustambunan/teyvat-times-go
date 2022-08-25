package repositories

import (
	"final-project-backend/httperror"
	"final-project-backend/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	MatchingCredential(email, password string) (*models.User, error)
	FindUser(user *models.User) (*models.User, error)
	FindUserByReferralCode(refCode string) (*models.User, error)
	CheckUsernameAndEmail(user *models.User) error
	Save(user *models.User) (*models.User, error)
	SaveUserReferral(userRef *models.UserReferral) error
	UpdateCoins(user *models.User, coins int) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

type URConfig struct {
	DB *gorm.DB
}

func NewUserRepository(c *URConfig) UserRepository {
	return &userRepository{db: c.DB}
}

func (repo *userRepository) MatchingCredential(email, password string) (*models.User, error) {
	var user models.User
	result := repo.db.
		Where("email = ?", email).
		First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	if notMatchErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); notMatchErr != nil {
		return nil, notMatchErr
	}
	return &user, nil
}

func (repo *userRepository) FindUser(user *models.User) (*models.User, error) {
	result := repo.db.
		First(&user)
	return user, result.Error
}

func (repo *userRepository) FindUserByReferralCode(refCode string) (*models.User, error) {
	user := &models.User{}
	result := repo.db.
		Where("referral_code = ?", refCode).
		First(&user)
	return user, result.Error
}

func (repo *userRepository) CheckUsernameAndEmail(user *models.User) error {
	resultUserName := repo.db.
		Where("username = ?", user.Username).
		First(&user)
	if resultUserName.Error == nil {
		return httperror.BadRequestError("DUPLICATE_USERNAME", "Username already taken")
	}

	resultEmail := repo.db.
		Where("email = ?", user.Email).
		First(&user)
	if resultEmail.Error == nil {
		return httperror.BadRequestError("DUPLICATE_EMAIL", "Email already used")
	}

	return nil
}

func (repo *userRepository) Save(user *models.User) (*models.User, error) {
	result := repo.db.
		Select("Email", "Name", "Username", "Phone", "ReferralCode", "Password", "Address", "AddressID").
		Create(user)
	return user, result.Error
}

func (repo *userRepository) SaveUserReferral(userRef *models.UserReferral) error {
	result := repo.db.
		Create(&userRef)
	return result.Error
}

func (repo *userRepository) UpdateCoins(user *models.User, coins int) (*models.User, error) {
	result := repo.db.
		Model(&user).
		UpdateColumn("coins", gorm.Expr("coins + ?", coins))
	return user, result.Error
}
