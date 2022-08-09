package service

import (
	"SimpleDY/dao"
	"SimpleDY/global"
	"SimpleDY/status"
)

type UserService struct {
}

//Register
/**
param:用户名 密码 string
response:注册结果，用户id,错误码
*/
func (userservice UserService) Register(username, password string) (bool, uint64, int) {
	var count int64
	global.Db.Model(&dao.User{}).Where("username = ?", username).Count(&count)
	if count > 0 {
		return false, 0, status.UsernameHasExistedError
	}
	user := dao.User{
		Name:     "",
		Username: username,
		Password: password,
	}
	if global.Db.Create(&user).RowsAffected == 1 {
		return true, user.Id, 0
	}
	return false, 0, status.UnknownError
}

//Login
/**
param 用户名 密码
response 登陆用户id 错误码
*/
func (userservice UserService) Login(username, password string) (uint64, uint64) {
	var user dao.User
	var count int64
	global.Db.Model(&dao.User{}).Where("username = ?", username).Find(&user).Count(&count)
	if count == 0 {
		return 0, status.UserNotExistOrPasswordWrongError
	}

	if user.Password != password {
		return 0, status.UserNotExistOrPasswordWrongError
	}
	return user.Id, status.Success
}

//GetInfoByUserId
/*
param 用户id
response 用户结构体
*/
func (userservice UserService) GetInfoByUserId(userid uint64) *dao.User {
	var user dao.User
	global.Db.Model(&dao.User{}).Where("id = ?", userid).First(&user)
	return &user
}
