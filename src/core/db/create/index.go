package create

import (
	"core/db/connect"
	"fmt"
	"strconv"
	"strings"
)

type SCreate struct {
	Name  string
	Model interface{}
}

func (entity *SCreate) Create(data map[string]interface{}) interface{} {
	values, keys := entity.parseData(data)

	fmt.Println(strings.Join(keys, ","))

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		entity.Name,
		strings.Join(keys, ","),
		strings.Join(values, ","))

	_, err := connect.DB.Exec(query)

	if err != nil {
		return err
	}

	return nil
}

func (entity *SCreate) parseData(data map[string]interface{}) ([]string, []string) {
	var values []string
	var keys []string

	for key, value := range data {
		switch value.(type) {
		case float64:
			values = append(values, strconv.FormatFloat(value.(float64), 'f', 0, 64))
		default:
			values = append(values, fmt.Sprintf("'%s'", strings.ToLower(value.(string))))
		}

		keys = append(keys, strings.ToLower(key))
	}

	return values, keys
}
