package user

import (
	"constants/requestError"
	"core/Router"
	"core/crypto"
	"core/db"
	"core/db/types"
	"core/logger"
	"models/User"
	"net/http"
)

var log = logger.Logger{"user controller"}

type Controller struct {
	db struct {
		user db.Instance
	}
}

func (entity *Controller) Init() {
	entity.db.user = db.Instance{Name: "users"}
}

func (entity *Controller) handleSignUp(ctx Router.Context) {
	user := User.Find(types.QueryOptions{
		Where: types.Where{"username": ctx.Body["username"]},
	})

	if user.IsNotExist() {
		ctx.Reject(requestError.USER_ALREADY_EXIST)
		return
	}

	ctx.Body["password"] = crypto.GenerateHash(ctx.Body["password"].(string))

	userCreated := User.CreateAndFind(ctx.Body)
	userCreated.AddAllowIP(ctx.Req.Header.Get("X-Real-IP"))
	userCreated.GenerateToken()

	ctx.SendJson(userCreated, http.StatusCreated)
}

func (entity *Controller) delete(ctx Router.Context) {
	user := User.FindById(ctx.User.ID, []string{"_id"})

	if user.IsNotExist() {
		ctx.Reject(requestError.NOT_FOUND)
		return
	}

	user.Drop()
	ctx.Send("ok", http.StatusOK)
}

func (entity *Controller) handleUpdate(ctx Router.Context) {
	user := User.FindById(ctx.User.ID, []string{"_id", "username"})
	if user.ID == 0 {
		ctx.Reject(requestError.UNAUTHORIZED)
	}
	ctx.SendJson(user.UpdateAndFind(ctx.Body), http.StatusCreated)
}
