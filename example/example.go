package main

import (
	"jumpserver-sdk/api"
	"fmt"
)

func main() {
	Api := api.New("http://127.0.0.1:8080", "safsdf", "asdfasdf")
	fmt.Println(Api.Users.UserProfile())
}
