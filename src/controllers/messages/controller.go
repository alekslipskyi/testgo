package messages

import (
	"core/Router"
	"core/db/types"
	"models/Message"
	"net/http"
)

type Controller struct {
}

func (conn *Controller) Get(ctx Router.Context) {
	messages := Message.Find(types.QueryOptions{
		Attributes: types.Attributes{"messages.body", "messages.file_url", "messages._id", "messages.channel_id", "json_build_object('username', u.username, '_id', u._id) as user"},
		Includes: []types.Include{{
			TableName:  "users",
			FkTableId:  "user_id",
			RefTableID: "u._id",
			AS:         "u",
		}},
		Where: types.Where{
			"user_id":    ctx.User.ID,
			"channel_id": ctx.Params["ChannelID"].(int64),
		},
		AS: "messages",
	})

	ctx.SendJson(messages, http.StatusOK)
}
