package tests

import (
	"./utils"
	"core/crypto"
	"core/db/connect"
	"core/logger"
	. "github.com/smartystreets/goconvey/convey"
	"helpers/errors"
	"models/User"
	"net/http"
	"testing"
)

func TestGettingUserSpec(t *testing.T) {
	Convey("Test getting user spec", t, func() {
		var log = logger.Logger{Context: "getting user tests", Colors: logger.Colors{Info: logger.GREEN}}
		connect.DB.Exec("delete from users")

		requester := utils.Requester{}
		requester.Init("/api/v0/user/", map[string]interface{}{})

		createdUser := User.CreateAndFind(map[string]interface{}{
			"firstName": "string",
			"lastName":  "string",
			"password":  crypto.GenerateHash("string"),
			"username":  "string",
		})
		createdUser.AddAllowIP("127.0.0.1")
		createdUser.GenerateToken()

		Convey("Getting user should return user", func() {
			log.Info("Getting user should return user")
			res, responseBody := requester.GET(createdUser.ID)

			expectedBody := map[string]interface{}{
				"id":         float64(createdUser.ID),
				"first_name": createdUser.Firstname,
				"last_name":  createdUser.Lastname,
				"username":   createdUser.Username,
			}

			So(res.StatusCode, ShouldEqual, http.StatusOK)
			So(responseBody, ShouldResemble, expectedBody)
		})

		Convey("Getting user with param id as string should fail", func() {
			log.Info("Getting user with param id as string should fail")
			res, responseBody := requester.GET("test")

			expectedError := utils.StructToMap(errors.IRequestError{http.StatusBadRequest, "userID must be a number", "NOT_VALID"})

			So(res.StatusCode, ShouldEqual, http.StatusBadRequest)
			So(responseBody, ShouldResemble, expectedError)
		})
	})
}
