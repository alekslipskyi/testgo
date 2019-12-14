package tests

import (
	"./utils"
	"constants/requestError"
	channelController "controllers/channel"
	"core/db/connect"
	"core/db/types"
	. "github.com/smartystreets/goconvey/convey"
	"models/Channel"
	"net/http"
	"testing"
)

func TestDeleteChannelSpec(t *testing.T) {
	Convey("Delete channel tests", t, func() {
		connect.DB.Exec("delete from users")

		createdUser := utils.CreateUser()
		createdUser2 := utils.CreateUser()

		requester := utils.Requester{}
		requester.Init("/api/v0/channel/", map[string]interface{}{
			"auth": createdUser.Token,
		})

		Convey("Delete created channel should delete channel and return message\"ok\"", func() {
			channelID := utils.CreateChannel(map[string]interface{}{
				"name": "test",
			}, createdUser.ID, "rwd")
			res, _ := requester.DELETE(channelID)

			So(res.StatusCode, ShouldEqual, http.StatusOK)

			channel := Channel.FindOne(types.QueryOptions{Where: types.Where{"name": "test"}})
			So(channel.IsExist(), ShouldBeFalse)
		})

		Convey("Delete created channel with params as string should return NOT_FOUND error", func() {
			utils.CreateChannel(map[string]interface{}{
				"name": "test",
			}, createdUser.ID, "rwd")
			res, responseBody := requester.DELETE("test")

			So(res.StatusCode, ShouldEqual, http.StatusNotFound)
			So(responseBody, ShouldResemble, utils.StructToMap(requestError.NOT_FOUND))

			channel := Channel.FindOne(types.QueryOptions{Where: types.Where{"name": "test"}})
			So(channel.IsExist(), ShouldBeFalse)
		})

		Convey("Delete not existed channel should return NOT_FOUND error", func() {
			utils.CreateChannel(map[string]interface{}{
				"name": "test",
			}, createdUser.ID, "rwd")
			res, responseBody := requester.DELETE(123)

			So(res.StatusCode, ShouldEqual, http.StatusNotFound)
			So(responseBody, ShouldResemble, utils.StructToMap(requestError.NOT_FOUND))

			channel := Channel.FindOne(types.QueryOptions{Where: types.Where{"name": "test"}})
			So(channel.IsExist(), ShouldBeFalse)
		})

		Convey("Delete created channel by another user should return NOT_FOUND error", func() {
			channelID := utils.CreateChannel(map[string]interface{}{
				"name": "test",
			}, createdUser2.ID, "rwd")
			res, responseBody := requester.DELETE(channelID)

			So(res.StatusCode, ShouldEqual, http.StatusNotFound)
			So(responseBody, ShouldResemble, utils.StructToMap(requestError.NOT_FOUND))

			channel := Channel.FindOne(types.QueryOptions{Where: types.Where{"name": "test"}})
			So(channel.IsExist(), ShouldBeFalse)
		})

		Convey("Delete created channel without such permissions should return NOT_ALLOWED error", func() {
			channelID := utils.CreateChannel(map[string]interface{}{
				"name": "test",
			}, createdUser.ID, "rw")
			res, responseBody := requester.DELETE(channelID)

			So(res.StatusCode, ShouldEqual, http.StatusForbidden)
			So(responseBody, ShouldResemble, utils.StructToMap(channelController.NOT_ALLOWED_TO_DROP))

			channel := Channel.FindOne(types.QueryOptions{Where: types.Where{"name": "test"}})
			So(channel.IsExist(), ShouldBeFalse)
		})

		Convey("Delete created channel without auth header should return UNAUTHORIZED error", func() {
			requester.UnsetAuth()
			utils.CreateChannel(map[string]interface{}{
				"name": "test",
			}, createdUser.ID, "rwd")
			res, responseBody := requester.DELETE("test")

			So(res.StatusCode, ShouldEqual, http.StatusUnauthorized)
			So(responseBody, ShouldResemble, utils.StructToMap(requestError.UNAUTHORIZED))

			channel := Channel.FindOne(types.QueryOptions{Where: types.Where{"name": "test"}})
			So(channel.IsExist(), ShouldBeFalse)
		})

		Reset(func() {
			connect.DB.Exec("delete from channels")
		})

	})
}
