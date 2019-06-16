package list

import (
	"helpers/auth"
	"lib/Router"
)

func Routes() {
	controller := Controller{}
	router := Router.Instance{Prefix: "/list"}

	router.GET("/", controller.index, auth.IsAuthenticated)
}
