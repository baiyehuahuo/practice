package model

type Course struct {
	Cid   int    `form:"cid"`
	Cname string `form:"cname"`
	Tid   int    `form:"tid"`
}
