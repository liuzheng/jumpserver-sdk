package main

import (
	"jumpserver-sdk/api"
	"flag"
	log "github.com/liuzheng/golog"
)

func main() {
	flag.Parse()
	log.Logs("", "DEBUG", "ERROR")
	Api := api.New("http://127.0.0.1:5000", "cccc", "./keys/.access_key")
	//Api.Terminal.TerminalRegister()

	Api.Terminal.TerminalAccessKey()
}
