package models

import (
	"ginchat/utils"
	"gorm.io/gorm"
	"log"
)

type Contact struct {
	gorm.Model
	OwnerID  uint // who's contact
	TargetID uint // target user id
	Type     int
	Desc     string `gorm:"type:varchar(255);"`
}

func (c *Contact) TableName() string {
	return "contact"
}

func SearchFriends(userID uint) []*UserBasic {
	var contacts []*Contact
	utils.GetDB().Where("owner_id = ? and type = 1", userID).Find(&contacts)
	friendIDs := make([]int, len(contacts))
	for i, con := range contacts {
		friendIDs[i] = int(con.TargetID)
	}
	return FindUserByIDs(friendIDs)
}

func AddFriend(userID uint, targetID uint) bool {
	user := FindUserByID(int(userID))
	target := FindUserByID(int(targetID))
	if user.ID == 0 || target.ID == 0 {
		log.Println("user ID or target ID is none:", user.ID, target.ID)
		return false
	}
	db := utils.GetDB()
	err := db.Transaction(func(tx *gorm.DB) error {
		con := Contact{}
		con.OwnerID = user.ID
		con.TargetID = target.ID
		con.Type = 1
		if err := tx.Create(&con).Error; err != nil {
			return err
		}
		con = Contact{}
		con.OwnerID = target.ID
		con.TargetID = user.ID
		con.Type = 1
		if err := tx.Create(&con).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}

func AutoMigrateContact() error {
	return utils.GetDB().AutoMigrate(&Contact{})
}
