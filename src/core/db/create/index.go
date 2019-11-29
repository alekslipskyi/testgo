package create

import (
	"core/db/connect"
	"core/db/converter"
	"core/db/find"
	"core/db/types"
	"core/logger"
	"fmt"
	"strings"
)

type SCreate struct {
	Name  string
	Model interface{}
}

var log = logger.Logger{"create instance DB"}

func (entity *SCreate) Create(data map[string]interface{}, isReturnID bool) (int64, error) {
	var ID int64
	returning := "RETURNING _id"

	if !isReturnID {
		returning = ""
	}

	values, keys := converter.ParseData(data)
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) %s",
		entity.Name,
		strings.Join(keys, ","),
		strings.Join(values, ","),
		returning)
	log.Debug("query:", query)

	err := connect.DB.QueryRow(query).Scan(&ID)

	if err != nil {
		log.Error(err)
		return 0, err
	}

	return ID, err
}

func (entity *SCreate) CreateAndFind(data map[string]interface{}) interface{} {
	id, _ := entity.Create(data, true)
	findInstance := find.SFind{entity.Name, entity.Model}
	return findInstance.FindOne(types.QueryOptions{Where: types.Where{"_id": id}})
}
