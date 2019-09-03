package auth

import (
	"constants/requestError"
	"core/Router"
	"core/db/types"
	"core/logger"
	"helpers"
	"models/User"
	"net/http"
)

var log = logger.Logger{"controller AUTH"}

type Controller struct{}

func (controller *Controller) handleAuth(ctx Router.Context) {
	user := User.Find(types.QueryOptions{
		Where: types.Where{"username": ctx.Req.URL.Query().Get("username")},
	})

	if user.IsNotExist() || !user.IsValidPassword(ctx.Req.URL.Query().Get("password")) {
		ctx.Reject(requestError.INVALID_CREDENTIAL)
		return
	}

	log.Log("user", user)

	helpers.OmitPrivateFields(&user)
	if user.IsIPNotExist(ctx.RequestIP) {
		if status := user.AddAllowIP(ctx.RequestIP); !status {
			ctx.Reject(requestError.UNEXPECTED_ERROR)
		}
	}

	user.GenerateToken()

	ctx.SendJson(user, http.StatusOK)
}
