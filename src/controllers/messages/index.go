package messages

import (
	"constants/permissions"
	"core/Router"
	"core/validation"
	"helpers"
	"helpers/auth"
)

func Routes() {
	controller := Controller{}

	router := Router.Instance{Prefix: "/message"}

	router.GET("/{channelID}/messages", controller.Get, auth.IsAuthenticated, validation.IsValid("params", checkGetMessagesParams))
	router.POST("/{channelID}", controller.Create,
		auth.IsAuthenticated,
		validation.IsValid("params", checkCreateMessageParams),
		validation.IsValid("body", checkCreateMessageBody),
		helpers.HasPermissions(permissions.WRITE, WRITE_IS_NOT_ALLOWED),
	)
}
