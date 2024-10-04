package user

import (
	"context"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(ctx context.Context, nickname, fullName, password string) (*User, error) {
	args := m.Called(ctx, nickname, fullName, password)
	if user, ok := args.Get(0).(*User); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserService) GetUserByID(ctx context.Context, id primitive.ObjectID) (*User, error) {
	args := m.Called(ctx, id)
	if user, ok := args.Get(0).(*User); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserService) GetUserByNickname(ctx context.Context, nickname string) (*User, error) {
	args := m.Called(ctx, nickname)
	if user, ok := args.Get(0).(*User); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserService) UpdateUser(ctx context.Context, id primitive.ObjectID, updateFields bson.D) error {
	args := m.Called(ctx, id, updateFields)
	return args.Error(0)
}

func (m *MockUserService) DeleteUser(ctx context.Context, id primitive.ObjectID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockUserService) ValidatePassword(ctx context.Context, nickname, password string) (*User, error) {
	args := m.Called(ctx, nickname, password)
	if user, ok := args.Get(0).(*User); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}
