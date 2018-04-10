package api

import (
	"time"
	log "github.com/liuzheng/golog"
	"encoding/json"
)


type User struct {
	Id              string      `json:"id"`
	Username        string      `json:"username"`
	Name            string      `json:"name"`
	Email           string      `json:"email"`
	IsActive        bool        `json:"is_active"`
	IsSuperuser     bool        `json:"is_superuser"`
	Role            string      `json:"role"`
	GroupDisplay    string      `json:"group_display"`
	Groups          []string    `json:"groups"`
	WeChat          string      `json:"wechat"`
	Avatar          string      `json:"avatar"`
	Comment         string      `json:"comment"`
	EnableOtp       bool        `json:"enable_otp"`
	IsFirstLogin    bool        `json:"is_first_login"`
	SecretKeyOtp    string      `json:"secret_key_otp"`
	CreatedBy       string      `json:"created_by"`
	UserPermissions []string    `json:"user_permissions"`
	GetRoleDisplay  string      `json:"get_role_display"`
	IsValid         bool        `json:"is_valid"`
	Phone           interface{} `json:"phone"`
	DateExpired     string      `json:"date_expired"`
	DateJoined      string      `json:"date_joined"`
	LastLogin       string      `json:"last_login"`
}

type UserGroup struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Comment     string    `json:"comment"`
	DateCreated time.Time `json:"date_created"`
}
type UsersInterface interface {
	UserViewSet() []User
	UserProfile() User
}
type UserServer struct {
}

func (u *UserServer) UserViewSet() (users []User) {
	res, _ := app.Http("GET", Actions["users"], nil)
	log.Debug("UserViewSet", "%v", res)
	return
}
func (u *UserServer) UserProfile() (user User) {
	res, _ := app.Http("GET", Actions["user-profile"], nil)
	log.Debug("UserViewSet", "%v", string(res))
	ok := json.Unmarshal(res, &user)
	log.Debug("UserViewSet", "%v", ok)

	return
}

func (u *UserServer) UserGroupViewSet() (user User) {
	return
}
