package messages

import (
	"core/Router"
	"core/validation"
	"helpers/auth"
)

func Routes() {
	controller := Controller{}

	router := Router.Instance{Prefix: "/message"}

	router.GET("/{ChannelID}/messages", controller.Get, auth.IsAuthenticated, validation.IsValid("params", checkGetMessagesParams))
	router.POST("/{ChannelID}", controller.Create,
		auth.IsAuthenticated,
		validation.IsValid("params", checkCreateMessageParams),
		validation.IsValid("body", checkCreateMessageBody),
	)
}
