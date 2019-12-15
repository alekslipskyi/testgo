package channel

import (
	"constants/dbError"
	"constants/permissions"
	"constants/requestError"
	"core/Router"
	"core/db"
	"core/db/types"
	"core/logger"
	"models/Channel"
	"models/ChannelUsers"
	"net/http"
)

var log = logger.Logger{"controller channel"}

type Controller struct {
	channelUsers db.Instance
}

func (entity *Controller) Init() {
}

func (entity *Controller) index(ctx Router.Context) {
	ctx.SendJson(Channel.FindChannels(ctx.User.ID), http.StatusOK)
}

func (entity *Controller) create(ctx Router.Context) {
	channelID, err := Channel.Create(ctx.Body)

	if err != nil {
		if err.Error() == dbError.DUPLICATE_KEY {
			ctx.Reject(CHANNEL_ALREADY_EXISTS)
		} else {
			ctx.Reject(requestError.SOMETHING_WENT_WRONG)
		}
		return
	}

	err2 := ChannelUsers.Create(map[string]interface{}{
		"user_id":     ctx.User.ID,
		"channel_id":  channelID,
		"permissions": "rwdui",
	})

	channel := Channel.FindOne(types.QueryOptions{
		Attributes: types.Attributes{"_id", "name", "array_to_json(array_agg(cs.user_id)) AS Users"},
		Where:      types.Where{"_id": channelID},
		Includes: []types.Include{
			{TableName: "channel_users", FkTableId: "channel_id", RefTableID: "_id", AS: "cs"},
		},
		GroupBy: types.GroupBy{"_id"},
	})

	if err2 != nil && err2.Error() != dbError.NO_ROWS {
		log.Debug("error from creating channel users", err2)
		channel.Drop()
		ctx.Reject(requestError.SOMETHING_WENT_WRONG)
		return
	}

	ctx.SendJson(channel, http.StatusOK)
}

func (entity *Controller) drop(ctx Router.Context) {
	channelUser := ChannelUsers.Find(types.QueryOptions{Where: types.Where{
		"user_id":    ctx.User.ID,
		"channel_id": ctx.Params["channelID"],
	}})

	if !channelUser.IsExist() {
		ctx.Reject(requestError.NOT_FOUND)
		return
	}

	if !channelUser.HasPermission(permissions.DROP) {
		ctx.Reject(NOT_ALLOWED_TO_DROP)
		return
	}

	channel := Channel.FindOne(types.QueryOptions{
		Attributes: types.Attributes{"_id"},
		Where:      types.Where{"_id": ctx.Params["channelID"]},
	})

	if !channel.IsExist() {
		ctx.Reject(requestError.NOT_FOUND)
		return
	}

	ok := channel.Drop()

	if !ok {
		ctx.Reject(requestError.SOMETHING_WENT_WRONG)
		return
	}

	ctx.Send("ok", http.StatusOK)
}

func (entity *Controller) invite(ctx Router.Context) {
	currentUser := ChannelUsers.FindOne(ctx.User.ID, ctx.Params["channelID"].(int64))

	if !currentUser.HasPermission(permissions.INVITE) {
		ctx.Reject(NOT_ALLOWED_TO_INVITE)
		return
	}

	channelUser := ChannelUsers.FindOne(ctx.Params["userID"].(int64), ctx.Params["channelID"].(int64))

	if channelUser.IsExist() {
		ctx.Reject(USER_ALREADY_INVITED)
		return
	}

	err := ChannelUsers.Create(map[string]interface{}{
		"user_id":    ctx.Params["userID"].(int64),
		"channel_id": ctx.Params["channelID"].(int64),
	})

	if err != nil && err.Error() != dbError.NO_ROWS {
		ctx.Reject(requestError.NOT_FOUND)
		return
	}

	ctx.Send("ok", http.StatusOK)
}
