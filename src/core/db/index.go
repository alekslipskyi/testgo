package db

import (
	"core/db/create"
	"core/db/find"
	"core/db/types"
)

type Instance struct {
	Name  string
	Model interface{}
}

func (entity *Instance) Find(options types.QueryOptions) interface{} {
	findInstance := find.SFind{entity.Name, entity.Model}
	return findInstance.Find(options)
}

func (entity *Instance) FindById(id int8, options types.QueryOptions) interface{} {
	findInstance := find.SFind{entity.Name, entity.Model}
	return findInstance.Find(options)
}

func (entity *Instance) Create(data map[string]interface{}) interface{} {
	createInstance := create.SCreate{entity.Name, entity.Model}
	return createInstance.Create(data)
}
