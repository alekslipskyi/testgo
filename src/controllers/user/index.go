package user

import (
	"helpers/validation"
	"helpers/validation/types"
	"lib/Router"
)

func Routes() {
	controller := Controller{}
	router := Router.Instance{Prefix: "/user"}

	body := validation.MustBe{
		"username":  types.String{Required: true, Min: 3},
		"firstName": types.String{Required: true, Min: 1},
		"lastName":  types.String{Required: true, Min: 1},
		"password":  types.String{Required: true, Min: 6},
	}

	router.POST("/{id}", controller.handleSignUp, validation.IsValid("body", body))
}
