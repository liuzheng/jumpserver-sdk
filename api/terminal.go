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

type TerminalView struct {
	Id      string `json:"id"`
	Token   string `json:"token"`
	Message string `json:"msg"`
}

type TerminalServer struct {
}

func (t *TerminalServer) TerminalRegister() error {
	if app.secret != "" {
		return nil
	}
	data := app.CreateQueryData()
	data.Data["name"] = app.AppName
	res, _ := app.Http("POST", Actions["terminal-register"], data)
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
		app.token = Res.Token
		app.AppId = Res.Id
	}
	return nil
}

func (t *TerminalServer) TerminalAccessKey() {
	if app.secret != "" {
		return
	}
	data := app.CreateQueryData()
	data.Params["token"] = app.token
	res, _ := app.Http("GET", fmt.Sprintf(Actions["terminal-access-key"], app.AppId), data)
	Res := struct {
		AccessKey struct {
			Id     string `json:"id"`
			Secret string `json:"secret"`
		} `json:"access_key"`
	}{}
	json.Unmarshal(res, &Res)
	if Res.AccessKey.Id != "" && Res.AccessKey.Secret != "" {
		f, err := os.Create(app.appKeyPath)
		if err != nil {
			log.Error("TerminalAccessKey", "%v", err)
			return
		}
		app.secret = Res.AccessKey.Secret
		app.accessKey = Res.AccessKey.Id
		f.WriteString(Res.AccessKey.Id + ":" + Res.AccessKey.Secret)
		f.Close()
	}
}
