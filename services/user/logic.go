package user

import (
	"github.com/pwn233/golang-short-course/common"
)

type UserService struct {
	repo *UserRepository
}

func NewUserService() *UserService {
	repo := UserRepository{}
	return &UserService{
		repo: repo.NewUserRepository(),
	}
}

// AddUser is a logic that contain validation and calling repository for create and execute statement.
func (u *UserService) AddUser(request AddUserRequest) (err error) {
	// Validating input step
	if err = common.CheckEmptyInput(request.Id); err != nil {
		return err
	}
	if err = common.CheckEmptyInput(request.FirstName); err != nil {
		return err
	}
	if err = common.CheckEmptyInput(request.LastName); err != nil {
		return err
	}
	// Call repository to create and execute statement
	if err = u.repo.AddUser(request); err != nil {
		return err
	} else {
		return nil
	}
}
