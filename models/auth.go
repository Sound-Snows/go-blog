package models

//Auth 用户模型
type Auth struct {
	ID       int    `gorm:"primary_key" josn:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//CheckAuth 根据账号和密码查询用户
func CheckAuth(username, password string) bool {
	var auth Auth
	db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	if auth.ID > 0 {
		return true
	}
	return false
}
