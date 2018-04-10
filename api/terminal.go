package api

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
	TerminalViewSet()
}

func (t *TerminalServer) TerminalViewSet(appid string) {
	data := app.CreateQueryData()
	data["name"] = appid
	res, _ := app.Http("POST", Actions["terminal-register"], data)

}
