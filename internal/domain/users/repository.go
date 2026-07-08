package users

import (
	"errors"

	"gorm.io/gorm"
)

type IRegisterRepository interface {
	Register(user *User) error
}

type registerRepository struct {
	db *gorm.DB
}

func NewRegisterRepository(db *gorm.DB) IRegisterRepository {
	return &registerRepository{
		db: db,
	}
}


func (r *registerRepository) Register(user *User) error {
	var existing User
	var err error

	if err = r.db.Where("email = ?", user.Email).First(&existing).Error; err == nil {
		return errors.New("email already exists")
	}

	if err = r.db.Where("phone = ?", user.Phone).First(&existing).Error; err == nil {
		return errors.New("phone already exists")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if err := r.db.Create(user).Error; err != nil {
		return err
	}

	return nil
}