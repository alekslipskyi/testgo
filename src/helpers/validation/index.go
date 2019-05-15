package validation

import (
	"helpers/validation/constants"
	"helpers/validation/types"
	"lib/Router"
	"reflect"
)

type (
	MustBe map[string]interface{}
)

func getFieldFromCTX(in string, ctx Router.Context) (bool, map[string]interface{}) {
	defaultParams := make(map[string]interface{})

	switch in {
	case "body":
		return true, ctx.Body
	case "params":
		return true, ctx.Params
	default:
		return false, defaultParams
	}
}

func IsValid(in string, validBody MustBe) func(ctx Router.Context) (bool, string) {
	return func(ctx Router.Context) (bool, string) {
		ok, data := getFieldFromCTX(in, ctx)

		if ok {
			for key := range data {
				if validBody[key] == nil {
					return false, key + constants.IsNotAllowed
				}
			}

			for key, field := range validBody {
				switch reflect.TypeOf(field) {
				case reflect.TypeOf(types.Number{}):
					{
						number := types.Number{}
						ok, errMessage := number.Validate(field, key, data[key])
						if !ok {
							return ok, errMessage
						}
					}
				case reflect.TypeOf(types.String{}):
					{
						str := types.String{}
						ok, errMessage := str.Validate(field, key, data[key])

						if !ok {
							return ok, errMessage
						}
					}
				}
			}
		}

		return true, "ok"
	}
}
