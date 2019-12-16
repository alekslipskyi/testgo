package tests

import (
	"./utils"
	"constants/requestError"
	"core/db/connect"
	"core/db/types"
	. "github.com/smartystreets/goconvey/convey"
	"helpers/errors"
	"models/User"
	"net/http"
	"testing"
)

func TestSignUpSpec(t *testing.T) {
	Convey("Sign up tests", t, func() {
		connect.DB.Exec("delete from users")

		requester := utils.Requester{}
		requester.Init("/api/v0/user", map[string]interface{}{})

		user := map[string]interface{}{
			"firstName": "string",
			"lastName":  "string",
			"username":  "string",
			"password":  "string",
		}

		Convey("sign up should be successful and return created user", func() {

			res, responseBody := requester.POST(user)

			createdUser := User.Find(types.QueryOptions{
				Where: types.Where{"username": "string"},
			})
			createdUser.GenerateToken()

			expectedBody := map[string]interface{}{
				"id":         float64(createdUser.ID),
				"first_name": "string",
				"last_name":  "string",
				"username":   "string",
				"token":      createdUser.Token,
			}

			So(res.StatusCode, ShouldEqual, http.StatusCreated)
			So(responseBody, ShouldResemble, expectedBody)
		})

		Convey("sign up with already existed credentials should be failed and return the USER_ALREADY_EXIST error", func() {

			requester.POST(user)
			res, responseBody := requester.POST(user)

			createdUser := User.Find(types.QueryOptions{
				Where: types.Where{"username": "string"},
			})
			createdUser.GenerateToken()

			So(res.StatusCode, ShouldEqual, http.StatusBadRequest)
			So(responseBody, ShouldResemble, utils.StructToMap(requestError.USER_ALREADY_EXIST))
		})

		Convey("sign up with unauthorized field should be failed and return the error with message \"test is not allowed\"", func() {

			wrongUser := map[string]interface{}{
				"firstName": "string",
				"lastName":  "string",
				"username":  "string",
				"password":  "string",
				"test":      "string",
			}
			res, responseBody := requester.POST(wrongUser)

			createdUser := User.Find(types.QueryOptions{
				Where: types.Where{"username": "string"},
			})
			createdUser.GenerateToken()

			expectedError := utils.StructToMap(errors.IRequestError{http.StatusBadRequest, "test is not allowed", "NOT_VALID"})

			So(res.StatusCode, ShouldEqual, http.StatusBadRequest)
			So(responseBody, ShouldResemble, expectedError)
		})

		Convey("sign up without firstName should be failed and return the error with message \"firstName is required\"", func() {

			wrongUser := map[string]interface{}{
				"lastName": "string",
				"username": "string",
				"password": "string",
			}
			res, responseBody := requester.POST(wrongUser)

			createdUser := User.Find(types.QueryOptions{
				Where: types.Where{"username": "string"},
			})
			createdUser.GenerateToken()

			expectedError := utils.StructToMap(errors.IRequestError{http.StatusBadRequest, "firstName is required", "NOT_VALID"})

			So(res.StatusCode, ShouldEqual, http.StatusBadRequest)
			So(responseBody, ShouldResemble, expectedError)
		})

		Convey("sign up without lastName should be failed and return the error with message \"lastName is required\"", func() {

			wrongUser := map[string]interface{}{
				"firstName": "string",
				"username":  "string",
				"password":  "string",
			}
			res, responseBody := requester.POST(wrongUser)

			createdUser := User.Find(types.QueryOptions{
				Where: types.Where{"username": "string"},
			})
			createdUser.GenerateToken()

			expectedError := utils.StructToMap(errors.IRequestError{http.StatusBadRequest, "lastName is required", "NOT_VALID"})

			So(res.StatusCode, ShouldEqual, http.StatusBadRequest)
			So(responseBody, ShouldResemble, expectedError)
		})

		Convey("sign up without username should be failed and return the error with message \"username is required\"", func() {

			wrongUser := map[string]interface{}{
				"firstName": "string",
				"lastName":  "string",
				"password":  "string",
			}
			res, responseBody := requester.POST(wrongUser)

			createdUser := User.Find(types.QueryOptions{
				Where: types.Where{"username": "string"},
			})
			createdUser.GenerateToken()

			expectedError := utils.StructToMap(errors.IRequestError{http.StatusBadRequest, "username is required", "NOT_VALID"})

			So(res.StatusCode, ShouldEqual, http.StatusBadRequest)
			So(responseBody, ShouldResemble, expectedError)
		})

		Convey("sign up without password should be failed and return the error with message \"password is required\"", func() {

			wrongUser := map[string]interface{}{
				"firstName": "string",
				"lastName":  "string",
				"username":  "string",
			}
			res, responseBody := requester.POST(wrongUser)

			createdUser := User.Find(types.QueryOptions{
				Where: types.Where{"username": "string"},
			})
			createdUser.GenerateToken()

			expectedError := utils.StructToMap(errors.IRequestError{http.StatusBadRequest, "password is required", "NOT_VALID"})

			So(res.StatusCode, ShouldEqual, http.StatusBadRequest)
			So(responseBody, ShouldResemble, expectedError)
		})

		Convey("sign up with not valid firstName should be failed and return the error with message \"firstName must be a string\"", func() {

			wrongUser := map[string]interface{}{
				"firstName": 1,
				"lastName":  "string",
				"username":  "string",
				"password":  "string",
			}
			res, responseBody := requester.POST(wrongUser)

			createdUser := User.Find(types.QueryOptions{
				Where: types.Where{"username": "string"},
			})
			createdUser.GenerateToken()

			expectedError := utils.StructToMap(errors.IRequestError{http.StatusBadRequest, "firstName must be a string", "NOT_VALID"})

			So(res.StatusCode, ShouldEqual, http.StatusBadRequest)
			So(responseBody, ShouldResemble, expectedError)
		})

		Convey("sign up with not valid lastName should be failed and return the error with message \"lastName must be a string\"", func() {

			wrongUser := map[string]interface{}{
				"firstName": "string",
				"lastName":  1,
				"username":  "string",
				"password":  "string",
			}
			res, responseBody := requester.POST(wrongUser)

			createdUser := User.Find(types.QueryOptions{
				Where: types.Where{"username": "string"},
			})
			createdUser.GenerateToken()

			expectedError := utils.StructToMap(errors.IRequestError{http.StatusBadRequest, "lastName must be a string", "NOT_VALID"})

			So(res.StatusCode, ShouldEqual, http.StatusBadRequest)
			So(responseBody, ShouldResemble, expectedError)
		})

		Convey("sign up with not valid username should be failed and return the error with message \"username must be a string\"", func() {

			wrongUser := map[string]interface{}{
				"firstName": "string",
				"lastName":  "string",
				"username":  1,
				"password":  "string",
			}
			res, responseBody := requester.POST(wrongUser)

			createdUser := User.Find(types.QueryOptions{
				Where: types.Where{"username": "string"},
			})
			createdUser.GenerateToken()

			expectedError := utils.StructToMap(errors.IRequestError{http.StatusBadRequest, "username must be a string", "NOT_VALID"})

			So(res.StatusCode, ShouldEqual, http.StatusBadRequest)
			So(responseBody, ShouldResemble, expectedError)
		})

		Convey("sign up with not valid password should be failed and return the error with message \"password must be a string\"", func() {

			wrongUser := map[string]interface{}{
				"firstName": "string",
				"lastName":  "string",
				"username":  "string",
				"password":  1,
			}
			res, responseBody := requester.POST(wrongUser)

			createdUser := User.Find(types.QueryOptions{
				Where: types.Where{"username": "string"},
			})
			createdUser.GenerateToken()

			expectedError := utils.StructToMap(errors.IRequestError{http.StatusBadRequest, "password must be a string", "NOT_VALID"})

			So(res.StatusCode, ShouldEqual, http.StatusBadRequest)
			So(responseBody, ShouldResemble, expectedError)
		})

		Convey("sign up with firstName empty string should be failed and return the error with message \"firstName must be longer than 1\"", func() {

			wrongUser := map[string]interface{}{
				"firstName": "",
				"lastName":  "string",
				"username":  "string",
				"password":  "string",
			}
			res, responseBody := requester.POST(wrongUser)

			createdUser := User.Find(types.QueryOptions{
				Where: types.Where{"username": "string"},
			})
			createdUser.GenerateToken()

			expectedError := utils.StructToMap(errors.IRequestError{http.StatusBadRequest, "firstName must be longer than 1", "NOT_VALID"})

			So(res.StatusCode, ShouldEqual, http.StatusBadRequest)
			So(responseBody, ShouldResemble, expectedError)
		})

		Convey("sign up with lastName empty string should be failed and return the error with message \"lastName must be longer than 1\"", func() {

			wrongUser := map[string]interface{}{
				"firstName": "string",
				"lastName":  "",
				"username":  "string",
				"password":  "string",
			}
			res, responseBody := requester.POST(wrongUser)

			createdUser := User.Find(types.QueryOptions{
				Where: types.Where{"username": "string"},
			})
			createdUser.GenerateToken()

			expectedError := utils.StructToMap(errors.IRequestError{http.StatusBadRequest, "lastName must be longer than 1", "NOT_VALID"})

			So(res.StatusCode, ShouldEqual, http.StatusBadRequest)
			So(responseBody, ShouldResemble, expectedError)
		})

		Convey("sign up with username empty string should be failed and return the error with message \"username must be longer than 1\"", func() {

			wrongUser := map[string]interface{}{
				"firstName": "string",
				"lastName":  "string",
				"username":  "",
				"password":  "string",
			}
			res, responseBody := requester.POST(wrongUser)

			createdUser := User.Find(types.QueryOptions{
				Where: types.Where{"username": "string"},
			})
			createdUser.GenerateToken()

			expectedError := utils.StructToMap(errors.IRequestError{http.StatusBadRequest, "username must be longer than 1", "NOT_VALID"})

			So(res.StatusCode, ShouldEqual, http.StatusBadRequest)
			So(responseBody, ShouldResemble, expectedError)
		})

		Convey("sign up with password empty string should be failed and return the error with message \"password must be longer than 1\"", func() {

			wrongUser := map[string]interface{}{
				"firstName": "string",
				"lastName":  "string",
				"username":  "string",
				"password":  "",
			}
			res, responseBody := requester.POST(wrongUser)

			createdUser := User.Find(types.QueryOptions{
				Where: types.Where{"username": "string"},
			})
			createdUser.GenerateToken()

			expectedError := utils.StructToMap(errors.IRequestError{http.StatusBadRequest, "password must be longer than 6", "NOT_VALID"})

			So(res.StatusCode, ShouldEqual, http.StatusBadRequest)
			So(responseBody, ShouldResemble, expectedError)
		})

		Convey("sign up with password length more than 10 should be failed and return the error with message \"password must be less than 1\"", func() {

			wrongUser := map[string]interface{}{
				"firstName": "string",
				"lastName":  "string",
				"username":  "string",
				"password":  "asdkaskdakdalkasdkskadskl",
			}
			res, responseBody := requester.POST(wrongUser)

			createdUser := User.Find(types.QueryOptions{
				Where: types.Where{"username": "string"},
			})
			createdUser.GenerateToken()

			expectedError := utils.StructToMap(errors.IRequestError{http.StatusBadRequest, "password must be less than 10", "NOT_VALID"})

			So(res.StatusCode, ShouldEqual, http.StatusBadRequest)
			So(responseBody, ShouldResemble, expectedError)
		})

		Reset(func() {
			connect.DB.Exec("delete from users")
		})
	})
}
