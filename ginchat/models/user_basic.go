package models

import (
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name          string    `gorm:"column:name" json:"name"`
	Password      string    `gorm:"column:password" json:"password"`
	Phone         string    `gorm:"column:phone" json:"phone"`
	Email         string    `gorm:"column:email" json:"email"`
	Identity      string    `gorm:"column:identity" json:"identity"`
	ClientIP      string    `gorm:"column:client_ip" json:"client_ip"`
	ClientPort    string    `gorm:"column:client_port" json:"client_port"`
	LoginTime     time.Time `gorm:"column:login_time" json:"login_time"`
	HeartBeatTime time.Time `gorm:"column:heart_beat_time" json:"heart_beat_time"`
	LoginOutTime  time.Time `gorm:"column:login_out_time" json:"login_out_time"`
	IsLoginOut    bool      `gorm:"column:is_login_out" json:"is_login_out"`
	DeviceInfo    string    `gorm:"column:device_info" json:"device_info"`
}

func (us *UserBasic) TableName() string {
	return "user_basic"
}
