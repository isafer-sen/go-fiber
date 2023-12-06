package repository

import (
	"app/core/database"
	"app/core/model"
)

type DeviceRepository struct {
}

// FindAll 查找所有设备
func (d *DeviceRepository) FindAll() (userAll []model.Devices, err error) {
	return userAll, database.DB.Find(&userAll).Error
}
