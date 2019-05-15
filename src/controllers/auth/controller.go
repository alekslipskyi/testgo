package auth

import (
	"core/db"
	"fmt"
	"lib/Router"
)

type Controller struct {
	db.Instance
}

func (controller *Controller) handleAuth(ctx Router.Context) {
	fmt.Fprintf(ctx.Res, "Auth")
}
