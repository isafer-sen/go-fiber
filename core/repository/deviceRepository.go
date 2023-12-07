package repository

import (
	"app/core/model"
	"app/db"
)

type DeviceRepository struct {
}

// FindAll 查找所有设备
func (d *DeviceRepository) FindAll() (userAll []model.Devices, err error) {
	return userAll, db.DB.Find(&userAll).Error
}
