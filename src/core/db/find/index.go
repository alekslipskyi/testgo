package find

import (
	"core/db/connect"
	"core/db/types"
	"fmt"
	"reflect"
	"strings"
)

type SFind struct {
	Name  string
	Model interface{}
}

func (entity *SFind) FindById(id int8, options types.QueryOptions) interface{} {
	return entity.Find(types.QueryOptions{})
}

func (entity *SFind) Find(options types.QueryOptions) interface{} {
	sqlQuery, attributes := entity.generateQuery(options)

	query := fmt.Sprintf(sqlQuery, attributes, entity.Name)
	row := connect.DB.QueryRow(query)

	modelPointer := entity.generateModelPointer()
	pointers := entity.generatePointers(modelPointer)

	err := row.Scan(pointers...)
	if err != nil {
		return modelPointer.Interface()
	}

	for i := range pointers {
		val := pointers[i].(*interface{})
		modelPointer.Field(i).Set(reflect.ValueOf(*val))
	}

	return modelPointer.Interface()
}

func (entity *SFind) generateQuery(options types.QueryOptions) (string, string) {
	sqlQuery := "SELECT %s FROM %s "
	attributes := "*"
	var Where string

	if len(options.Attributes) != 0 {
		attributes = strings.Join(options.Attributes, ", ")
	}

	if options.Where != nil {
		for key, value := range options.Where {
			Where = fmt.Sprintf(Where+"%s='%s'", key, value)
		}
	}

	if len(Where) != 0 {
		sqlQuery += "WHERE " + Where
	}

	return sqlQuery, attributes
}

func (entity *SFind) generateModelPointer() reflect.Value {
	elem := reflect.ValueOf(entity.Model).Elem()
	return reflect.Indirect(elem)
}

func (entity *SFind) generatePointers(modelPointer reflect.Value) []interface{} {
	pointerValues := make([]interface{}, modelPointer.NumField())
	pointers := make([]interface{}, modelPointer.NumField())

	for i := range pointers {
		pointers[i] = &pointerValues[i]
	}

	return pointers
}
