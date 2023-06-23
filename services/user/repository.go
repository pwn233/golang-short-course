package user

import (
	"errors"

	db "github.com/pwn233/golang-short-course/config"
	"gorm.io/gorm"
)

type UserRepository struct{}

func (r *UserRepository) NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// AddUser is a function that add user from request body to database.
func (r *UserRepository) AddUser(request AddUserRequest) (err error) {
	if db.DevSandbox.DB == nil {
		return errors.New("database connection unavailable")
	}
	db := db.DevSandbox.DB
	err = db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Table("organization").Create(&request).Error; err != nil {
			return err
		} else {
			return nil
		}
	})
	if err != nil {
		return err
	} else {
		return nil
	}
}
