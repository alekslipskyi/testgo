package user

import (
	"fmt"
	"lib/Router"
)

type Controller struct {
}

func (controller *Controller) handleSignUp(ctx Router.Context) {
	fmt.Println(ctx.Body["firstName"])
	fmt.Println(ctx.Params["id"])
}

func (controller *Controller) handleTest(ctx Router.Context) {
	fmt.Fprintf(ctx.Res, "Test")
}
