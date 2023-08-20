package repository

import (
	"fmt"
	"golang-microservice/db"
	"golang-microservice/model"
	"log"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Panicln(err)
	}

}

func TestUserSave(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err, "err connection db")
	defer conn.Close()

	id := bson.NewObjectId()

	user := &model.User{
		Id:          id,
		Name:        "TEST",
		Email:       fmt.Sprintf("%s@email.test", id.Hex()),
		Password:    "123456789",
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}

	r := NewUserRepository(conn)
	err = r.Save(user)
	assert.NoError(t, err)

	found, err := r.GetById(user.Id.Hex())
	assert.NoError(t, err)
	assert.NotNil(t, found)

}
