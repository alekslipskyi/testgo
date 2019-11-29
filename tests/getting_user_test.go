package tests

import (
	"../src"
	"./utils"
	"core/crypto"
	"core/db/connect"
	"encoding/json"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"helpers/errors"
	"models/User"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGettingUserSpec(t *testing.T) {
	Convey("Test getting user spec", t, func() {
		srv := httptest.NewServer(router.Handler())
		url := srv.URL + "/api/v0/user/"

		createdUser := User.CreateAndFind(map[string]interface{}{
			"firstName": "string",
			"lastName":  "string",
			"password":  crypto.GenerateHash("string"),
			"username":  "string",
		})
		createdUser.AddAllowIP("127.0.0.1")
		createdUser.GenerateToken()

		Convey("Getting user should return user", func() {
			client := &http.Client{}
			req, _ := http.NewRequest("GET", fmt.Sprintf("%s%d", url, createdUser.ID), nil)
			res, _ := client.Do(req)

			var responseBody map[string]interface{}
			_ = json.NewDecoder(res.Body).Decode(&responseBody)

			expectedBody := map[string]interface{}{
				"id":         float64(createdUser.ID),
				"first_name": createdUser.Firstname,
				"last_name":  createdUser.Lastname,
				"username":   createdUser.Username,
			}

			fmt.Printf("tes--- %s", responseBody)

			So(res.StatusCode, ShouldEqual, http.StatusOK)
			So(responseBody, ShouldResemble, expectedBody)
		})

		Convey("Getting user with param id as string should fail", func() {
			client := &http.Client{}
			req, _ := http.NewRequest("GET", fmt.Sprintf("%s%s", url, "test"), nil)
			res, _ := client.Do(req)

			var responseBody map[string]interface{}
			_ = json.NewDecoder(res.Body).Decode(&responseBody)

			expectedError := utils.StructToMap(errors.IRequestError{http.StatusBadRequest, "userID must be a number", "NOT_VALID"})

			So(res.StatusCode, ShouldEqual, http.StatusBadRequest)
			So(responseBody, ShouldResemble, expectedError)
		})

		connect.DB.Exec("delete from users")
	})
}
