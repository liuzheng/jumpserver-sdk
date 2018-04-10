package api

import (
	"bytes"
	"net/http"
	"io/ioutil"
	"encoding/json"
	log "github.com/liuzheng/golog"
	"bufio"
	"os"
	"strings"
	"time"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"encoding/hex"
)

const TimeFormat = "2006-01-02 15:04:05"

var (
	user_api     = "/api/users/"
	terminal_api = "/api/terminal/"
	Actions      = Action{
		"terminal-heartbeat":             "/api/terminal/v1/terminal/status/",
		"session-list":                   "/api/terminal/v1/sessions/",
		"session-detail":                 "/api/terminal/v1/sessions/%s/",
		"session-command":                "/api/terminal/v1/command/",
		"user-assets":                    "/api/perms/v1/user/%s/assets/",
		"user-asset-groups":              "/api/perms/v1/user/%s/nodes-assets/",
		"user-nodes-assets":              "/api/perms/v1/user/%s/nodes-assets/",
		"my-profile":                     "/api/users/v1/profile/",
		"system-user-auth-info":          "/api/assets/v1/system-user/%s/auth-info/",
		"validate-user-asset-permission": "/api/perms/v1/asset-permission/user/validate/",
		"finish-task":                    "/api/terminal/v1/tasks/%s/",
		"asset":                          "/api/assets/v1/assets/%s/",
		"system-user":                    "/api/assets/v1/system-user/%s",
		"user-user":                      "/api/users/v1/users/%s/",
		"token-asset":                    "/api/users/v1/connection-token/?token=%s",
		"domain-detail":                  "/api/assets/v1/domain/%s/",
		"ftp-log-list":                   "/api/audits/v1/ftp-log/",

		// user api
		"users":                  user_api + "v1/users",                    // UserViewSet
		"user-group":             user_api + "v1/groups",                   // UserGroupViewSet
		"user-token":             user_api + "v1/token/",                   // UserToken
		"connection-token":       user_api + "v1/connection-token/",        // UserConnectionTokenApi
		"user-profile":           user_api + "v1/profile/",                 // UserProfile
		"user-auth":              user_api + "v1/auth/",                    // UserAuthApi
		"change-user-password":   user_api + "v1/users/%s/password/",       // ChangeUserPasswordApi
		"user-reset-password":    user_api + "v1/users/%s/password/reset/", // UserResetPasswordApi
		"user-public-key-reset":  user_api + "v1/users/%s/pubkey/reset/",   // UserResetPKApi
		"user-public-key-update": user_api + "v1/users/%s/pubkey/update/",  // UserUpdatePKApi
		"user-update-group":      user_api + "v1/users/%s/groups/",         // UserUpdateGroupApi
		"user-group-update-user": user_api + "v1/groups/%s/users/",         // UserGroupUpdateUserApi

		// terminal api
		"terminal-status":     terminal_api + "v1/terminal/%s/status",     // StatusViewSet
		"terminal-sessions":   terminal_api + "v1/terminal/%s/sessions",   // SessionViewSet
		"tasks":               terminal_api + "v1/tasks",                  // TaskViewSet
		"terminal-register":   terminal_api + "v1/terminal/",              // TerminalViewSet
		"command":             terminal_api + "v1/command",                // CommandViewSet
		"session":             terminal_api + "v1/sessions",               // SessionViewSet
		"status":              terminal_api + "v1/status",                 // StatusViewSet
		"session-replay":      terminal_api + "v1/sessions/%s/replay/",    // SessionReplayViewSet
		"kill-session":        terminal_api + "v1/tasks/kill-session/",    // KillSessionAPI
		"terminal-access-key": terminal_api + "v1/terminal/%s/access-key", // TerminalTokenApi
		"terminal-config":     terminal_api + "v1/terminal/config",        // TerminalConfig
	}
	app = Server{
		http: &http.Client{},
	}
)

//初始化一个ApiServer
func New(JmsUrl, AppName, AppKeyPath string) *Server {
	app.Url = JmsUrl
	app.AppName = AppName
	app.appKeyPath = AppKeyPath
	_, err := os.Stat(AppKeyPath)
	if err == nil {
		file, _ := os.Open(AppKeyPath)
		f := bufio.NewReader(file)
		line, _, _ := f.ReadLine()
		if len(line) != 0 {
			q := strings.Split(string(line), ":")
			app.accessKey = q[0]
			app.secret = q[1]
		}
	} else if os.IsNotExist(err) {
		log.Warn("New", "file %s not exists", AppKeyPath)
	} else {
		log.Warn("New", "file %s stat error: %v", AppKeyPath, err)
	}

	return &app
}

func (s *Server) Check() bool {
	s.Users.UserViewSet()
	return true
}

func (s *Server) Http(method, uri string, params, data map[string]interface{}) (resBody []byte, err error) {

	dataJson, err := json.Marshal(data)
	if err != nil {
		log.Error("Http", "%v", err)
		return
	}
	log.Debug("Http", "%v", string(dataJson))
	reqNew := bytes.NewBuffer(dataJson)
	request, err := http.NewRequest(method, s.Url+uri, reqNew)
	if err != nil {
		log.Error("Http", "%v", err)
		return
	}

	q := request.URL.Query()
	for k, v := range params {
		q.Add(k, v.(string))
	}
	request.URL.RawQuery = q.Encode()

	request.Header.Set("Content-type", "application/json")
	for k, v := range s.AccessKeyAuth() {
		request.Header.Set(k, v)
	}

	//发起HTTP请求
	response, err := s.http.Do(request)
	if err != nil {
		if err != nil {
			log.Error("Http", "http.Do: %v", err)
			return
		}
	}
	resBody, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error("Http", "ioutil.ReadAll: %v", err)
		return
	}
	log.Debug("RESPONSE", "%v", string(resBody))
	return
}

//创建请求数据Map
func (s *Server) CreateQueryData() map[string]interface{} {
	return make(map[string]interface{})
}

func (s *Server) AccessKeyAuth() (result map[string]string) {
	result = make(map[string]string)
	result["Date"] = time.Now().Format("Mon, 2 Jan 2006 15:04:05 GMT")
	md5sum := md5.New()
	md5sum.Write([]byte(fmt.Sprintf("%s\n%s", s.secret, result["Date"])))
	signature := base64.StdEncoding.EncodeToString([]byte(hex.EncodeToString(md5sum.Sum(nil))))
	result["Authorization"] = "Sign " + s.accessKey + ":" + signature
	return
}

func (s *Server) TokenAuth() (result map[string]string) {
	result = make(map[string]string)
	result["Authorization"] = "Bearer " + s.token
	return
}

func (s *Server) SessionAuth() (result map[string]string) {
	return
}
