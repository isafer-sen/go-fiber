package repository

import (
	"app/core/model"
	"app/db"
)

type UserRepository struct {
}

// FindAll 查找所有用
func (u *UserRepository) FindAll(page int, pageSize int) (userAll []model.User, err error) {
	return userAll, db.DB.Offset(page).Limit(pageSize).Find(&userAll).Error
}
