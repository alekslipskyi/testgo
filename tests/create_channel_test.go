package tests

import (
	"./utils"
	"constants/requestError"
	"core/db/connect"
	"core/db/types"
	. "github.com/smartystreets/goconvey/convey"
	"helpers/errors"
	"models/Channel"
	"net/http"
	"testing"
)

func TestCreateChannelSpec(t *testing.T) {
	Convey("Create channel tests", t, func() {
		connect.DB.Exec("delete from users")
		connect.DB.Exec("delete from channels")

		createdUser := utils.CreateUser()

		requester := utils.Requester{}
		requester.Init("/api/v0/channel", map[string]interface{}{
			"auth": createdUser.Token,
		})

		Convey("Create channel request should create channel and return it", func() {

			res, responseBody := requester.POST(map[string]interface{}{
				"name": "test",
			})

			createdChannel := Channel.FindOne(types.QueryOptions{Where: types.Where{"name": "test"}, Attributes: types.Attributes{"_id"}})

			So(createdChannel.IsExist(), ShouldBeTrue)

			expectedBody := map[string]interface{}{
				"name":  "test",
				"id":    float64(createdChannel.ID),
				"users": []interface{}{float64(createdUser.ID)},
			}

			So(res.StatusCode, ShouldEqual, http.StatusOK)
			So(responseBody, ShouldResemble, expectedBody)
		})

		Convey("Create channel without name should return bad request with error \"name is required\"", func() {

			res, responseBody := requester.POST(map[string]interface{}{})

			createdChannel := Channel.FindOne(types.QueryOptions{Where: types.Where{"name": "test"}, Attributes: types.Attributes{"_id"}})

			So(!createdChannel.IsExist(), ShouldBeTrue)

			expectedError := utils.StructToMap(errors.IRequestError{http.StatusBadRequest, "name is required", "NOT_VALID"})

			So(res.StatusCode, ShouldEqual, http.StatusBadRequest)
			So(responseBody, ShouldResemble, expectedError)
		})

		Convey("Create channel with additional field should return bad request with error \"test1 is not allowed\"", func() {

			res, responseBody := requester.POST(map[string]interface{}{
				"name":  "test",
				"test1": "test2",
			})

			createdChannel := Channel.FindOne(types.QueryOptions{Where: types.Where{"name": "test"}, Attributes: types.Attributes{"_id"}})

			So(!createdChannel.IsExist(), ShouldBeTrue)

			expectedError := utils.StructToMap(errors.IRequestError{http.StatusBadRequest, "test1 is not allowed", "NOT_VALID"})

			So(res.StatusCode, ShouldEqual, http.StatusBadRequest)
			So(responseBody, ShouldResemble, expectedError)
		})

		Convey("Create channel without auth header should return UNAUTHORIZED error", func() {

			requester.UnsetAuth()
			res, responseBody := requester.POST(map[string]interface{}{
				"name": "test",
			})

			createdChannel := Channel.FindOne(types.QueryOptions{Where: types.Where{"name": "test"}, Attributes: types.Attributes{"_id"}})

			So(!createdChannel.IsExist(), ShouldBeTrue)

			So(res.StatusCode, ShouldEqual, http.StatusUnauthorized)
			So(responseBody, ShouldResemble, utils.StructToMap(requestError.UNAUTHORIZED))
		})
	})
}
