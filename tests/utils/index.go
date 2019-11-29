package utils

import "encoding/json"

func StructToMap(st interface{}) map[string]interface{} {
	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(st)
	json.Unmarshal(inrec, &inInterface)

	return inInterface
}
