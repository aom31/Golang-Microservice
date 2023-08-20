package repository

import (
	"golang-microservice/constant"
	"golang-microservice/db"
	"golang-microservice/model"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type IUserRepository interface {
	Save(user *model.User) error
	GetById(id string) (*model.User, error)
	GetByEmail(email string) (user *model.User, err error)
	GetAll() (users []*model.User, err error)
	UpdateUser(user *model.User) error
	DeleteUserById(id string) error
}

type userRepository struct {
	mgcollec *mgo.Collection
}

func NewUserRepository(conn db.Connection) IUserRepository {
	return &userRepository{
		mgcollec: conn.DB().C(constant.Collection_USER),
	}
}

func (repo *userRepository) Save(user *model.User) error {
	return repo.mgcollec.Insert(&user)

}

func (repo *userRepository) GetById(id string) (user *model.User, err error) {
	err = repo.mgcollec.FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

func (repo *userRepository) GetByEmail(email string) (user *model.User, err error) {
	query := bson.M{
		"email": email,
	}
	err = repo.mgcollec.Find(query).One(&user)
	return user, err
}

func (repo *userRepository) GetAll() (users []*model.User, err error) {

	err = repo.mgcollec.Find(bson.M{}).One(&users)
	return users, err
}

func (repo *userRepository) UpdateUser(user *model.User) error {

	return repo.mgcollec.UpdateId(user.Id, user)
}

func (repo *userRepository) DeleteUserById(id string) error {
	return repo.mgcollec.RemoveId(bson.ObjectIdHex(id))
}
