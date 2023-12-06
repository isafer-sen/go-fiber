package service

import (
	"app/core/model"
	"app/core/repository"
)

type DeviceService struct {
	deviceRepository repository.DeviceRepository
}

// FindAll 查找所有设备
func (d *DeviceService) FindAll() (userAll []model.Devices, err error) {
	return d.deviceRepository.FindAll()
}
