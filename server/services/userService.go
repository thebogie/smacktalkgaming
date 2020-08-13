package services

import (
	//"errors"

	"github.com/thebogie/smacktalkgaming/config"
	"github.com/thebogie/smacktalkgaming/repos"
	"github.com/thebogie/smacktalkgaming/types"
)

// UserService interface
type UserService interface {
	GetUserByObjectID(*types.User) bool
	GetUserByEmail(*types.User) bool
	AddUser(*types.User)
}

type userService struct {
	UserRepo repos.UserRepo
}

// NewUserService will instantiate User Service
func NewUserService(
	userRepo repos.UserRepo) UserService {

	return &userService{
		UserRepo: userRepo,
	}
}

func (us *userService) AddUser(in *types.User) {

	if !us.GetUserByEmail(in) {
		us.UserRepo.AddUser(in)
	} else {
		config.Apex.Infof("User already exists: %+v", in)
	}

	return
}

func (us *userService) GetUserByObjectID(in *types.User) bool {
	//if id == 0 {
	//	return nil, errors.New("id param is required")
	//}

	return us.UserRepo.FindUserByObjectID(in)
}

func (us *userService) GetUserByEmail(in *types.User) bool {

	return us.UserRepo.FindUserByEmail(in)
}
