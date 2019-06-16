package user

import (
	"lib/validation"
	"lib/validation/types"
)

var checkBodySignUp = validation.IsValid("body", validation.MustBe{
	"username":  types.String{Required: true, Min: 1},
	"firstName": types.String{Required: true, Min: 1},
	"lastName":  types.String{Required: true, Min: 1},
	"password":  types.String{Required: true, Min: 6, Max: 10},
})
