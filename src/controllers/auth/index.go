package auth

import (
	"lib/Router"
)

func Routes() {
	controller := Controller{}
	router := Router.Instance{Prefix: "/auth"}

	router.GET("/token", controller.handleAuth)
}
