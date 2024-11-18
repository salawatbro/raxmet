package repository

import (
	"github.com/google/uuid"
	"github.com/salawatbro/raxmet/database"
	"github.com/salawatbro/raxmet/internal/interfaces"
	"github.com/salawatbro/raxmet/internal/models"
	"github.com/salawatbro/raxmet/pkg/response"
	"math"
)

type UserRepository struct {
}

func NewUserRepository() interfaces.UserRepositoryInterface {
	return &UserRepository{}
}

func (repo *UserRepository) Create(user models.User) (models.User, error) {
	if err := database.DB.Model(&models.User{}).Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (repo *UserRepository) FindAll(paginate response.Pagination) (*response.Pagination, error) {
	var users []models.User
	var totalRows int64
	query := database.DB.Model(&models.User{})

	query.Count(&totalRows)
	paginate.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(paginate.GetLimit())))
	paginate.TotalPages = totalPages
	if err := query.Offset(paginate.GetOffset()).Limit(paginate.GetLimit()).Order(paginate.GetSort()).Find(&users).Error; err != nil {
		return nil, err
	}
	return &paginate, nil
}

func (repo *UserRepository) ExistsByEmail(email string) bool {
	var user models.User
	if err := database.DB.Model(&models.User{}).Where("email = ?", email).First(&user).Error; err != nil {
		return false
	}
	return true
}

func (repo *UserRepository) FindByEmail(email string) (models.User, error) {
	var user models.User
	if err := database.DB.Model(&models.User{}).Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (repo *UserRepository) FindByID(id uuid.UUID) (models.User, error) {
	var user models.User
	if err := database.DB.Model(&models.User{}).Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (repo *UserRepository) Update(id uuid.UUID, user models.User) (models.User, error) {
	if err := database.DB.Model(&models.User{}).Where("id = ?", id).Updates(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (repo *UserRepository) UpdatePassword(id uuid.UUID, password string) error {
	if err := database.DB.Model(&models.User{}).Where("id = ?", id).Update("password", password).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) Delete(id uuid.UUID) error {
	if err := database.DB.Model(&models.User{}).Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		return err
	}
	return nil
}
