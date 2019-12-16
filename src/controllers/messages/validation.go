package messages

import (
	"core/validation"
	"core/validation/types"
)

var checkGetMessagesParams = validation.MustBe{
	"ChannelID": types.Number{Min: 1},
}

var checkCreateMessageParams = validation.MustBe{
	"ChannelID": types.Number{Min: 1},
}

var checkCreateMessageBody = validation.MustBeOneOf{
	"fileURL": types.String{Min: 1, Max: 50},
	"body":    types.String{Min: 1, Max: 50},
}
