package models

import (
	"fmt"
	"ginchat/utils"
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name          string    `gorm:"column:name" json:"name"`
	Password      string    `gorm:"column:password" json:"password"`
	Salt          string    `gorm:"column:salt" json:"salt"`
	Phone         string    `gorm:"column:phone" json:"phone" valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string    `gorm:"column:email" json:"email" valid:"email"`
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

func GetUserList() []*UserBasic {
	var datas []*UserBasic
	utils.GetDB().Find(&datas)
	return datas
}

func CreateUser(user UserBasic) *gorm.DB {
	return utils.GetDB().Create(&user)
}

func DeleteUser(user UserBasic) *gorm.DB {
	return utils.GetDB().Delete(&user)
}

func UpdateUser(user UserBasic) *gorm.DB {
	return utils.GetDB().Updates(&user)
}

func FindUserByID(ID int) *UserBasic {
	var user UserBasic
	utils.GetDB().Where("id = ?", ID).First(&user)
	return &user
}

func FindUserByIDs(IDs []int) []*UserBasic {
	var users []*UserBasic
	utils.GetDB().Where("id in ?", IDs).Find(&users)
	return users
}

func FindUserByName(name string) *UserBasic {
	var user UserBasic
	utils.GetDB().Where("name = ?", name).First(&user)
	return &user
}

func FindUserByPhone(phone string) *UserBasic {
	var user UserBasic
	utils.GetDB().Where("phone = ?", phone).First(&user)
	return &user
}

func FindUserByEmail(email string) *UserBasic {
	var user UserBasic
	utils.GetDB().Where("email = ?", email).First(&user)
	return &user
}

func FindUserByNameAndPwd(name, password string) *UserBasic {
	user := UserBasic{}
	utils.GetDB().Where("name = ? and password = ?", name, password).First(&user)
	token := utils.Md5EncodeSmall(fmt.Sprintf("%d", time.Now().Unix()))
	utils.GetDB().Model(&user).Where("id = ?", user.ID).Update("identity", token)
	return &user
}

func AutoMigrateUserBasic() error {
	return utils.GetDB().AutoMigrate(&UserBasic{})
}
