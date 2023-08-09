package entity

import "douyin/pb"

// User in mysql
// gorm generated by https://www.printlove.cn/tools/sql2gorm (convert mysql table to gorm tag)
type User struct {
	ID              int64  `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT;comment:用户 ID" json:"id"`       // 用户id
	Name            string `gorm:"column:name;type:varchar(32);unique;comment:用户昵称;NOT NULL" json:"name"`                    // 用户名称
	Password        string `gorm:"column:password;type:varchar(32);comment:用户密码;NOT NULL" json:"password"`                   // 用户密码
	FollowCount     int64  `gorm:"column:follow_count;type:int(11) unsigned;default:0;comment:关注总数" json:"follow_count"`     // 关注总数
	FollowerCount   int64  `gorm:"column:follower_count;type:int(11) unsigned;default:0;comment:粉丝总数" json:"follower_count"` // 粉丝总数
	IsFollow        bool   `gorm:"column:is_follow;type:tinyint(1);default:0;comment:true-已关注, false-未关注" json:"is_follow"`  // true-已关注，false-未关注
	Avatar          string `gorm:"column:avatar;type:varchar(64);comment:用户头像" json:"avatar"`                                //用户头像
	BackgroundImage string `gorm:"column:background_image;type:varchar(64);comment:用户个人顶部大图" json:"background_image"`        //用户个人页顶部大图
	Signature       string `gorm:"column:signature;type:varchar(64);comment:用户简介" json:"signature"`                          //个人简介
	TotalFavorited  int64  `gorm:"column:total_favorited;type:int(11);default:0;comment:获赞数量" json:"total_favorited"`        //获赞数量
	WorkCount       int64  `gorm:"column:work_count;type:int(11);default:0;comment:作品数量" json:"work_count"`                  //作品数量
	FavoriteCount   int64  `gorm:"column:favorite_count;type:int(11);default:0;comment:点赞数量" json:"favorite_count"`          //点赞数量
}

func (User) TableName() string {
	return "Users"
}

func (user *User) GetPBUser() *pb.User {
	pbUser := &pb.User{
		Id:              new(int64),
		Name:            new(string),
		FollowCount:     new(int64),
		FollowerCount:   new(int64),
		IsFollow:        new(bool),
		Avatar:          new(string),
		BackgroundImage: new(string),
		Signature:       new(string),
		TotalFavorited:  new(int64),
		WorkCount:       new(int64),
		FavoriteCount:   new(int64),
	}
	*pbUser.Id = user.ID
	*pbUser.Name = user.Name
	*pbUser.FollowCount = user.FollowerCount
	*pbUser.FollowerCount = user.FollowerCount
	*pbUser.IsFollow = user.IsFollow
	*pbUser.Avatar = user.Avatar
	*pbUser.BackgroundImage = user.BackgroundImage
	*pbUser.Signature = user.Signature
	*pbUser.TotalFavorited = user.TotalFavorited
	*pbUser.WorkCount = user.WorkCount
	*pbUser.FavoriteCount = user.FavoriteCount
	return pbUser
}
