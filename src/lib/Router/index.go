package Router

import (
	"encoding/json"
	"helpers/errors"
	"net/http"
	"strings"
)

type failedHandler func(error errors.RequestError)

type New struct {
	thrown failedHandler
}

type Context struct {
	Res    http.ResponseWriter
	Req    *http.Request
	Body   map[string]interface{}
	Params map[string]interface{}
	Thrown failedHandler
}

func (router *New) EntryPoint(w http.ResponseWriter, r *http.Request) {
	ctx := Context{Res: w, Req: r}
	router.thrown = router.failed(ctx)

	exist, handler, params := router.getRoute(r.Method, r.RequestURI)

	if exist {
		ok, result := router.parseBody(w, r)
		ctx := Context{w, r, result, params, router.thrown}

		if !ok {
			return
		}

		if len(handler.middleware) != 0 {
			for _, middleware := range handler.middleware {
				ok, errMessage := middleware(ctx)

				if !ok {
					requestError := errors.RequestError{http.StatusBadRequest, errMessage, "NOT_VALID"}

					router.thrown(requestError)
					return
				}
			}
		}

		handler.success(ctx)
	} else {
		http.NotFound(w, r)
	}
}

func (router *New) getRoute(method string, url string) (bool, Route, map[string]interface{}) {
	var params = make(map[string]interface{})
	var findedRoute Route

	for route, routeParams := range routes {
		if route == method+url {
			return true, routeParams, params
		}

		routesParts := strings.Split(strings.Replace(url, routeParams.prefix, "", -1), "/")

		for i, param := range routeParams.params {
			if len(routesParts[i+1]) != 0 {
				params[param] = routesParts[i+1]
				findedRoute = routeParams
			} else {
				return false, Route{}, params
			}
		}
	}

	if len(params) != 0 {
		return true, findedRoute, params
	}

	return false, Route{}, params
}

func (router *New) parseBody(w http.ResponseWriter, r *http.Request) (bool, map[string]interface{}) {
	var body map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		println(err.Error())
		requestError := errors.RequestError{http.StatusBadRequest, "json is not valid", "JSON_NOT_VALID"}

		router.thrown(requestError)
		return false, body
	} else {
		return true, body
	}
}

func (router *New) failed(ctx Context) failedHandler {
	return func(error errors.RequestError) {
		ctx.Res.Header().Set("Content-Type", "application/json")
		ctx.Res.WriteHeader(error.StatusCode)

		js, _ := json.Marshal(error)

		_, _ = ctx.Res.Write(js)
	}
}
