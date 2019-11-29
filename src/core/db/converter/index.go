package converter

import "C"
import (
	"core/logger"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var log = logger.Logger{"DB CONVERTER"}

func ParseData(data map[string]interface{}) ([]string, []string) {
	var values []string
	var keys []string

	for key, value := range data {
		switch value.(type) {
		case int64:
			values = append(values, strconv.FormatInt(value.(int64), 10))
		case float64:
			values = append(values, strconv.FormatFloat(value.(float64), 'f', 0, 64))
		default:
			values = append(values, fmt.Sprintf("'%s'", strings.ToLower(value.(string))))
		}

		keys = append(keys, strings.ToLower(key))
	}

	return values, keys
}

func isStringType(value string) bool {
	if _, err := strconv.ParseBool(value); err == nil {
		return false
	}

	if _, err := strconv.ParseFloat(value, 64); err == nil {
		return false
	}

	if _, err := strconv.ParseInt(value, 10, 64); err == nil {
		return false
	}

	return true
}

func castToString(val interface{}) string {
	switch val.(type) {
	case int64:
		return strconv.FormatInt(val.(int64), 10)
	case bool:
		return strconv.FormatBool(val.(bool))
	case float64:
		return strconv.FormatFloat(val.(float64), 'E', -1, 64)
	}

	return val.(string)
}

func DataToQueryString(data map[string]interface{}, intermediateDelimiter, delimiter string) string {
	var queryString []string

	for key, value := range data {
		val := castToString(value)

		if isStringType(val) {
			queryString = append(queryString, " "+key+intermediateDelimiter+fmt.Sprintf("'%s'", val)+" ")
		} else {
			queryString = append(queryString, " "+key+intermediateDelimiter+val+" ")
		}
	}

	return strings.Join(queryString, delimiter)
}

func CastToNeededType(val interface{}, modelPointer reflect.Value, property string) {
	name := property

	if strings.Contains(name, " AS ") {
		name = strings.Split(name, " AS ")[1]
	}

	switch val.(type) {
	default:
		modelPointer.FieldByName(strings.Title(strings.ToLower(name))).Set(reflect.ValueOf(val))
	case []uint8:
		var arr []int64

		err := json.Unmarshal([]byte(string(val.([]uint8))), &arr)

		if err != nil {
			log.Error(fmt.Sprintf("Error with casting type[%s] to array: %s", reflect.TypeOf(val), err))
		}

		modelPointer.FieldByName(strings.Title(strings.ToLower(name))).Set(reflect.ValueOf(arr))
	}
}
