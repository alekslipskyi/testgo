package tests

import (
	"../src"
	"./utils"
	"constants/requestError"
	"core/crypto"
	"core/db/connect"
	"encoding/json"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"models/User"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSpec(t *testing.T) {
	Convey("Auth tests", t, func() {
		srv := httptest.NewServer(router.Handler())
		prefix := srv.URL + "/api/v0/auth"

		createdUser := User.CreateAndFind(map[string]interface{}{
			"firstName": "string",
			"lastName":  "string",
			"password":  crypto.GenerateHash("string"),
			"username":  "string",
		})
		createdUser.GenerateToken()

		Convey("Login with right credentials should be successful and return a user object", func() {
			res, _ := http.Get(fmt.Sprintf("%s/token?username=%s&password=%s", prefix, "string", "string"))
			var responseBody map[string]interface{}
			_ = json.NewDecoder(res.Body).Decode(&responseBody)

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
			res, _ := http.Get(fmt.Sprintf("%s/token?username=%s&password=%s", prefix, "string", "string2"))

			var responseBody map[string]interface{}
			_ = json.NewDecoder(res.Body).Decode(&responseBody)

			So(res.StatusCode, ShouldEqual, http.StatusBadRequest)
			So(responseBody, ShouldResemble, utils.StructToMap(requestError.INVALID_CREDENTIAL))
		})

		Convey("Login with wrong email should be failed and return an error", func() {
			res, _ := http.Get(fmt.Sprintf("%s/token?username=%s&password=%s", prefix, "string2", "string"))

			var responseBody map[string]interface{}
			_ = json.NewDecoder(res.Body).Decode(&responseBody)

			So(res.StatusCode, ShouldEqual, http.StatusBadRequest)
			So(responseBody, ShouldResemble, utils.StructToMap(requestError.INVALID_CREDENTIAL))
		})

		connect.DB.Exec("delete from users")
	})
}
