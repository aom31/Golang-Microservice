package mockdata

import (
	"golang-microservice/model"

	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock implementation of IUserRepository
type MockUserRepository struct {
	*mock.Mock
}

func (mock *MockUserRepository) Save(user *model.User) (string, error) {
	args := mock.Called(user)
	return args.String(0), args.Error(1)
}
