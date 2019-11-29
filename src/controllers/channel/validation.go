package channel

import (
	"core/validation"
	"core/validation/types"
)

var bodyChannelCreate = validation.IsValid("body", validation.MustBe{
	"name": types.String{Min: 1, Max: 10},
})

var paramsChannelInvite = validation.IsValid("params", validation.MustBe{
	"channelID": types.Number{Min: 1},
	"userID":    types.Number{Min: 1},
})

var paramsChannelDrop = validation.IsValid("params", validation.MustBe{
	"channelID": types.Number{Min: 1},
})
