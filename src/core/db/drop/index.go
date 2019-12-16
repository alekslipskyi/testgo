package drop

import (
	"core/db/connect"
	"core/db/converter"
	"core/db/types"
	"core/logger"
	"fmt"
)

type SDrop struct {
	Name string
}

var log = logger.Logger{Context: "DROP QUERY"}

func (entity *SDrop) Drop(options types.QueryOptions) bool {
	query := fmt.Sprintf("delete from %s where %s", entity.Name, converter.DataToQueryString(options.Where, "=", "and"))

	_, err := connect.DB.Query(query)

	log.Debug("query", query)

	if err != nil {
		log.Error(err)

		return false
	}

	return true
}
