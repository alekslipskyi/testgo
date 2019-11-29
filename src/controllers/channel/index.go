package channel

import (
	"core/Router"
	"helpers/auth"
)

func Routes() {
	router := Router.Instance{Prefix: "/channel"}
	controller := Controller{}
	controller.Init()

	router.DELETE("/{channelID}", controller.drop, auth.IsAuthenticated)

	router.GET("/", controller.index, auth.IsAuthenticated)
	router.POST("/", controller.create, auth.IsAuthenticated, bodyChannelCreate)
	router.PUT("/{userID}/invite/{channelID}", controller.invite, auth.IsAuthenticated, paramsChannelInvite)
}
