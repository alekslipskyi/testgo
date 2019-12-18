package channel

import (
	"constants/permissions"
	"core/Router"
	"helpers"
	"helpers/auth"
)

func Routes() {
	router := Router.Instance{Prefix: "/channel"}
	controller := Controller{}
	controller.Init()

	router.DELETE("/{channelID}", controller.drop,
		auth.IsAuthenticated,
		paramsChannelDrop,
		helpers.HasPermissions(permissions.DROP, NOT_ALLOWED_TO_DROP),
	)

	router.GET("/", controller.index, auth.IsAuthenticated)
	router.POST("/", controller.create, auth.IsAuthenticated, bodyChannelCreate)
	router.PUT("/{userID}/invite/{channelID}", controller.invite,
		auth.IsAuthenticated,
		paramsChannelInvite,
		helpers.HasPermissions(permissions.INVITE, NOT_ALLOWED_TO_INVITE),
	)
}
