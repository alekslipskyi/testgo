package list

import (
	"core/validation"
	"core/validation/types"
)

var bodyCreate = validation.IsValid("body", validation.MustBe{
	"title":       types.String{Min: 1, Max: 6, Required: true},
	"description": types.String{Min: 1, Required: true},
})

var bodyUpdate = validation.IsValid("body", validation.MustBeOneOf{
	"title":       types.String{Min: 1, Max: 6},
	"description": types.String{Min: 1},
})
