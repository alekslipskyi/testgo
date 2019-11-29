package user

import (
	"core/Router"
	"helpers/auth"
)

func Routes() {
	controller := Controller{}
	controller.Init()

	router := Router.Instance{Prefix: "/user"}

	router.GET("/{userID}", controller.getByID, checkParamsUserID)

	router.GET("/", controller.getMe, auth.IsAuthenticated)
	router.POST("/", controller.handleSignUp, checkBodySignUp)
	router.PUT("/", controller.handleUpdate, checkBodyUpdate, auth.IsAuthenticated)
	router.DELETE("/", controller.delete, auth.IsAuthenticated)
}
