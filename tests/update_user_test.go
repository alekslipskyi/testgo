package tests

import (
	"./utils"
	"constants/requestError"
	"core/crypto"
	"core/db/connect"
	. "github.com/smartystreets/goconvey/convey"
	"helpers/errors"
	"models/User"
	"net/http"
	"testing"
)

func TestUpdateUserSpec(t *testing.T) {
	connect.Init()
	defer connect.DB.Close()

	Convey("Update user tests", t, func() {
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
		requester.Init("/api/v0/user", map[string]interface{}{
			"auth": createdUser.Token,
		})

		userToUpdate := map[string]interface{}{
			"firstName": "updated",
			"lastName":  "updated",
			"username":  "updated",
		}

		Convey("Update user should be successful and return updated user", func() {

			res, responseBody := requester.PUT(userToUpdate)

			expectedBody := map[string]interface{}{
				"id":         float64(createdUser.ID),
				"first_name": "updated",
				"last_name":  "updated",
				"username":   "updated",
			}

			So(res.StatusCode, ShouldEqual, http.StatusCreated)
			So(responseBody, ShouldResemble, expectedBody)
		})

		Convey("Update user with additional unauthorized filed \"test\" should be failed and return error \"test is not allowed\"", func() {

			userToUpdate := map[string]interface{}{
				"firstName": "updated",
				"lastName":  "updated",
				"username":  "updated",
				"test":      "test",
			}

			res, responseBody := requester.PUT(userToUpdate)

			expectedError := utils.StructToMap(errors.IRequestError{http.StatusBadRequest, "test is not allowed", "NOT_VALID"})

			So(res.StatusCode, ShouldEqual, http.StatusBadRequest)
			So(responseBody, ShouldResemble, expectedError)
		})

		Convey("Update only firstName should be successful and return user with updated firstName", func() {

			userToUpdate := map[string]interface{}{
				"firstName": "updated",
			}

			res, responseBody := requester.PUT(userToUpdate)

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

			userToUpdate := map[string]interface{}{
				"lastName": "updated",
			}

			res, responseBody := requester.PUT(userToUpdate)

			expectedBody := map[string]interface{}{
				"id":         float64(createdUser.ID),
				"first_name": "string",
				"last_name":  "updated",
				"username":   "string",
			}

			So(res.StatusCode, ShouldEqual, http.StatusCreated)
			So(responseBody, ShouldResemble, expectedBody)
		})

		Convey("Update with empty body should be failed and return user with error \"One of fields should be provided\"", func() {

			userToUpdate := map[string]interface{}{}

			res, responseBody := requester.PUT(userToUpdate)

			expectedError := utils.StructToMap(errors.IRequestError{http.StatusBadRequest, "One of fields should be provided", "NOT_VALID"})

			So(res.StatusCode, ShouldEqual, http.StatusBadRequest)
			So(responseBody, ShouldResemble, expectedError)
		})

		Convey("Update user without providing auth header should be failed and return UNAUTHORIZED error", func() {
			requester.UnsetAuth()
			res, responseBody := requester.PUT(userToUpdate)

			So(res.StatusCode, ShouldEqual, http.StatusUnauthorized)
			So(responseBody, ShouldResemble, utils.StructToMap(requestError.UNAUTHORIZED))
		})
	})
}
