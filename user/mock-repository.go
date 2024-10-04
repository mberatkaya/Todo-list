package user

import (
	"context"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(ctx context.Context, user *User) (*User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockUserRepository) GetUserByID(ctx context.Context, id primitive.ObjectID) (*User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockUserRepository) GetUserByNickname(ctx context.Context, nickname string) (*User, error) {
	args := m.Called(ctx, nickname)
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockUserRepository) UpdateUser(ctx context.Context, id primitive.ObjectID, updateFields bson.D) error {
	args := m.Called(ctx, id, updateFields)
	return args.Error(0)
}

func (m *MockUserRepository) DeleteUser(ctx context.Context, id primitive.ObjectID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
