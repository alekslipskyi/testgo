package tests

import (
	"./utils"
	"controllers/channel"
	"core/db/connect"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"helpers/errors"
	"net/http"
	"testing"
)

func TestInviteToChannelSpec(t *testing.T) {
	Convey("Test invite to channel", t, func() {
		connect.DB.Exec("delete from users")

		createdUser := utils.CreateUser()
		createdUser2 := utils.CreateUser("test2")

		requester := utils.Requester{}
		requester.Init("/api/v0/channel/", map[string]interface{}{
			"auth": createdUser.Token,
		})

		Convey("Invite user to channel should be successful", func() {


			channelID := utils.CreateChannel(map[string]interface{}{
				"name": "test",
			}, createdUser.ID, "rwdui")
			res, _ := requester.PUT(fmt.Sprintf("%d/invite/%d", createdUser2.ID, int64(channelID)))

			So(res.StatusCode, ShouldEqual, http.StatusOK)
		})

		Convey("Invite user to channel without needed permission should be failed with error BAD_PERMISSION_INVITE", func() {

			channelID := utils.CreateChannel(map[string]interface{}{
				"name": "test",
			}, createdUser.ID, "rwdu")
			res, responseBody := requester.PUT(fmt.Sprintf("%d/invite/%d", createdUser2.ID, int64(channelID)))

			So(res.StatusCode, ShouldEqual, http.StatusForbidden)
			So(responseBody, ShouldResemble, utils.StructToMap(channel.NOT_ALLOWED_TO_INVITE))
		})

		Convey("Invite user to channel with wrong userID should be failed and return error\"userID must be a number\"", func() {

			res, responseBody := requester.PUT(fmt.Sprintf("%s/invite/%d", "test", 12))

			So(res.StatusCode, ShouldEqual, http.StatusBadRequest)

			expectedError := utils.StructToMap(errors.IRequestError{
				StatusCode: http.StatusBadRequest,
				Message:    "userID must be a number",
				Token:      "NOT_VALID",
			})

			So(responseBody, ShouldResemble, expectedError)
		})

		Reset(func() {
			connect.DB.Exec("delete from channels")
		})
	})
}
