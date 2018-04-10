package api

import (
	log "github.com/liuzheng/golog"
	"encoding/json"
	"fmt"
)

type User struct {
	Id              string   `json:"id"`
	Username        string   `json:"username"`
	Name            string   `json:"name"`
	Email           string   `json:"email"`
	IsActive        bool     `json:"is_active"`
	IsSuperuser     bool     `json:"is_superuser"`
	Role            string   `json:"role"`
	GroupDisplay    string   `json:"group_display"`
	Groups          []string `json:"groups"`
	WeChat          string   `json:"wechat"`
	Avatar          string   `json:"avatar"`
	Comment         string   `json:"comment"`
	EnableOtp       bool     `json:"enable_otp"`
	IsFirstLogin    bool     `json:"is_first_login"`
	SecretKeyOtp    string   `json:"secret_key_otp"`
	CreatedBy       string   `json:"created_by"`
	UserPermissions []string `json:"user_permissions"`
	GetRoleDisplay  string   `json:"get_role_display"`
	IsValid         bool     `json:"is_valid"`
	Phone           int      `json:"phone"`
	DateExpired     string   `json:"date_expired"`
	DateJoined      string   `json:"date_joined"`
	LastLogin       string   `json:"last_login"`
}

type UserGroup struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Comment     string   `json:"comment"`
	Users       []string `json:"users"`
	IsDiscard   bool     `json:"is_discard"`
	DiscardTime string   `json:"discard_time"`
	DateCreated string   `json:"date_created"`
	CreatedBy   string   `json:"created_by"`
}

type UserServer struct {
}

// Get all users profile list
func (u *UserServer) UserViewSet() (users []User) {
	res, _ := app.Http("GET", Actions["users"], nil, nil)
	log.Debug("UserViewSet", "%v", string(res))
	return
}

// Get a user profile
func (u *UserServer) GetUserProfile(uid string) (user User) {
	res, _ := app.Http("GET", fmt.Sprintf(Actions["user-user"], uid), nil, nil)
	json.Unmarshal(res, &user)
	return
}

// Get one user group detial
func (u *UserServer) GetUserGroupDetial(gid string) (group UserGroup) {
	res, _ := app.Http("GET", fmt.Sprintf(Actions["user-group"], gid), nil, nil)
	json.Unmarshal(res, &group)
	return
}

// Get all user groups detial list
func (u *UserServer) UserGroupViewSet() (groups []UserGroup) {
	res, _ := app.Http("GET", Actions["user-groups"], nil, nil)
	json.Unmarshal(res, &groups)
	return
}

func (u *UserServer) UserToken() (user User) {
	res, _ := app.Http("GET", Actions["user-token"], nil, nil)
	json.Unmarshal(res, &user)
	return
}

func (u *UserServer) UserConnectionTokenApi() (user User) {
	res, _ := app.Http("GET", Actions["connection-token"], nil, nil)
	json.Unmarshal(res, &user)
	return
}

func (u *UserServer) UserProfile() (user User) {
	res, _ := app.Http("GET", Actions["user-profile"], nil, nil)
	json.Unmarshal(res, &user)
	return
}

func (u *UserServer) UserAuthApi() (user User) {
	res, _ := app.Http("GET", Actions["user-auth"], nil, nil)
	json.Unmarshal(res, &user)
	return
}

func (u *UserServer) ChangeUserPasswordApi(uid string) (user User) {
	res, _ := app.Http("GET", fmt.Sprintf(Actions["change-user-password"], uid), nil, nil)
	json.Unmarshal(res, &user)
	return
}
func (u *UserServer) UserResetPasswordApi(uid string) (user User) {
	res, _ := app.Http("GET", fmt.Sprintf(Actions["user-reset-password"], uid), nil, nil)
	json.Unmarshal(res, &user)
	return
}
func (u *UserServer) UserResetPKApi(uid string) (user User) {
	res, _ := app.Http("GET", fmt.Sprintf(Actions["user-public-key-reset"], uid), nil, nil)
	json.Unmarshal(res, &user)
	return
}
func (u *UserServer) UserUpdatePKApi(uid string) (user User) {
	res, _ := app.Http("GET", fmt.Sprintf(Actions["user-public-key-update"], uid), nil, nil)
	json.Unmarshal(res, &user)
	return
}
func (u *UserServer) UserUpdateGroupApi(uid string) (user User) {
	res, _ := app.Http("GET", fmt.Sprintf(Actions["user-update-group"], uid), nil, nil)
	json.Unmarshal(res, &user)
	return
}
func (u *UserServer) UserGroupUpdateUserApi(gid string) (user User) {
	res, _ := app.Http("GET", fmt.Sprintf(Actions["user-group-update-user"], gid), nil, nil)
	json.Unmarshal(res, &user)
	return
}
