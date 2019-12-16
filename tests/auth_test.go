package tests

import (
	"./utils"
	"constants/requestError"
	"core/crypto"
	"core/db/connect"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"models/User"
	"net/http"
	"testing"
)

func TestSpec(t *testing.T) {
	connect.Init()
	Convey("Auth tests", t, func() {
		connect.DB.Exec("delete from users")
		requester := utils.Requester{}
		requester.Init("/api/v0/auth", map[string]interface{}{})

		createdUser := User.CreateAndFind(map[string]interface{}{
			"firstName": "string",
			"lastName":  "string",
			"password":  crypto.GenerateHash("string"),
			"username":  "string",
		})
		createdUser.GenerateToken()

		Convey("Login with right credentials should be successful and return a user object", func() {


			res, responseBody := requester.GET(fmt.Sprintf("/token?username=%s&password=%s", "string", "string"))

			expectedBody := map[string]interface{}{
				"id":         float64(createdUser.ID),
				"first_name": createdUser.Firstname,
				"last_name":  createdUser.Lastname,
				"username":   createdUser.Username,
				"token":      createdUser.Token,
			}

			So(res.StatusCode, ShouldEqual, http.StatusOK)
			So(responseBody, ShouldResemble, expectedBody)
		})

		Convey("Login with wrong password should be failed and return an error", func() {

			res, responseBody := requester.GET(fmt.Sprintf("/token?username=%s&password=%s", "string", "string2"))

			So(res.StatusCode, ShouldEqual, http.StatusBadRequest)
			So(responseBody, ShouldResemble, utils.StructToMap(requestError.INVALID_CREDENTIAL))
		})

		Convey("Login with wrong email should be failed and return an error", func() {

			res, responseBody := requester.GET(fmt.Sprintf("/token?username=%s&password=%s", "string2", "string"))

			So(res.StatusCode, ShouldEqual, http.StatusBadRequest)
			So(responseBody, ShouldResemble, utils.StructToMap(requestError.INVALID_CREDENTIAL))
		})
	})
}
