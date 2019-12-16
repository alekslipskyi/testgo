package tests

import (
	"./utils"
	"constants/requestError"
	channelController "controllers/channel"
	"core/db/connect"
	"core/db/types"
	"core/logger"
	. "github.com/smartystreets/goconvey/convey"
	"models/Channel"
	"net/http"
	"testing"
)

func TestDeleteChannelSpec(t *testing.T) {
	var log = logger.Logger{Context: "Delete channel tests", Colors: logger.Colors{Info: logger.GREEN}}

	Convey("Delete channel tests", t, func() {
		_, err := connect.DB.Exec("delete from users")
		log.LogOnError(err)

		createdUser := utils.CreateUser()
		createdUser2 := utils.CreateUser("string2")

		requester := utils.Requester{}
		requester.Init("/api/v0/channel/", map[string]interface{}{
			"auth": createdUser.Token,
		})

		Convey("Delete created channel should delete channel and return message\"ok\"", func() {
			log.Info("Delete created channel should delete channel and return message\"ok\"")
			channelID := utils.CreateChannel(map[string]interface{}{
				"name": "test",
			}, createdUser.ID, "rwd")
			requester.Debug()
			res, _ := requester.DELETE(channelID)

			So(res.StatusCode, ShouldEqual, http.StatusOK)

			channel := Channel.FindOnlyChannel(types.Where{"name": "test"})
			So(channel.IsExist(), ShouldBeFalse)
		})

		Convey("Delete created channel with params as string should return NOT_FOUND error", func() {
			log.Info("Delete created channel with params as string should return NOT_FOUND error")
			utils.CreateChannel(map[string]interface{}{
				"name": "test",
			}, createdUser.ID, "rwd")
			res, responseBody := requester.DELETE("test")

			So(res.StatusCode, ShouldEqual, http.StatusNotFound)
			So(responseBody, ShouldResemble, utils.StructToMap(requestError.NOT_FOUND))

			channel := Channel.FindOnlyChannel(types.Where{"name": "test"})
			So(channel.IsExist(), ShouldBeFalse)
		})

		Convey("Delete not existed channel should return NOT_FOUND error", func() {
			log.Debug("1")
			log.Info("Delete not existed channel should return NOT_FOUND error")
			utils.CreateChannel(map[string]interface{}{
				"name": "test",
			}, createdUser.ID, "rwd")
			res, responseBody := requester.DELETE(123)
			log.Debug("2")

			So(res.StatusCode, ShouldEqual, http.StatusNotFound)
			So(responseBody, ShouldResemble, utils.StructToMap(requestError.NOT_FOUND))
			log.Debug("3")

			channel := Channel.FindOnlyChannel(types.Where{"name": "test"})
			So(channel.IsExist(), ShouldBeFalse)
		})

		Convey("Delete created channel by another user should return NOT_FOUND error", func() {
			log.Info("Delete created channel by another user should return NOT_FOUND error")
			channelID := utils.CreateChannel(map[string]interface{}{
				"name": "test",
			}, createdUser2.ID, "rwd")
			res, responseBody := requester.DELETE(channelID)

			So(res.StatusCode, ShouldEqual, http.StatusNotFound)
			So(responseBody, ShouldResemble, utils.StructToMap(requestError.NOT_FOUND))

			channel := Channel.FindOnlyChannel(types.Where{"name": "test"})
			So(channel.IsExist(), ShouldBeFalse)
		})

		Convey("Delete created channel without such permissions should return NOT_ALLOWED error", func() {
			log.Info("Delete created channel without such permissions should return NOT_ALLOWED error")
			channelID := utils.CreateChannel(map[string]interface{}{
				"name": "test",
			}, createdUser.ID, "rw")
			res, responseBody := requester.DELETE(channelID)

			So(res.StatusCode, ShouldEqual, http.StatusForbidden)
			So(responseBody, ShouldResemble, utils.StructToMap(channelController.NOT_ALLOWED_TO_DROP))

			channel := Channel.FindOnlyChannel(types.Where{"name": "test"})
			So(channel.IsExist(), ShouldBeFalse)
		})

		Convey("Delete created channel without auth header should return UNAUTHORIZED error", func() {
			log.Info("Delete created channel without auth header should return UNAUTHORIZED error")
			requester.UnsetAuth()
			utils.CreateChannel(map[string]interface{}{
				"name": "test",
			}, createdUser.ID, "rwd")
			res, responseBody := requester.DELETE("test")

			So(res.StatusCode, ShouldEqual, http.StatusUnauthorized)
			So(responseBody, ShouldResemble, utils.StructToMap(requestError.UNAUTHORIZED))

			channel := Channel.FindOnlyChannel(types.Where{"name": "test"})
			So(channel.IsExist(), ShouldBeFalse)
		})

		Reset(func() {
			connect.DB.Exec("delete from channels")
		})

	})
}
