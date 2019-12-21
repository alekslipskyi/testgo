package tests

import (
	"./utils"
	"constants/requestError"
	"core/db/connect"
	"core/logger"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"helpers/errors"
	"models/Channel"
	"models/ChannelUsers"
	"models/Message"
	"net/http"
	"testing"
)

var log = logger.Logger{Context: "Gettting messages tests"}

func TestGettingMessagesSpec(t *testing.T) {
	connect.Init()
	defer connect.DB.Close()

	Convey("Getting messages tests", t, func() {
		connect.DB.Exec("delete from users")

		requester := utils.Requester{}
		createdUser1 := utils.CreateUser()
		createdUser2 := utils.CreateUser("string2")

		channelID, err := Channel.Create(map[string]interface{}{
			"name": "test",
		})
		log.LogOnError(err, "error from creating channel", err)

		err = ChannelUsers.Create(map[string]interface{}{
			"user_id":     createdUser1.ID,
			"channel_id":  channelID,
			"permissions": "rwdui",
		})
		log.LogOnError(err, "error from assigning user to channel", err)

		requester.Init("/api/v0/message", map[string]interface{}{
			"auth": createdUser1.Token,
		})

		messageID, err := Message.CreateAndReturnID(channelID, createdUser1.ID, map[string]interface{}{
			"body": "test",
		})
		log.LogOnError(err, "error from creating message")

		Convey("Getting all messages relative to createdUser1", func() {
			res, responseBody := requester.GET(fmt.Sprintf("/%d/messages", channelID))

			So(res.StatusCode, ShouldEqual, http.StatusOK)

			expectedBody := []interface{}{map[string]interface{}{
				"body":       "test",
				"id":         float64(messageID),
				"channel_id": float64(channelID),
				"user": map[string]interface{}{
					"_id":      float64(createdUser1.ID),
					"username": createdUser1.Username,
				},
			}}

			So(responseBody, ShouldResemble, expectedBody)
		})

		Convey("Getting all messages with channel id belongs to not existed channel should return empty array", func() {
			res, responseBody := requester.GET(fmt.Sprintf("/%d/messages", 123))

			So(res.StatusCode, ShouldEqual, http.StatusOK)

			So(responseBody, ShouldResemble, []interface{}{})
		})

		Convey("Getting all messages with user id which have't channels should return empty array", func() {
			requester.UnsetAuth()
			requester.SetAuth(createdUser2.Token)
			res, responseBody := requester.GET(fmt.Sprintf("/%d/messages", channelID))

			So(res.StatusCode, ShouldEqual, http.StatusOK)

			So(responseBody, ShouldResemble, []interface{}{})
		})

		Convey("Getting all messages with wrong channel id should return error \"channelID must be a number\"", func() {
			res, responseBody := requester.GET(fmt.Sprintf("/%s/messages", "sadsa"))

			So(res.StatusCode, ShouldEqual, http.StatusBadRequest)

			expectedError := utils.StructToMap(errors.IRequestError{
				StatusCode: http.StatusBadRequest,
				Message:    "channelID must be a number",
				Token:      "NOT_VALID",
			})

			So(responseBody, ShouldResemble, expectedError)
		})

		Convey("Getting all messages without being authenticated should return UNAUTHORIZED error", func() {
			requester.UnsetAuth()
			res, responseBody := requester.GET(fmt.Sprintf("/%d/messages", channelID))

			So(res.StatusCode, ShouldEqual, http.StatusUnauthorized)
			So(responseBody, ShouldResemble, utils.StructToMap(requestError.UNAUTHORIZED))
		})

		connect.DB.Exec("delete from channels")
	})
}
