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
}
