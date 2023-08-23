package repository

import (
	"context"
	"errors"
	"fmt"
	"golang-microservice/config"
	"golang-microservice/constants"
	"golang-microservice/db"
	"golang-microservice/model"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type IUserRepository interface {
	Save(user *model.User) (string, error)
	GetById(id string) (*model.User, error)
	GetByEmail(email string) (user *model.User, err error)
	GetAll() (users []*model.User, err error)
	UpdateUser(user *model.User) error
	DeleteUserById(id string) error
}

type userRepository struct {
	Client *mongo.Client
	Cfg    *config.Config
}

func NewUserRepository(Client *mongo.Client, Cfg *config.Config) IUserRepository {
	return &userRepository{
		Client: Client,
		Cfg:    Cfg,
	}
}

func (repo *userRepository) getCollectionUser() *mongo.Collection {
	db := db.DBConn(repo.Cfg)
	userCollection := db.Database(repo.Cfg.DB.DBName).Collection(constants.Collection_USER)

	return userCollection
}

func (repo *userRepository) Save(user *model.User) (string, error) {
	insertResult, err := repo.getCollectionUser().InsertOne(context.Background(), &user)
	if err != nil {
		return "", err
	}
	idInserted := fmt.Sprintf("%v", insertResult.InsertedID)

	return idInserted, nil

}

func (repo *userRepository) GetById(id string) (user *model.User, err error) {

	filter := bson.ObjectIdHex(id)

	if err := repo.getCollectionUser().FindOne(context.Background(), filter).Decode(&user); err != nil {
		return &model.User{}, err
	}

	return user, nil
}

func (repo *userRepository) GetByEmail(email string) (user *model.User, err error) {
	filter := bson.M{
		"email": email,
	}
	if err := repo.getCollectionUser().FindOne(context.Background(), filter).Decode(&user); err != nil {
		return &model.User{}, err
	}

	return user, nil
}

func (repo *userRepository) GetAll() (users []*model.User, err error) {
	var user *model.User
	cursor, err := repo.getCollectionUser().Find(context.Background(), bson.D{})
	if err != nil {
		defer cursor.Close(context.Background())
		return nil, err
	}
	for cursor.Next(context.Background()) {
		if err := cursor.Decode(&user); err != nil {
			return users, errors.New("user not found")
		}
		users = append(users, user)
	}
	return users, err
}

func (repo *userRepository) UpdateUser(user *model.User) error {

	_, err := repo.getCollectionUser().UpdateOne(context.Background(), user.Id, user)
	return err
}

func (repo *userRepository) DeleteUserById(id string) error {
	_, err := repo.getCollectionUser().DeleteOne(context.Background(), bson.ObjectIdHex(id))
	if err != nil {
		return err
	}
	return nil
}
