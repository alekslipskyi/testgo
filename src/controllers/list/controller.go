package list

import (
	"core/db"
	"fmt"
	"lib/Router"
)

type Controller struct {
	list db.Instance
}

func (controller *Controller) index(ctx Router.Context) {
	fmt.Println("list are", ctx.User.ID)
}
