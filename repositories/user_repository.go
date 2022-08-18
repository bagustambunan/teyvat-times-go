package repositories

import (
	"final-project-backend/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	MatchingCredential(email, password string) (*models.User, error)
	FindUser(user *models.User) (*models.User, error)
	FindUserByReferralCode(refCode string) (*models.User, error)
	Save(user *models.User) (*models.User, int, error)
	SaveUserReferral(userRef *models.UserReferral) error
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

func (repo *userRepository) Save(user *models.User) (*models.User, int, error) {
	result := repo.db.
		Select("Email", "Name", "Phone", "ReferralCode", "Password", "Address", "AddressID").
		Clauses(clause.OnConflict{DoNothing: true}).
		Create(user)
	return user, int(result.RowsAffected), result.Error
}

func (repo *userRepository) SaveUserReferral(userRef *models.UserReferral) error {
	result := repo.db.
		Create(&userRef)
	return result.Error
}
