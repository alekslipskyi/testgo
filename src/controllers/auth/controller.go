package auth

import (
	"core/db/types"
	"encoding/json"
	"helpers"
	"helpers/errors"
	"lib/Router"
	"lib/crypto"
	"models/User"
	"net/http"
)

type Controller struct{}

func (controller *Controller) handleAuth(ctx Router.Context) {
	user := User.Find(types.QueryOptions{
		Where: types.Where{"username": ctx.Req.URL.Query().Get("username")},
	})

	if len(user.Password) == 0 || user.Password != crypto.GenerateHash(ctx.Req.URL.Query().Get("password")) {
		requestError := errors.RequestError{
			http.StatusNotFound,
			"invalid credentials",
			"INVALID_CREDENTIALS",
		}
		ctx.Reject(requestError)
		return
	}

	helpers.OmitPrivateFields(&user)
	user.GenerateToken()

	payload, _ := json.Marshal(&user)

	ctx.SendJson(payload, http.StatusOK)
}
