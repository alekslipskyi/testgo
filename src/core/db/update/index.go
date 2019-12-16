package update

import (
	"core/db/connect"
	"core/db/converter"
	"core/db/find"
	"core/db/types"
	"core/logger"
	"fmt"
)

var log = logger.Logger{Context: "DB UPDATE"}

type SUpdate struct {
	Name  string
	Model interface{}
}

func (entity *SUpdate) Update(data map[string]interface{}, where types.Where) bool {
	valuesToUpdate := converter.DataToQueryString(data, "=", ",")
	whereConverted := converter.DataToQueryString(where, "=", "and")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s", entity.Name, valuesToUpdate, whereConverted)
	log.Debug("query:", query)
	connect.DB.QueryRow(query)
	return true
}

func (entity *SUpdate) UpdateAndFind(data map[string]interface{}, where types.Where) interface{} {
	findInstance := find.SFind{entity.Name, entity.Model}
	entity.Update(data, where)
	return findInstance.FindOne(types.QueryOptions{Where: where})
}
