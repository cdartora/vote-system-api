package services

import (
	"example.com/vote-system-api/src/models"
	uuid "github.com/satori/go.uuid"
)

type UserService interface {
	CreateUser(*models.User) error
	GetUser(*uuid.UUID) (*models.User, error)
	GetAll() ([]*models.User, error)
	UpdateUser(*models.User, *uuid.UUID) error
	DeleteUser(*uuid.UUID) error
	Login(user models.RequestUser) (token string, ok bool)
}
