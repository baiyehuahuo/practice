package model

type User struct {
	Username string   `form:"username"`
	Password string   `form:"password"`
	Hobby    []string `form:"hobby"`
}
