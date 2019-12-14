package tests

import (
	"./utils"
	"constants/requestError"
	"core/db/connect"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"testing"
)

func TestGettingChannels(t *testing.T) {
	Convey("Test getting channels", t, func() {
		connect.DB.Exec("delete from users")

		createdUser := utils.CreateUser()
		createdUser2 := utils.CreateUser()

		requester := utils.Requester{}
		requester.Init("/api/v0/channel", map[string]interface{}{
			"auth": createdUser.Token,
		})

		Convey("Getting all channels when channels is not created should return empty array", func() {
			res, responseBody := requester.GET()

			So(res.StatusCode, ShouldEqual, http.StatusOK)
			So(responseBody, ShouldResemble, []interface{}{})
		})

		Convey("Getting all channels related to User should return list of one channel", func() {
			utils.CreateChannel(map[string]interface{}{
				"name": "test2",
			}, createdUser2.ID)
			channelID := utils.CreateChannel(map[string]interface{}{
				"name": "test",
			}, createdUser.ID)

			res, responseBody := requester.GET()

			expectedBody := []interface{}{map[string]interface{}{
				"id":    channelID,
				"name":  "test",
				"users": []interface{}{float64(createdUser.ID)},
			}}

			So(res.StatusCode, ShouldEqual, http.StatusOK)
			So(responseBody, ShouldResemble, expectedBody)
		})

		Convey("Getting all channels when two channel is created should return list of two channels", func() {
			channelID := utils.CreateChannel(map[string]interface{}{
				"name": "test",
			}, createdUser.ID)

			channelID2 := utils.CreateChannel(map[string]interface{}{
				"name": "test2",
			}, createdUser.ID)

			res, responseBody := requester.GET()

			expectedBody := []interface{}{map[string]interface{}{
				"id":    channelID,
				"name":  "test",
				"users": []interface{}{float64(createdUser.ID)},
			}, map[string]interface{}{
				"id":    channelID2,
				"name":  "test2",
				"users": []interface{}{float64(createdUser.ID)},
			}}

			So(res.StatusCode, ShouldEqual, http.StatusOK)
			So(responseBody, ShouldResemble, expectedBody)
		})

		Convey("Getting all channels relat when one channel is created should return list of one channel", func() {
			channelID := utils.CreateChannel(map[string]interface{}{
				"name": "test",
			}, createdUser.ID)

			res, responseBody := requester.GET()

			expectedBody := []interface{}{map[string]interface{}{
				"id":    channelID,
				"name":  "test",
				"users": []interface{}{float64(createdUser.ID)},
			}}

			So(res.StatusCode, ShouldEqual, http.StatusOK)
			So(responseBody, ShouldResemble, expectedBody)
		})

		Convey("Getting all channels when token is not provided should return UNAUTHORIZED error", func() {
			requester.UnsetAuth()
			res, responseBody := requester.GET()

			So(res.StatusCode, ShouldEqual, http.StatusUnauthorized)
			So(responseBody, ShouldResemble, utils.StructToMap(requestError.UNAUTHORIZED))
		})

		Reset(func() {
			connect.DB.Exec("delete from channels")
		})
	})
}
