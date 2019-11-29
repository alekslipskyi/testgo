package db

import (
	"core/db/create"
	"core/db/drop"
	"core/db/find"
	"core/db/types"
	"core/db/update"
)

type Instance struct {
	Name  string
	Model interface{}
}

func (entity *Instance) Find(options types.QueryOptions) interface{} {
	findInstance := find.SFind{entity.Name, entity.Model}
	return findInstance.FindOne(options)
}

func (entity *Instance) FindMany(options types.QueryOptions) []interface{} {
	findInstance := find.SFind{entity.Name, entity.Model}
	return findInstance.FindMany(options)
}

func (entity *Instance) FindById(id int64, attributes []string) interface{} {
	findInstance := find.SFind{entity.Name, entity.Model}
	return findInstance.FindOne(types.QueryOptions{
		Attributes: attributes,
		Where:      types.Where{"_id": id},
	})
}

func (entity *Instance) Drop(options types.QueryOptions) bool {
	dropInstance := drop.SDrop{entity.Name}
	return dropInstance.Drop(options)
}

func (entity *Instance) Create(data map[string]interface{}) (interface{}, error) {
	createInstance := create.SCreate{entity.Name, entity.Model}
	return createInstance.Create(data, false)
}

func (entity *Instance) CreateAndReturnID(data map[string]interface{}) (int64, error) {
	createInstance := create.SCreate{entity.Name, entity.Model}
	return createInstance.Create(data, true)
}

func (entity *Instance) CreateAndFind(data map[string]interface{}) interface{} {
	createInstance := create.SCreate{entity.Name, entity.Model}
	return createInstance.CreateAndFind(data)
}

func (entity *Instance) Update(data map[string]interface{}, where types.Where) bool {
	updateInstance := update.SUpdate{entity.Name, entity.Model}
	return updateInstance.Update(data, where)
}

func (entity *Instance) UpdateAndFind(data map[string]interface{}, where types.Where) interface{} {
	updateInstance := update.SUpdate{entity.Name, entity.Model}
	return updateInstance.UpdateAndFind(data, where)
}
