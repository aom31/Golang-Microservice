package model

import (
	"golang-microservice/pkg/proto-pb/pb"
	"time"

	"gopkg.in/mgo.v2/bson"
)

func (u *User) ToProtoBuffer() *pb.User {
	return &pb.User{
		Id:       u.Id.Hex(),
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		Created:  u.CreatedTime.Unix(),
		Updated:  u.UpdatedTime.Unix(),
	}
}

func (u *User) FromProtoBuffer(user *pb.User) {
	u.Id = bson.ObjectIdHex(user.GetId())
	u.Name = user.GetName()
	u.Email = user.GetEmail()
	u.Password = user.GetPassword()
	u.CreatedTime = time.Unix(user.Created, 0)
	u.UpdatedTime = time.Unix(user.Updated, 0)
}
