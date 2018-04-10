package api

import "net/http"

//jms服务器配置
type Server struct {
	Url        string
	AppName    string
	AppId      string
	appKeyPath string
	token      string
	accessKey  string
	secret     string
	http       *http.Client
	Users      UserServer
	Terminal   TerminalServer
}

// api操作对应URL
type Action map[string]string

// 服务器登陆凭证
type LoginCredit struct {
	Ip         string
	Port       int
	Username   string
	PrivateKey string `json:"private_key"`
}

// 服务器列表单项
type Machine struct {
	Sid    int
	Name   string
	Ip     string
	Port   int
	Remark string
	Users  []MachineUser
	Type   string
	Uuid   string
}

//服务器组列表
type MachineGroup struct {
	Gid    int
	Name   string
	Remark string
}

// 服务器列表可用用户
type MachineUser struct {
	Uid      int
	Username string
}

// 用户pubkey返回
type UserPubKey struct {
	Ticket string
	Key    string
}

// 监控权限返回
type ResponsePass struct {
	Pass bool
}
