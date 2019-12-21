package tests

import (
	"./utils"
	"constants/requestError"
	"core/db/connect"
	"core/db/types"
	. "github.com/smartystreets/goconvey/convey"
	"helpers/errors"
	"models/Message"
	"net/http"
	"testing"
)

func TestCreatingMessageSpec(t *testing.T) {
	connect.Init()
	defer connect.DB.Close()

	Convey("Creating message tests", t, func() {
		connect.DB.Exec("delete from users")

		createdUser1 := utils.CreateUser()
		channelID := utils.CreateChannel(map[string]interface{}{
			"name": "tset",
		}, createdUser1.ID, "rwdui")

		requester := utils.InitRequester("/api/v0/message/", map[string]interface{}{
			"auth": createdUser1.Token,
		})

		Convey("Creating message should create a message and return message \"ok\"", func() {
			res, _ := requester.POST(channelID, map[string]interface{}{
				"body": "hello",
			})

			So(res.StatusCode, ShouldEqual, http.StatusCreated)

			message := Message.FindOne(types.Where{
				"body": "hello",
			})

			So(message.IsExist(), ShouldBeTrue)
		})

		Convey("Creating message in channel which is not exist should return NOT_FOUND error", func() {
			res, responseBody := requester.POST(123, map[string]interface{}{
				"body": "hello",
			})

			So(res.StatusCode, ShouldEqual, http.StatusNotFound)
			So(responseBody, ShouldResemble, utils.StructToMap(requestError.NOT_FOUND))

			message := Message.FindOne(types.Where{
				"body": "hello",
			})

			So(message.IsExist(), ShouldBeFalse)
		})

		Convey("Creating message with wrong channel id should return error message \"channelID must be a number\"", func() {
			res, responseBody := requester.POST("hello", map[string]interface{}{
				"body": "hello",
			})

			So(res.StatusCode, ShouldEqual, http.StatusBadRequest)

			expectedError := utils.StructToMap(errors.IRequestError{
				StatusCode: http.StatusBadRequest,
				Message:    "channelID must be a number",
				Token:      "NOT_VALID",
			})

			So(responseBody, ShouldResemble, expectedError)

			message := Message.FindOne(types.Where{
				"body": "hello",
			})

			So(message.IsExist(), ShouldBeFalse)
		})

		Convey("Creating message with wrong body return error message \"body must be a string\"", func() {
			res, responseBody := requester.POST(channelID, map[string]interface{}{
				"body": 123,
			})

			So(res.StatusCode, ShouldEqual, http.StatusBadRequest)

			expectedError := utils.StructToMap(errors.IRequestError{
				StatusCode: http.StatusBadRequest,
				Message:    "body must be a string",
				Token:      "NOT_VALID",
			})

			So(responseBody, ShouldResemble, expectedError)

			message := Message.FindOne(types.Where{
				"body": "hello",
			})

			So(message.IsExist(), ShouldBeFalse)
		})

		connect.DB.Exec("delete from messages")
		connect.DB.Exec("delete from channels")
	})
}
