package messages

import (
	"constants/requestError"
	"core/Router"
	"core/amqp"
	"core/db/types"
	"models/ChannelUsers"
	"models/Message"
	"net/http"
	"os"
)

var broker = amqp.Init()

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

func (conn *Controller) Update(ctx Router.Context) {
	message := Message.FindByID(ctx.Params["messageID"].(int64))

	if !message.IsExist() {
		ctx.Reject(requestError.NOT_FOUND)
		return
	}

	message.Update(ctx.Body)

	ctx.Send("ok", http.StatusOK)
}

func (conn *Controller) Create(ctx Router.Context) {
	Message.Create(ctx.Params["ChannelID"].(int64), ctx.User.ID, ctx.Body)

	userIDS := ChannelUsers.GetUserIDS(ctx.Params["ChannelID"].(int64))

	broker.SendMessage(os.Getenv("QUEUE_MESSAGE"), map[string]interface{} {
		"type": "messages",
		"toUsers": userIDS,
		"data": map[string]interface{} {
			"message": ctx.Body,
		},
	})

	ctx.Send("ok", http.StatusOK)
}
