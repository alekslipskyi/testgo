package List

import (
	"core/db"
	"core/db/types"
)

type List struct {
	Title       string `json:"title, omitempty"`
	ID          int64  `json:"id, omitempty"`
	Description string `json:"description, omitempty"`
	User_id     int64  `json:"user_id, omitempty"`
}

func (list *List) UpdateAndFind(data map[string]interface{}) List {
	return getInstance().UpdateAndFind(data, types.Where{"_id": list.ID}).(List)
}

func (list *List) IsNotExist() bool {
	return list.ID == 0
}

func (list *List) Drop() bool {
	return getInstance().Drop(types.QueryOptions{Where: types.Where{"_id": list.ID}})
}

func FindById(id int64, userId int64) List {
	return getInstance().Find(types.QueryOptions{Where: types.Where{"user_id": userId, "_id": id}}).(List)
}

func FindManyByUser(userId int64) []List {
	result := getInstance().FindMany(types.QueryOptions{Where: types.Where{"user_id": userId}})
	listResult := make([]List, len(result))
	for key, val := range result {
		listResult[key] = val.(List)
	}

	return listResult
}

func CreateAndFind(data map[string]interface{}) List {
	return getInstance().CreateAndFind(data).(List)
}

func getInstance() *db.Instance {
	return &db.Instance{"list", &List{}}
}
