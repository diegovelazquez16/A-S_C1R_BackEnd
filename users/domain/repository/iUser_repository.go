package repository

import (
	"gorm.io/gorm"
	"holamundo/users/domain/models"
)


type IUserRepository interface {
	Create(user *models.User) error
	GetAll() ([]models.User, error)
	GetByID(id uint) (*models.User, error)
	Update(user *models.User) error// Interfaz para conectar el modelo con el repositorio
	Delete(id uint) error
}

type UserRepositoryImpl struct {
	DB *gorm.DB  // Implementación de la interfaz

}

func (r *UserRepositoryImpl) Create(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepositoryImpl) GetAll() ([]models.User, error) {
	var users []models.User
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *UserRepositoryImpl) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, id).Error
	return &user, err
}

func (r *UserRepositoryImpl) Update(user *models.User) error{
	return r.DB.Save(user).Error
}

func (r *UserRepositoryImpl) Delete(id uint) error {
	return r.DB.Delete(&models.User{}, id).Error
}
