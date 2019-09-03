package auth

import (
	"core/Router"
)

func Routes() {
	controller := Controller{}
	router := Router.Instance{Prefix: "/auth"}

	router.GET("/token", controller.handleAuth)
}
