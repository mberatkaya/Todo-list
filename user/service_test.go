package user

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	mockUser := &User{
		Nickname: "testuser",
		FullName: "Test User",
		Password: "password",
	}
	mockRepo.On("GetUserByNickname", mock.Anything, "testuser").Return(nil, errors.New("not found"))
	mockRepo.On("CreateUser", mock.Anything, mockUser).Return(mockUser, nil)

	createdUser, err := service.CreateUser(context.TODO(), "testuser", "Test User", "password")

	assert.NoError(t, err)
	assert.Equal(t, "testuser", createdUser.Nickname)
	mockRepo.AssertExpectations(t)
}

func TestGetUserByID(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	objectID := primitive.NewObjectID()
	mockUser := &User{
		ID:       objectID,
		Nickname: "testuser",
		FullName: "Test User",
	}

	mockRepo.On("GetUserByID", mock.Anything, objectID).Return(mockUser, nil)

	user, err := service.GetUserByID(context.TODO(), objectID)

	assert.NoError(t, err)
	assert.Equal(t, objectID, user.ID)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	objectID := primitive.NewObjectID()
	updateFields := bson.D{{"fullName", "Updated User"}}
	mockRepo.On("GetUserByID", mock.Anything, objectID).Return(&User{
		ID:       objectID,
		Nickname: "testuser",
		FullName: "Test User",
	}, nil)
	mockRepo.On("UpdateUser", mock.Anything, objectID, updateFields).Return(nil)

	err := service.UpdateUser(context.TODO(), objectID, updateFields)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	objectID := primitive.NewObjectID()
	mockRepo.On("DeleteUser", mock.Anything, objectID).Return(nil)

	err := service.DeleteUser(context.TODO(), objectID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestValidatePassword(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	mockUser := &User{
		Nickname: "testuser",
		FullName: "Test User",
		Password: "$2a$10$7ZyTBOx8hR1PZ8mnCwKJpO.Z4tA9UzHgQFll3oeIEihh7gweu9EmS",
	}

	mockRepo.On("GetUserByNickname", mock.Anything, "testuser").Return(mockUser, nil)

	user, err := service.ValidatePassword(context.TODO(), "testuser", "password")

	assert.NoError(t, err)
	assert.Equal(t, "testuser", user.Nickname)
	mockRepo.AssertExpectations(t)
}
