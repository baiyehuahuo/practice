package model

type Course struct {
	Cid   int    `form:"cid" gorm:"primaryKey"`
	Cname string `form:"cname"`
	Tid   int    `form:"tid"`
}
