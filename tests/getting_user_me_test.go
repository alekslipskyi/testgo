package tests

import (
	"../src"
	"./utils"
	"constants/requestError"
	"core/crypto"
	"core/db/connect"
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"models/User"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGettingMeSpec(t *testing.T) {
	Convey("Test getting me spec", t, func() {
		srv := httptest.NewServer(router.Handler())
		url := srv.URL + "/api/v0/user"

		createdUser := User.CreateAndFind(map[string]interface{}{
			"firstName": "string",
			"lastName":  "string",
			"password":  crypto.GenerateHash("string"),
			"username":  "string",
		})
		createdUser.AddAllowIP("127.0.0.1")
		createdUser.GenerateToken()

		Convey("Getting me should return my user data", func() {
			client := &http.Client{}
			req, _ := http.NewRequest("GET", url, nil)
			req.Header.Set("Authorization", createdUser.Token)
			req.Header.Set("X-Real-IP", "127.0.0.1")
			res, _ := client.Do(req)

			var responseBody map[string]interface{}
			_ = json.NewDecoder(res.Body).Decode(&responseBody)

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
			client := &http.Client{}
			req, _ := http.NewRequest("GET", url, nil)
			req.Header.Set("X-Real-IP", "127.0.0.1")
			res, _ := client.Do(req)

			var responseBody map[string]interface{}
			_ = json.NewDecoder(res.Body).Decode(&responseBody)

			So(res.StatusCode, ShouldEqual, http.StatusUnauthorized)
			So(responseBody, ShouldResemble, utils.StructToMap(requestError.UNAUTHORIZED))
		})

		Convey("Getting me with auth header but unauthorized ip should return WRONG_IP error", func() {
			client := &http.Client{}
			req, _ := http.NewRequest("GET", url, nil)
			req.Header.Set("Authorization", createdUser.Token)
			req.Header.Set("X-Real-IP", "127.0.0.2")
			res, _ := client.Do(req)

			var responseBody map[string]interface{}
			_ = json.NewDecoder(res.Body).Decode(&responseBody)

			So(res.StatusCode, ShouldEqual, http.StatusUnauthorized)
			So(responseBody, ShouldResemble, utils.StructToMap(requestError.WRONG_IP))
		})

		Convey("Getting me with fake auth header should return UNAUTHORIZED error", func() {
			client := &http.Client{}
			req, _ := http.NewRequest("GET", url, nil)
			req.Header.Set("Authorization", "fake")
			req.Header.Set("X-Real-IP", "127.0.0.1")
			res, _ := client.Do(req)

			var responseBody map[string]interface{}
			_ = json.NewDecoder(res.Body).Decode(&responseBody)

			So(res.StatusCode, ShouldEqual, http.StatusUnauthorized)
			So(responseBody, ShouldResemble, utils.StructToMap(requestError.UNAUTHORIZED))
		})

		connect.DB.Exec("delete from users")
	})
}
