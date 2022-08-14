package services

import (
	"context"
	"errors"

	helper "example.com/vote-system-api/helpers"
	"example.com/vote-system-api/models"
	uuid "github.com/satori/go.uuid"
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
	newId := uuid.NewV4()
	user.Id = newId
	_, err := u.usercollection.InsertOne(u.ctx, user)
	return err
}

func (u *UserServiceImpl) GetUser(id *uuid.UUID) (*models.User, error) {
	var user *models.User
	query := bson.D{bson.E{Key: "_id", Value: id}}
	err := u.usercollection.FindOne(u.ctx, query).Decode(&user)
	return user, err
}

func (u *UserServiceImpl) GetAll() ([]*models.User, error) {
	var users []*models.User

	cursor, err := u.usercollection.Find(u.ctx, bson.D{{}})

	if err != nil {
		return nil, err
	}

	for cursor.Next((u.ctx)) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(u.ctx)

	if len(users) == 0 {
		return nil, errors.New(("documents not found"))
	}

	return users, nil
}

func (u *UserServiceImpl) UpdateUser(user *models.User, id *uuid.UUID) error {
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "name", Value: user.Name}, bson.E{Key: "password", Value: user.Password}}}}

	result, _ := u.usercollection.UpdateOne(u.ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}

	return nil
}

func (u *UserServiceImpl) DeleteUser(id *uuid.UUID) error {
	query := bson.D{bson.E{Key: "_id", Value: id}}
	result, _ := u.usercollection.DeleteOne(u.ctx, query)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}

// validate bcrypt
// func (u *UserServiceImpl) ValidatePassword(password *string, encryptedPassword *string) (isValid bool, err error) {

// }

func (u *UserServiceImpl) Login(user models.RequestUser) (token string, ok bool) {
	var foundUser *models.User

	query := bson.D{bson.E{Key: "name", Value: user.Name}}
	err := u.usercollection.FindOne(u.ctx, query).Decode(&foundUser)

	if err != nil || user.Password != foundUser.Password {
		return "", false
	}

	token, err = helper.CreateToken(user.Name, user.Password)

	if err != nil {
		return "", false
	}

	return token, true
}
