package user

import (
	"core/validation"
	"core/validation/types"
)

var checkBodySignUp = validation.IsValid("body", validation.MustBe{
	"username":  types.String{Min: 1},
	"firstName": types.String{Min: 1},
	"lastName":  types.String{Min: 1},
	"password":  types.String{Min: 6, Max: 10},
})

var checkParamsUserID = validation.IsValid("params", validation.MustBe{"userID": types.Number{}})

var checkBodyUpdate = validation.IsValid("body", validation.MustBeOneOf{
	"username":  types.String{Min: 1},
	"firstName": types.String{Min: 1},
	"lastName":  types.String{Min: 1},
})
