package list

import (
	"core/Router"
	"helpers/auth"
)

func Routes() {
	controller := Controller{}
	router := Router.Instance{Prefix: "/list"}

	router.GET("/", controller.index, auth.IsAuthenticated)
	router.GET("/{id}", controller.get, auth.IsAuthenticated)
	router.DELETE("/{id}", controller.delete, auth.IsAuthenticated)
	router.PATCH("/{id}", controller.update, auth.IsAuthenticated)
	router.POST("/", controller.create, bodyCreate, auth.IsAuthenticated)
}
