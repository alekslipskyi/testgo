package helpers

import (
	channel2 "controllers/channel"
	"core/Router"
	dbTypes "core/db/types"
	"models/ChannelUsers"
	"reflect"
)

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

func HasPermissions(permission rune, error interface{}) func(ctx Router.Context) (bool, interface{}, interface{}) {
	return func(ctx Router.Context) (bool, interface{}, interface{}) {
		channel := ChannelUsers.Find(dbTypes.QueryOptions{
			Where: map[string]interface{}{"channel_id": ctx.Params["ChannelID"].(int64), "user_id": ctx.User.ID},
		})

		if !channel.IsExist() {
			return false, channel2.CHANNEL_NOT_FOUND, nil
		}

		if !channel.HasPermission(permission) {
			return false, error, nil
		}

		return true, "ok", nil
	}
}
