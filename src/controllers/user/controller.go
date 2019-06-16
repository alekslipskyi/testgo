package user

import (
	"core/db"
	"core/db/types"
	"fmt"
	"helpers/errors"
	"lib/Router"
	"lib/crypto"
	"models/User"
	"net/http"
)

type Controller struct {
	db struct {
		user db.Instance
	}
}

func (c *Controller) Init() {
	c.db.user = db.Instance{Name: "users"}
}

func (c *Controller) handleSignUp(ctx Router.Context) {
	user := User.Find(types.QueryOptions{
		Where: types.Where{"username": ctx.Body["username"]},
	})

	if user.ID != 0 {
		ctx.Reject(errors.RequestError{
			http.StatusBadRequest,
			"User already exist",
			"USER_ALREADY_EXISTS"})
		return
	}

	ctx.Body["password"] = crypto.GenerateHash(ctx.Body["password"].(string))

	c.db.user.Create(ctx.Body)
	ctx.Send("User is created", http.StatusCreated)
}

func (c *Controller) handleTest(ctx Router.Context) {
	fmt.Fprintf(ctx.Res, "Test")
}
