package list

import (
	"constants/requestError"
	"core/Router"
	"core/logger"
	"models/List"
	"net/http"
	"strconv"
)

type Controller struct{}

var log = logger.Logger{Context: "Controller list"}

func (controller *Controller) getListById(ctx Router.Context) List.List {
	id, _ := strconv.ParseInt(ctx.Params["id"].(string), 10, 64)
	list := List.FindById(id, ctx.User.ID)

	if list.IsNotExist() {
		ctx.Reject(requestError.NOT_FOUND)
		return List.List{}
	}

	return list
}

func (controller *Controller) index(ctx Router.Context) {
	ctx.SendJson(List.FindManyByUser(ctx.User.ID), http.StatusOK)
}

func (controller *Controller) create(ctx Router.Context) {
	body := ctx.Body
	body["user_id"] = ctx.User.ID

	ctx.SendJson(List.CreateAndFind(body), http.StatusCreated)
}

func (controller *Controller) get(ctx Router.Context) {
	list := controller.getListById(ctx)
	ctx.SendJson(list, http.StatusOK)
}

func (controller *Controller) delete(ctx Router.Context) {
	list := controller.getListById(ctx)

	status := list.Drop()

	if status {
		ctx.Send("ok", http.StatusOK)
	} else {
		ctx.SendJson(requestError.UNEXPECTED_ERROR, http.StatusInternalServerError)
	}
}

func (controller *Controller) update(ctx Router.Context) {
	list := controller.getListById(ctx)
	ctx.SendJson(list.UpdateAndFind(ctx.Body), http.StatusOK)
}
