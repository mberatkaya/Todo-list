package user

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	Repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, nickname, fullName, password string) (*User, error) {
	user := &User{
		Nickname: nickname,
		FullName: fullName,
		Password: password,
	}
	return s.Repo.CreateUser(ctx, user)
}

func (s *UserService) GetUserByID(ctx context.Context, id primitive.ObjectID) (*User, error) {
	return s.Repo.GetUserByID(ctx, id)
}

func (s *UserService) GetUserByNickname(ctx context.Context, nickname string) (*User, error) {
	return s.Repo.GetUserByNickname(ctx, nickname)
}

func (s *UserService) UpdateUser(ctx context.Context, id primitive.ObjectID, nickname, fullName string) error {
	updateFields := bson.D{
		{"nickname", nickname},
		{"fullName", fullName},
	}
	return s.Repo.UpdateUser(ctx, id, updateFields)
}

func (s *UserService) DeleteUser(ctx context.Context, id primitive.ObjectID) error {
	return s.Repo.DeleteUser(ctx, id)
}
