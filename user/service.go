package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	Repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, nickname, fullName string) (*User, error) {
	user := &User{
		Nickname: nickname,
		FullName: fullName,
	}
	return s.Repo.CreateUser(ctx, user)
}

func (s *UserService) GetUserByID(ctx context.Context, id primitive.ObjectID) (*User, error) {
	return s.Repo.GetUserByID(ctx, id)
}

func (s *UserService) DeleteUser(ctx context.Context, id primitive.ObjectID) error {
	return s.Repo.DeleteUser(ctx, id)
}
