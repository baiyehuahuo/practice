package model

type User struct {
	Id              int64  `json:"id,omitempty"`               // 用户id
	Name            string `json:"name,omitempty"`             // 用户名称
	FollowCount     int64  `json:"follow_count,omitempty"`     // 关注总数
	FollowerCount   int64  `json:"follower_count,omitempty"`   // 粉丝总数
	IsFollow        bool   `json:"is_follow,omitempty"`        // true-已关注，false-未关注
	Avatar          string `json:"avatar,omitempty"`           //用户头像
	BackgroundImage string `json:"background_image,omitempty"` //用户个人页顶部大图
	Signature       string `json:"signature,omitempty"`        //个人简介
	TotalFavorited  int64  `json:"total_favorited,omitempty"`  //获赞数量
	WorkCount       int64  `json:"work_count,omitempty"`       //作品数量
	FavoriteCount   int64  `json:"favorite_count,omitempty"`   //点赞数量
}
