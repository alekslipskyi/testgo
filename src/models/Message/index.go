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

func Find(query types.QueryOptions) []Message {
	result := getInstance().FindMany(query)

	messages := make([]Message, len(result))

	for key, message := range result {
		messages[key] = message.(Message)
	}

	return messages
}

func getInstance() *db.Instance {
	return &db.Instance{"messages", &Message{}}
}
