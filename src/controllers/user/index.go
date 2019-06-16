package user

import (
	"lib/Router"
)

func Routes() {
	controller := Controller{}
	controller.Init()

	router := Router.Instance{Prefix: "/user"}

	router.POST("/", controller.handleSignUp, checkBodySignUp)
}
