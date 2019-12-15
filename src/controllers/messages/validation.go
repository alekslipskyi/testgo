package messages

import (
	"core/validation"
	"core/validation/types"
)

var checkGetMessagesParams = validation.MustBe{
	"ChannelID": types.Number{Min: 1},
}
