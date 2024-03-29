package Message

import (
	"core/db"
	"core/db/types"
)

type Message struct {
	Body      string                 `json:"body,omitempty"`
	Fileurl   string                 `json:"file_url,omitempty"`
	ID        int64                  `json:"id,omitempty"`
	User      map[string]interface{} `json:"user,omitempty"`
	Channelid int64                  `json:"channel_id,omitempty"`
}

func (mess *Message) IsExist() bool {
	return mess.ID != 0
}

func (mess *Message) Update(data map[string]interface{}) {
	getInstance().Update(data, types.Where{"_id": mess.ID})
}

func FindByID(messageID int64) Message {
	return getInstance().FindById(messageID, []string{"_id"}).(Message)
}

func FindOne(where types.Where) Message {
	return getInstance().Find(types.QueryOptions{
		Attributes: types.Attributes{"messages.body", "messages.file_url", "messages._id", "messages.channel_id", "json_build_object('username', u.username, '_id', u._id) as user"},
		Includes: []types.Include{{
			TableName:  "users",
			FkTableId:  "user_id",
			RefTableID: "u._id",
			AS:         "u",
		}},
		Where: where,
		AS: "messages",
	}).(Message)
}

func Find(query types.QueryOptions) []Message {
	result := getInstance().FindMany(query)

	messages := make([]Message, len(result))

	for key, message := range result {
		messages[key] = message.(Message)
	}

	return messages
}

func CreateAndReturnID(channelID int64, userID int64, payload map[string]interface{}) (int64, error) {
	return getInstance().CreateAndReturnID(map[string]interface{}{
		"channel_id": channelID,
		"user_id":    userID,
		"file_url":   payload["fileURL"],
		"body":       payload["body"],
	})
}

func Create(channelID int64, userID int64, payload map[string]interface{}) {
	_, _ = getInstance().Create(map[string]interface{}{
		"channel_id": channelID,
		"user_id":    userID,
		"file_url":   payload["fileURL"],
		"body":       payload["body"],
	})
}

func getInstance() *db.Instance {
	return &db.Instance{"messages", &Message{}}
}
