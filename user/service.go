package user

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(ctx context.Context, nickname, fullName, password string) (*User, error)
	GetUserByID(ctx context.Context, id primitive.ObjectID) (*User, error)
	GetUserByNickname(ctx context.Context, nickname string) (*User, error)
	UpdateUser(ctx context.Context, id primitive.ObjectID, updateFields bson.D) error
	DeleteUser(ctx context.Context, id primitive.ObjectID) error
	ValidatePassword(ctx context.Context, nickname, password string) (*User, error)
}

type userService struct {
	Repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{Repo: repo}
}

func (s *userService) CreateUser(ctx context.Context, nickname, fullName, password string) (*User, error) {
	_, err := s.Repo.GetUserByNickname(ctx, nickname)
	if err == nil {
		return nil, errors.New("nickname already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &User{
		Nickname: nickname,
		FullName: fullName,
		Password: string(hashedPassword),
	}

	return s.Repo.CreateUser(ctx, user)
}

func (s *userService) GetUserByID(ctx context.Context, id primitive.ObjectID) (*User, error) {
	return s.Repo.GetUserByID(ctx, id)
}

func (s *userService) GetUserByNickname(ctx context.Context, nickname string) (*User, error) {
	return s.Repo.GetUserByNickname(ctx, nickname)
}

func (s *userService) UpdateUser(ctx context.Context, id primitive.ObjectID, updateFields bson.D) error {
	existingUser, err := s.Repo.GetUserByID(ctx, id)
	if err != nil {
		return err
	}

	for _, field := range updateFields {
		if field.Key == "nickname" {
			if existingUser.Nickname != field.Value {
				_, err := s.Repo.GetUserByNickname(ctx, field.Value.(string))
				if err == nil {
					return errors.New("nickname already exists")
				}
			}
		}
	}

	return s.Repo.UpdateUser(ctx, id, updateFields)
}

func (s *userService) DeleteUser(ctx context.Context, id primitive.ObjectID) error {
	return s.Repo.DeleteUser(ctx, id)
}

func (s *userService) ValidatePassword(ctx context.Context, nickname, password string) (*User, error) {
	user, err := s.Repo.GetUserByNickname(ctx, nickname)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	return user, nil
}
