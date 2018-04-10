package api

import (
	log "github.com/liuzheng/golog"
	"encoding/json"
	"github.com/pkg/errors"
	"fmt"
	"os"
)

type Terminal struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	RemoteAddr     string `json:"remote_addr"`
	SshPort        uint16 `json:"ssh_port"`
	HttpPort       uint16 `json:"http_port"`
	CommandStorage string `json:"command_storage"`
	ReplayStorage  string `json:"replay_storage"`
	IsAccepted     bool   `json:"is_accepted"`
	IsDeleted      bool   `json:"is_deleted"`
	DateCreated    string `json:"date_created"`
	Comment        string `json:"comment"`
	//User = models.OneToOneField(User, related_name='terminal', verbose_name='Application User', null=True, on_delete=models.CASCADE)
}
type TerminalServer struct {
}
type TerminalInterface interface {
	TerminalRegister() error
	TerminalAccessKey()
}

type TerminalView struct {
	Id      string `json:"id"`
	Token   string `json:"token"`
	Message string `json:"msg"`
}

func (t *TerminalServer) TerminalRegister() error {
	data := app.CreateQueryData()
	data["name"] = app.AppName
	res, _ := app.Http("POST", Actions["terminal-register"], nil, data)
	Res := struct {
		Id      string `json:"id"`
		Token   string `json:"token"`
		Message string `json:"msg"`
	}{}
	json.Unmarshal(res, &Res)
	if Res.Id == "" || Res.Token == "" {
		log.Error("TerminalRegister", "%v", Res.Message)
		return errors.New(Res.Message)
	} else {
		log.Info("TerminalRegister", "%v", Res.Message)
		app.Token = Res.Token
		app.AppId = Res.Id
	}
	return nil
}

func (t *TerminalServer) TerminalAccessKey() {
	params := app.CreateQueryData()
	params["token"] = app.Token
	res, _ := app.Http("GET", fmt.Sprintf(Actions["terminal-access-key"], app.AppId), params, nil)
	Res := struct {
		AccessKey struct {
			Id     string `json:"id"`
			Secret string `json:"secret"`
		} `json:"access_key"`
	}{}
	json.Unmarshal(res, &Res)
	f, err := os.Create(app.appKeyPath)
	if err != nil {
		log.Error("TerminalAccessKey", "%v", err)
		return
	}
	f.WriteString(Res.AccessKey.Id + ":" + Res.AccessKey.Secret)
	f.Close()
}
