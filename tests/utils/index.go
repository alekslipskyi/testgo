package utils

import (
	"bytes"
	"core/crypto"
	"encoding/json"
	"fmt"
	"models/Channel"
	"models/ChannelUsers"
	"models/User"
)

func CreateUser(username ...string) User.User {
	userName := "string"

	if len(username) > 0 {
		userName = username[0]
	}

	createdUser := User.CreateAndFind(map[string]interface{}{
		"firstName": "string",
		"lastName":  "string",
		"password":  crypto.GenerateHash("string"),
		"username":  userName,
	})
	createdUser.GenerateToken()
	createdUser.AddAllowIP("127.0.0.1")

	return createdUser
}

func MapToString(m []map[string]interface{}) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%s=\"%s\"\n", key, value)
	}
	return b.String()
}

func CreateChannel(channel map[string]interface{}, userID int64, permissions ...string) float64 {
	channelID, _ := Channel.Create(channel)

	if len(permissions) > 0 {
		_ = ChannelUsers.Create(map[string]interface{}{
			"user_id":     userID,
			"channel_id":  channelID,
			"permissions": permissions[0],
		})
	} else {
		_ = ChannelUsers.Create(map[string]interface{}{
			"user_id":    userID,
			"channel_id": channelID,
		})
	}

	return float64(channelID)
}

func StructToMap(st interface{}) map[string]interface{} {
	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(st)
	json.Unmarshal(inrec, &inInterface)

	return inInterface
}
