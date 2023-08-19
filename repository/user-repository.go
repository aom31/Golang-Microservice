package repository

import (
	"golang-microservice/constant"
	"golang-microservice/db"

	"gopkg.in/mgo.v2"
)

type IUserRepository interface {
	Save() error
}

type userRepository struct {
	c *mgo.Collection
}

func NewUserRepository(conn db.Connection) IUserRepository {
	return &userRepository{
		c: conn.DB().C(constant.Collection_USER),
	}
}

func (repo *userRepository) Save() error {

	return nil
}
