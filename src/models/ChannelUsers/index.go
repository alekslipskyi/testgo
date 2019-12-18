package ChannelUsers

import (
	"core/db"
	"core/db/types"
	"core/logger"
	"encoding/json"
	"strings"
)

var log = logger.Logger{Context:"Model channel users"}

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

func GetUserIDS(channelID int64) []float64 {
	var userIDS interface{}
	var ids []float64
	newInstance := db.Instance{Name:"channel_users"}
	rows := newInstance.FindCustom(types.QueryOptions{
		Attributes: types.Attributes{"array_to_json(array_agg(user_id)) as user"},
		Where: types.Where{"channel_id": channelID},
	})

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&userIDS)

		err := json.Unmarshal([]byte(string(userIDS.([]uint8))), &ids)
		log.LogOnError(err, "error from unmarshal data")
	}

	return ids
}

func Find(options types.QueryOptions) ChannelUsers {
	return getInstance().Find(options).(ChannelUsers)
}

func getInstance() *db.Instance {
	return &db.Instance{"channel_users", &ChannelUsers{}}
}
