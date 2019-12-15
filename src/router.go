package router

import (
	"controllers/auth"
	"controllers/channel"
	"controllers/list"
	"controllers/messages"
	"controllers/user"
	"core/Router"
	"core/db/connect"
	"net/http"
)

func Handler() http.Handler {
	handler := http.NewServeMux()

	router := Router.New{}
	connect.Init()

	user.Routes()
	auth.Routes()
	list.Routes()
	channel.Routes()
	messages.Routes()

	handler.HandleFunc("/", router.EntryPoint)

	return handler
}
