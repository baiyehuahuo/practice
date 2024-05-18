package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Name          string
	Password      string
	Phone         string
	Email         string
	Identity      string
	ClientIP      string
	ClientPort    string
	LoginTime     uint64
	HeartBeatTime uint64
	LoginOutTime  uint64
	IsLoginOut    bool
	DeviceInfo    string
}

func (us *UserBasic) TableName() string {
	return "user_basic"
}
