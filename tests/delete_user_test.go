package tests

import (
	"../src"
	"./utils"
	"constants/requestError"
	"core/crypto"
	"core/db/connect"
	"core/db/types"
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"models/User"
	"net/http"
	"net/http/httptest"
	"testing"
)

func createUser() User.User {
	createdUser := User.CreateAndFind(map[string]interface{}{
		"firstName": "string",
		"lastName":  "string",
		"password":  crypto.GenerateHash("string"),
		"username":  "string",
	})
	createdUser.GenerateToken()
	createdUser.AddAllowIP("127.0.0.1")

	return createdUser
}

func TestDeleteUserSpec(t *testing.T) {
	Convey("Delete user Test", t, func() {
		srv := httptest.NewServer(router.Handler())
		url := srv.URL + "/api/v0/user"

		Convey("Delete user should be successful and return ok", func() {
			createdUser := createUser()

			client := &http.Client{}
			req, _ := http.NewRequest("DELETE", url, nil)
			req.Header.Set("Authorization", createdUser.Token)
			req.Header.Set("X-Real-IP", "127.0.0.1")
			res, _ := client.Do(req)

			user := User.Find(types.QueryOptions{
				Where: types.Where{"username": "string"},
			})

			So(res.StatusCode, ShouldEqual, http.StatusOK)
			So(user.IsNotExist(), ShouldBeTrue)
		})

		Convey("Delete user without providing auth header should be failed and return UNAUTHORIZED error", func() {
			createUser()

			client := &http.Client{}
			req, _ := http.NewRequest("DELETE", url, nil)
			req.Header.Set("X-Real-IP", "127.0.0.1")
			res, _ := client.Do(req)

			user := User.Find(types.QueryOptions{
				Where: types.Where{"username": "string"},
			})

			var responseBody map[string]interface{}
			_ = json.NewDecoder(res.Body).Decode(&responseBody)

			So(res.StatusCode, ShouldEqual, http.StatusUnauthorized)
			So(!user.IsNotExist(), ShouldBeTrue)
			So(responseBody, ShouldResemble, utils.StructToMap(requestError.UNAUTHORIZED))
		})

		connect.DB.Exec("delete from users")
	})
}
