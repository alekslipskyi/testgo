package tests

import (
	"./utils"
	"constants/requestError"
	"core/db/connect"
	"core/db/types"
	"core/logger"
	. "github.com/smartystreets/goconvey/convey"
	"models/User"
	"net/http"
	"testing"
)

func TestDeleteUserSpec(t *testing.T) {
	var log = logger.Logger{Context: "delete user tests", Colors: logger.Colors{Info: logger.GREEN}}

	Convey("Delete user Test", t, func() {
		connect.DB.Exec("delete from users")

		requester := utils.Requester{}
		requester.Init("/api/v0/user", map[string]interface{}{})

		Convey("Delete user should be successful and return ok", func() {
			log.Info("Delete user should be successful and return ok")

			createdUser := utils.CreateUser()
			requester.SetAuth(createdUser.Token)

			res, _ := requester.DELETE()

			user := User.Find(types.QueryOptions{
				Where: types.Where{"username": "string"},
			})

			So(res.StatusCode, ShouldEqual, http.StatusOK)
			So(user.IsNotExist(), ShouldBeTrue)
		})

		Convey("Delete user without providing auth header should be failed and return UNAUTHORIZED error", func() {
			log.Info("Delete user without providing auth header should be failed and return UNAUTHORIZED error")
			utils.CreateUser()

			res, responseBody := requester.DELETE()

			user := User.Find(types.QueryOptions{
				Where: types.Where{"username": "string"},
			})

			So(res.StatusCode, ShouldEqual, http.StatusUnauthorized)
			So(!user.IsNotExist(), ShouldBeTrue)
			So(responseBody, ShouldResemble, utils.StructToMap(requestError.UNAUTHORIZED))
		})

		Reset(func() {
			requester.UnsetAuth()
		})
	})
}
