package services

import (
	"context"

	"example.com/vote-system-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(usercollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		usercollection: usercollection,
		ctx:            ctx,
	}
}

func (u *UserServiceImpl) CreateUser(user *models.User) error {
	_, err := u.usercollection.InsertOne(u.ctx, user)
	return err
}

func (u *UserServiceImpl) GetUser(id *string) (*models.User, error) {
	var user *models.User
	query := bson.D{bson.E{Key: "_id", Value: id}}
	err := u.usercollection.FindOne(u.ctx, query).Decode(&user)
	return user, err
}

func (u *UserServiceImpl) GetAll() ([]*models.User, error) {
	return nil, nil
}

func (u *UserServiceImpl) UpdateUser(*models.User) error {
	return nil
}

func (u *UserServiceImpl) DeleteUser(name *string) error {
	return nil
}
