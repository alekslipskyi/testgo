package find

import (
	"core/db/connect"
	"core/db/converter"
	"core/db/types"
	"core/logger"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
)

type SFind struct {
	Name  string
	Model interface{}
}

func (entity *SFind) FindOne(options types.QueryOptions) interface{} {
	var log = logger.Logger{"DB FIND ONE"}

	maxColumns := len(options.Attributes)
	sqlQuery, attributes := entity.generateQuery(options)

	query := fmt.Sprintf(sqlQuery, attributes, entity.Name)
	log.Debug("query:", query)

	row := connect.DB.QueryRow(query)
	result := entity.PointToModel(row, maxColumns, options).Interface()

	log.Debug("result:", result)
	return result
}

func (entity *SFind) FindMany(options types.QueryOptions) []interface{} {
	var log = logger.Logger{"DB FIND MANY"}

	maxColumns := len(options.Attributes)
	var result []interface{}
	sqlQuery, attributes := entity.generateQuery(options)

	query := fmt.Sprintf(sqlQuery, attributes, entity.Name)
	log.Debug("query:", query)

	rows, err := connect.DB.Query(query)

	defer rows.Close()

	if err != nil {
		log.Error("execute query: ", query, "error: ", err)
		return result
	}

	for rows.Next() {
		result = append(result, entity.PointToModel(rows, maxColumns, options).Interface())
	}

	log.Debug("result:", result)
	return result
}

func (entity *SFind) PointToModel(result interface{}, maxColumns int, options types.QueryOptions) reflect.Value {
	var log = logger.Logger{"DB FIND"}

	modelPointer := entity.generateModelPointer()
	pointers := entity.generatePointers(modelPointer, maxColumns)

	var err error

	if reflect.TypeOf(result) == reflect.TypeOf(&sql.Rows{}) {
		err = result.(*sql.Rows).Scan(pointers...)
	} else {
		err = result.(*sql.Row).Scan(pointers...)
	}

	if err != nil {
		log.Error("scan error", err)
		return modelPointer
	}

	if maxColumns == 0 {
		for i := range pointers {
			val := pointers[i].(*interface{})
			modelPointer.Field(i).Set(reflect.ValueOf(*val))
		}
	} else {
		for key, property := range options.Attributes {
			val := pointers[key].(*interface{})

			if property == "_id" {
				modelPointer.FieldByName("ID").Set(reflect.ValueOf(*val))
			} else {
				modelPointer.FieldByName(strings.Title(strings.ToLower(property))).Set(reflect.ValueOf(*val))
			}
		}
	}

	return modelPointer
}

func (entity *SFind) generateQuery(options types.QueryOptions) (string, string) {
	sqlQuery := "SELECT %s FROM %s "
	attributes := "*"
	var Where string

	if len(options.Attributes) != 0 {
		attributes = strings.Join(options.Attributes, ", ")
	}

	if options.Where != nil {
		Where = converter.DataToQueryString(options.Where, "=", "and")
	}

	if len(Where) != 0 {
		sqlQuery += "WHERE" + Where
	}

	return sqlQuery, attributes
}

func (entity *SFind) generateModelPointer() reflect.Value {
	elem := reflect.ValueOf(entity.Model).Elem()
	return reflect.Indirect(elem)
}

func (entity *SFind) generatePointers(modelPointer reflect.Value, maxColumns int) []interface{} {
	numColumns := maxColumns

	if maxColumns == 0 {
		numColumns = modelPointer.NumField()
	}

	pointerValues := make([]interface{}, numColumns)
	pointers := make([]interface{}, numColumns)

	for i := range pointers {
		pointers[i] = &pointerValues[i]
	}

	return pointers
}
