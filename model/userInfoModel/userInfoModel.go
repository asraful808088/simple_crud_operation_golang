package userinfo

type UserInfo struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Staff bool   `json:"staff"`
	Admin bool   `json:"admin"`
}