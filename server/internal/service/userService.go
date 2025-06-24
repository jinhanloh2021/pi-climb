package service

import (
	"github.com/jinhanloh2021/beta-blocker/internal/repository"
)

type UserService interface {
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{userRepo: repo}
}
