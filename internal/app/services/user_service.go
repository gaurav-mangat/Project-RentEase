package services

import (
	"context"
	"fmt"
	"rentease/internal/domain/entities"
	"rentease/internal/domain/interfaces"
	"rentease/pkg/utils"
)

type UserService struct {
	userRepo interfaces.UserRepo
}

func NewUserService(userRepo interfaces.UserRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

//func (us *UserService) Login(user entities.User) error {
//}

func (us *UserService) SignUp(user entities.User) bool {
	// chota mota logic no db involved
	err := us.userRepo.SaveUser(user)
	if err != nil {
		fmt.Println(err)
	}
	return true

}

func (us *UserService) Findbyuname(username string) (entities.User, error) {
	user, err := us.userRepo.FindByUsername(context.Background(), username)
	if err != nil {
		return entities.User{}, err
	}
	if user == nil {
		return entities.User{}, nil // No user found
	}
	return *user, nil
}

func (us *UserService) Login(username, password string) (bool, error) {
	exist, err := us.userRepo.CheckPassword(context.Background(), username, password)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, nil
	}
	utils.ActiveUser = username
	return true, nil

}

//func (us *UserService) getUser(username string) (entities.User, error) {}
