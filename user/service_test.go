package user

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func TestCreateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	nickname := "testuser"
	fullName := "Test User"
	password := "password123"

	t.Run("Success", func(t *testing.T) {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		mockUser := &User{
			ID:       primitive.NewObjectID(),
			Nickname: nickname,
			FullName: fullName,
			Password: string(hashedPassword),
		}

		mockRepo.On("GetUserByNickname", mock.Anything, nickname).Return(nil, errors.New("user not found")).Once()

		mockRepo.On("CreateUser", mock.Anything, mock.AnythingOfType("*user.User")).Return(mockUser, nil).Once()

		user, err := service.CreateUser(context.TODO(), nickname, fullName, password)

		assert.NoError(t, err, "Error should be nil when user is created successfully")
		assert.NotNil(t, user, "User should not be nil")
		assert.Equal(t, nickname, user.Nickname, "Nicknames should match")
		assert.Equal(t, fullName, user.FullName, "Full names should match")

		mockRepo.AssertExpectations(t)
	})

	t.Run("NicknameAlreadyExists", func(t *testing.T) {
		// Expect GetUserByNickname to return an existing user
		existingUser := &User{Nickname: nickname}
		mockRepo.On("GetUserByNickname", mock.Anything, nickname).Return(existingUser, nil)

		user, err := service.CreateUser(context.TODO(), nickname, fullName, password)

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, "nickname already exists", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestGetUserByID(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	id := primitive.NewObjectID()
	mockUser := &User{
		ID:       id,
		Nickname: "testuser",
		FullName: "Test User",
	}

	// Expect GetUserByID to return the mockUser
	mockRepo.On("GetUserByID", mock.Anything, id).Return(mockUser, nil)

	user, err := service.GetUserByID(context.TODO(), id)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, id, user.ID)
	assert.Equal(t, "testuser", user.Nickname)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	id := primitive.NewObjectID()
	updateFields := bson.D{{"fullName", "Updated User"}}
	mockUser := &User{ID: id, Nickname: "testuser", FullName: "Old User"}

	// Expect GetUserByID to return the existing user
	mockRepo.On("GetUserByID", mock.Anything, id).Return(mockUser, nil)

	// Expect UpdateUser to be called successfully
	mockRepo.On("UpdateUser", mock.Anything, id, updateFields).Return(nil)

	err := service.UpdateUser(context.TODO(), id, updateFields)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	id := primitive.NewObjectID()

	// Expect DeleteUser to be called successfully
	mockRepo.On("DeleteUser", mock.Anything, id).Return(nil)

	err := service.DeleteUser(context.TODO(), id)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
