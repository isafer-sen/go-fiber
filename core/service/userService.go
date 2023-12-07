package service

import (
	"app/core/model"
	"app/core/repository"
	"app/helper/cache"
	"encoding/json"
	"time"
)

type UserService struct {
	userRepository repository.UserRepository
}

func (u *UserService) FindAll(page int, pageSize int) (userAll []model.User, err error) {
	return u.userRepository.FindAll(page, pageSize)
}

func (u *UserService) FindAllCache(page int, pageSize int, md5Query string) (userAll []model.User, err error) {
	if cache.HasCache(md5Query) {
		//log.Info("走缓存")
		result, _ := cache.Get(md5Query)
		var u []model.User
		err := json.Unmarshal([]byte(result), &u)
		return u, err
	}
	//log.Info("不走缓存")
	data, err := u.userRepository.FindAll(page, pageSize)
	jsonData, _ := json.Marshal(data)
	if err := cache.Set(md5Query, jsonData, time.Second*60); err != nil {
		return nil, err
	}
	return data, nil
}
