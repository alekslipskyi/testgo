package ChannelUsers

import (
	"core/db"
	"core/db/types"
	"strings"
)

type ChannelUsers struct {
	UserId      int64
	ChannelId   int64
	Permissions string
}

func (channelUser *ChannelUsers) IsExist() bool {
	return channelUser.UserId != 0
}

func (channelUser *ChannelUsers) HasPermission(operation rune) bool {
	return strings.Contains(channelUser.Permissions, string(operation))
}

func Create(data map[string]interface{}) error {
	_, err := getInstance().Create(data)
	return err
}

func FindOne(userID int64, channelID int64) ChannelUsers {
	return getInstance().Find(types.QueryOptions{Where: types.Where{"user_id": userID, "channel_id": channelID}}).(ChannelUsers)
}

func Find(options types.QueryOptions) ChannelUsers {
	return getInstance().Find(options).(ChannelUsers)
}

func getInstance() *db.Instance {
	return &db.Instance{"channel_users", &ChannelUsers{}}
}
