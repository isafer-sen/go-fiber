package repository

import (
	"app/core/database"
	"app/core/model"
)

type UserRepository struct {
}

// FindAll 查找所有用
func (u *UserRepository) FindAll(page int, pageSize int) (userAll []model.User, err error) {
	return userAll, database.DB.Offset(page).Limit(pageSize).Find(&userAll).Error
}
