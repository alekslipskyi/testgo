package validation

import (
	"core/Router"
	"core/validation/constants"
	"core/validation/types"
	"reflect"
)

type (
	MustBe      map[string]interface{}
	MustBeOneOf map[string]interface{}
)

func getFieldFromCTX(in string, ctx *Router.Context) (bool, map[string]interface{}) {
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

func IsValid(in string, schema interface{}) func(ctx *Router.Context) (bool, interface{}, interface{}) {
	return func(ctx *Router.Context) (bool, interface{}, interface{}) {
		ok, data := getFieldFromCTX(in, ctx)
		var schemaBody map[string]interface{}

		isRequiredDefault := reflect.TypeOf(schema) == reflect.TypeOf(MustBe{})

		if reflect.TypeOf(schema) == reflect.TypeOf(MustBeOneOf{}) {
			schemaBody = schema.(MustBeOneOf)
			if len(data) == 0 {
				return false, constants.OneOfShouldBe, nil
			}
		} else {
			schemaBody = schema.(MustBe)
		}

		if ok {
			for key := range data {
				if schemaBody[key] == nil {
					return false, key + constants.IsNotAllowed, nil
				}
			}

			for key, field := range schemaBody {
				switch reflect.TypeOf(field) {
				case reflect.TypeOf(types.Number{}):
					{
						number := types.Number{}
						ok, errMessage := number.Validate(field, key, data[key], isRequiredDefault)
						if !ok {
							return ok, errMessage, nil
						}
					}
				case reflect.TypeOf(types.String{}):
					{
						str := types.String{}
						ok, errMessage := str.Validate(field, key, data[key], isRequiredDefault)

						if !ok {
							return ok, errMessage, nil
						}
					}
				}
			}
		}

		return true, "ok", nil
	}
}
