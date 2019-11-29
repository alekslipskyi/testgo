package tests

import (
	"../src"
	"./utils"
	"bytes"
	"constants/requestError"
	"core/crypto"
	"core/db/connect"
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"helpers/errors"
	"models/User"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdateUserSpec(t *testing.T) {
	Convey("Update user tests", t, func() {
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

		userToUpdate, _ := json.Marshal(map[string]interface{}{
			"firstName": "updated",
			"lastName":  "updated",
			"username":  "updated",
		})

		Convey("Update user should be successful and return updated user", func() {
			client := &http.Client{}
			req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(userToUpdate))
			req.Header.Set("Authorization", createdUser.Token)
			req.Header.Set("X-Real-IP", "127.0.0.1")
			res, _ := client.Do(req)

			var responseBody map[string]interface{}
			_ = json.NewDecoder(res.Body).Decode(&responseBody)

			expectedBody := map[string]interface{}{
				"id":         float64(createdUser.ID),
				"first_name": "updated",
				"last_name":  "updated",
				"username":   "updated",
			}

			So(res.StatusCode, ShouldEqual, http.StatusCreated)
			So(responseBody, ShouldResemble, expectedBody)
		})

		Convey("Update user without providing auth header should be failed and return UNAUTHORIZED error", func() {
			client := &http.Client{}
			req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(userToUpdate))
			req.Header.Set("X-Real-IP", "127.0.0.1")
			res, _ := client.Do(req)

			var responseBody map[string]interface{}
			_ = json.NewDecoder(res.Body).Decode(&responseBody)

			So(res.StatusCode, ShouldEqual, http.StatusUnauthorized)
			So(responseBody, ShouldResemble, utils.StructToMap(requestError.UNAUTHORIZED))
		})

		Convey("Update user with providing auth header but with another ip should be failed and return WRONG_IP error", func() {
			client := &http.Client{}
			req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(userToUpdate))
			req.Header.Set("Authorization", createdUser.Token)
			req.Header.Set("X-Real-IP", "127.0.0.2")
			res, _ := client.Do(req)

			var responseBody map[string]interface{}
			_ = json.NewDecoder(res.Body).Decode(&responseBody)

			So(res.StatusCode, ShouldEqual, http.StatusUnauthorized)
			So(responseBody, ShouldResemble, utils.StructToMap(requestError.WRONG_IP))
		})

		Convey("Update user with additional unauthorized filed \"test\" should be failed and return error \"test is not allowed\"", func() {
			userToUpdate, _ := json.Marshal(map[string]interface{}{
				"firstName": "updated",
				"lastName":  "updated",
				"username":  "updated",
				"test":      "test",
			})

			client := &http.Client{}
			req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(userToUpdate))
			req.Header.Set("Authorization", createdUser.Token)
			req.Header.Set("X-Real-IP", "127.0.0.1")
			res, _ := client.Do(req)

			var responseBody map[string]interface{}
			_ = json.NewDecoder(res.Body).Decode(&responseBody)

			expectedError := utils.StructToMap(errors.IRequestError{http.StatusBadRequest, "test is not allowed", "NOT_VALID"})

			So(res.StatusCode, ShouldEqual, http.StatusBadRequest)
			So(responseBody, ShouldResemble, expectedError)
		})

		Convey("Update only firstName should be successful and return user with updated firstName", func() {
			userToUpdate, _ := json.Marshal(map[string]interface{}{
				"firstName": "updated",
			})

			client := &http.Client{}
			req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(userToUpdate))
			req.Header.Set("Authorization", createdUser.Token)
			req.Header.Set("X-Real-IP", "127.0.0.1")
			res, _ := client.Do(req)

			var responseBody map[string]interface{}
			_ = json.NewDecoder(res.Body).Decode(&responseBody)

			expectedBody := map[string]interface{}{
				"id":         float64(createdUser.ID),
				"first_name": "updated",
				"last_name":  "string",
				"username":   "string",
			}

			So(res.StatusCode, ShouldEqual, http.StatusCreated)
			So(responseBody, ShouldResemble, expectedBody)
		})

		Convey("Update only lastName should be successful and return user with updated lastName", func() {
			userToUpdate, _ := json.Marshal(map[string]interface{}{
				"lastName": "updated",
			})

			client := &http.Client{}
			req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(userToUpdate))
			req.Header.Set("Authorization", createdUser.Token)
			req.Header.Set("X-Real-IP", "127.0.0.1")
			res, _ := client.Do(req)

			var responseBody map[string]interface{}
			_ = json.NewDecoder(res.Body).Decode(&responseBody)

			expectedBody := map[string]interface{}{
				"id":         float64(createdUser.ID),
				"first_name": "string",
				"last_name":  "updated",
				"username":   "string",
			}

			So(res.StatusCode, ShouldEqual, http.StatusCreated)
			So(responseBody, ShouldResemble, expectedBody)
		})

		Convey("Update only username should be successful and return user with updated username", func() {
			userToUpdate, _ := json.Marshal(map[string]interface{}{
				"username": "updated",
			})

			client := &http.Client{}
			req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(userToUpdate))
			req.Header.Set("Authorization", createdUser.Token)
			req.Header.Set("X-Real-IP", "127.0.0.1")
			res, _ := client.Do(req)

			var responseBody map[string]interface{}
			_ = json.NewDecoder(res.Body).Decode(&responseBody)

			expectedBody := map[string]interface{}{
				"id":         float64(createdUser.ID),
				"first_name": "string",
				"last_name":  "string",
				"username":   "updated",
			}

			So(res.StatusCode, ShouldEqual, http.StatusCreated)
			So(responseBody, ShouldResemble, expectedBody)
		})

		Convey("Update with empty body should be failed and return user with error \"One of fields should be provided\"", func() {
			userToUpdate, _ := json.Marshal(map[string]interface{}{})

			client := &http.Client{}
			req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(userToUpdate))
			req.Header.Set("Authorization", createdUser.Token)
			req.Header.Set("X-Real-IP", "127.0.0.1")
			res, _ := client.Do(req)

			var responseBody map[string]interface{}
			_ = json.NewDecoder(res.Body).Decode(&responseBody)

			expectedError := utils.StructToMap(errors.IRequestError{http.StatusBadRequest, "One of fields should be provided", "NOT_VALID"})

			So(res.StatusCode, ShouldEqual, http.StatusBadRequest)
			So(responseBody, ShouldResemble, expectedError)
		})

		connect.DB.Exec("delete from users")
	})
}
