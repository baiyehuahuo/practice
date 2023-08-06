package model

type User struct {
	Id              int64  `gorm:"primaryKey"` // 用户id
	Name            string `gorm:"not null"`   // 用户名称
	Password        string `gorm:"not null"`   // 用户密码
	FollowCount     int64  ``                  // 关注总数
	FollowerCount   int64  ``                  // 粉丝总数
	IsFollow        bool   ``                  // true-已关注，false-未关注
	Avatar          string ``                  //用户头像
	BackgroundImage string ``                  //用户个人页顶部大图
	Signature       string ``                  //个人简介
	TotalFavorited  int64  ``                  //获赞数量
	WorkCount       int64  ``                  //作品数量
	FavoriteCount   int64  ``                  //点赞数量
}
