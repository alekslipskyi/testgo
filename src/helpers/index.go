package helpers

import (
	"constants/requestError"
	"core/Router"
	dbTypes "core/db/types"
	"core/logger"
	"models/ChannelUsers"
	"reflect"
)

var log = logger.Logger{Context:"Helpers"}

func OmitPrivateFields(obj interface{}) {
	svPTR := reflect.ValueOf(obj).Elem()
	sv := reflect.Indirect(svPTR)
	st := sv.Type()

	for i := 0; i < sv.NumField(); i++ {
		key := st.Field(i)

		access := key.Tag.Get("json")
		if access == "private" {
			fieldValue := sv.Field(i)

			fieldValue.Set(reflect.Zero(fieldValue.Type()))
		}
	}
}

func HasPermissions(permission rune, error interface{}) func(ctx *Router.Context) (bool, interface{}, interface{}) {
	return func(ctx *Router.Context) (bool, interface{}, interface{}) {
		channel := ChannelUsers.Find(dbTypes.QueryOptions{
			Where: map[string]interface{}{"channel_id": ctx.Params["channelID"].(int64), "user_id": ctx.User.ID},
		})

		if !channel.IsExist() {
			return false, nil, requestError.NOT_FOUND
		}

		if !channel.HasPermission(permission) {
			return false, nil, error
		}

		return true, "ok", nil
	}
}
