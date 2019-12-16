package tests

import (
	"./utils"
	"constants/requestError"
	"core/crypto"
	"core/db/connect"
	. "github.com/smartystreets/goconvey/convey"
	"models/User"
	"net/http"
	"testing"
)

func TestGettingMeSpec(t *testing.T) {
	Convey("Test getting me spec", t, func() {
		connect.DB.Exec("delete from users")

		createdUser := User.CreateAndFind(map[string]interface{}{
			"firstName": "string",
			"lastName":  "string",
			"password":  crypto.GenerateHash("string"),
			"username":  "string",
		})
		createdUser.AddAllowIP("127.0.0.1")
		createdUser.GenerateToken()

		requester := utils.Requester{}
		requester.Init("/api/v0/user", map[string]interface{}{})

		Convey("Getting me should return my user data", func() {

			requester.SetAuth(createdUser.Token)
			res, responseBody := requester.GET()

			expectedBody := map[string]interface{}{
				"id":         float64(createdUser.ID),
				"first_name": createdUser.Firstname,
				"last_name":  createdUser.Lastname,
				"username":   createdUser.Username,
			}

			So(res.StatusCode, ShouldEqual, http.StatusOK)
			So(responseBody, ShouldResemble, expectedBody)
		})

		Convey("Getting me without auth header should return UNAUTHORIZED error", func() {

			res, responseBody := requester.GET()

			So(res.StatusCode, ShouldEqual, http.StatusUnauthorized)
			So(responseBody, ShouldResemble, utils.StructToMap(requestError.UNAUTHORIZED))
		})

		Convey("Getting me with auth header but unauthorized ip should return WRONG_IP error", func() {

			requester.SetHeader("Authorization", createdUser.Token)
			requester.SetHeader("X-Real-IP", "127.0.0.2")
			res, responseBody := requester.GET()

			So(res.StatusCode, ShouldEqual, http.StatusUnauthorized)
			So(responseBody, ShouldResemble, utils.StructToMap(requestError.WRONG_IP))
		})

		Convey("Getting me with fake auth header should return UNAUTHORIZED error", func() {

			requester.SetHeader("Authorization", "fake")
			requester.SetHeader("X-Real-IP", "127.0.0.2")
			res, responseBody := requester.GET()

			So(res.StatusCode, ShouldEqual, http.StatusUnauthorized)
			So(responseBody, ShouldResemble, utils.StructToMap(requestError.UNAUTHORIZED))
		})

		Reset(func() {
			requester.UnsetAuth()
		})
	})
}
