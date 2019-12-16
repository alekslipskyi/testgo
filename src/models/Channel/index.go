package Channel

import (
	"core/db"
	"core/db/connect"
	"core/db/types"
	"core/logger"
	"fmt"
	"github.com/lib/pq"
)

type Channel struct {
	Name  string  `json:"name,omitempty"`
	Users []int64 `json:"users,omitempty"`
	ID    int64   `json:"id,omitempty"`
}

func (channel *Channel) IsExist() bool {
	return channel.ID != 0
}

func (channel *Channel) Drop() bool {
	return getInstance().Drop(types.QueryOptions{Where: types.Where{"_id": channel.ID}})
}

func FindChannels(userId int64) []Channel {
	var channels []Channel
	log := logger.Logger{Context: " find channel with custom query "}

	query := fmt.Sprintf(`
		select *, (select array_agg(user_id) from channel_users where channel_id=_id)
		from channels
		where %d = ANY (select unnest(array_agg(user_id)) from channel_users where channel_id=_id)
	`, userId)

	log.Debug("query are", query)

	rows, err := connect.DB.Query(query)

	defer rows.Close()

	if err != nil {
		log.Error("error from sql query", err)
		return []Channel{}
	}

	for rows.Next() {
		channel := Channel{}
		if err := rows.Scan(&channel.Name, &channel.ID, pq.Array(&channel.Users)); err != nil {
			log.Error("error from scan results", err)
			return channels
		}

		channels = append(channels, channel)
	}

	log.Debug("result are----", channels)

	return channels
}

func Find(options types.QueryOptions) []Channel {
	result := getInstanceChannelUsers().FindMany(options)
	listResult := make([]Channel, len(result))
	for key, val := range result {
		listResult[key] = val.(Channel)
	}

	return listResult
}

func FindOnlyChannel(where types.Where) Channel {
	return getInstance().Find(types.QueryOptions{Where: where, Attributes: types.Attributes{"name", "_id"}}).(Channel)
}

func FindOne(options types.QueryOptions) Channel {
	return getInstance().Find(options).(Channel)
}

func Create(data map[string]interface{}) (int64, error) {
	return getInstance().CreateAndReturnID(data)
}

func getInstanceChannelUsers() *db.Instance {
	return &db.Instance{"channel_users", &Channel{}}
}

func getInstance() *db.Instance {
	return &db.Instance{"channels", &Channel{}}
}
