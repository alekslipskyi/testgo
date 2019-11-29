package ChannelUsers

import (
	"core/db"
	"core/db/types"
)

type ChannelUsers struct {
	UserId    int64
	ChannelId int64
}

func (channelUser *ChannelUsers) IsExist() bool {
	return channelUser.UserId != 0
}

func Create(data map[string]interface{}) error {
	_, err := getInstance().Create(data)
	return err
}

func Find(options types.QueryOptions) ChannelUsers {
	return getInstance().Find(options).(ChannelUsers)
}

func getInstance() *db.Instance {
	return &db.Instance{"channel_users", &ChannelUsers{}}
}
