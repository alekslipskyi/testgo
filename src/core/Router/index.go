package Router

import (
	"core/logger"
	"encoding/json"
	"fmt"
	"helpers/errors"
	"io"
	"models/User"
	"net/http"
	"net/url"
	"reflect"
	"runtime"
	"strings"
)

var log = logger.Logger{"ROUTER"}

type New struct {
	reject failedHandler
	closed bool
}

type Context struct {
	Res       http.ResponseWriter
	Req       *http.Request
	Body      map[string]interface{}
	Params    map[string]interface{}
	Reject    failedHandler
	SendJson  sendJsonHandler
	Send      sendHandler
	User      User.Token
	RequestIP string
}

func (router *New) EntryPoint(w http.ResponseWriter, r *http.Request) {
	requestURI, _ := url.Parse(r.RequestURI)
	requestURI.RawQuery = ""

	exist, handler, params := router.getRoute(r.Method, requestURI.String())

	if exist {
		conn := connection{w: w, r: r}
		ok, result := router.parseBody(w, r, conn)
		ctx := Context{
			w,
			r,
			result,
			params,
			conn.reject,
			conn.sendJson,
			conn.send,
			User.Token{},
			r.Header.Get("X-Real-IP"),
		}

		if !ok {
			return
		}

		if len(handler.middleware) != 0 {
			for _, middleware := range handler.middleware {
				ok, errMessage, providedError := middleware(&ctx)

				if !ok {
					var requestError errors.IRequestError

					if reflect.TypeOf(providedError) == reflect.TypeOf(errors.IRequestError{}) {
						requestError = providedError.(errors.IRequestError)
					} else {
						message := commonMessageError

						if reflect.TypeOf(message) == reflect.TypeOf(string(0)) {
							message = errMessage.(string)
						}

						requestError = errors.IRequestError{http.StatusBadRequest, message, tokenNotValid}
					}

					conn.reject(requestError)
					return
				}
			}
		}

		log.Debug("Receive request:", fmt.Sprintf("%s %s", r.Method, r.URL))

		if len(ctx.Body) != 0 {
			log.Debug("Body:", ctx.Body)
		}

		if len(ctx.Params) != 0 {
			log.Debug("Params:", ctx.Params)
		}
		log.Debug("handler is", runtime.FuncForPC(reflect.ValueOf(handler.success).Pointer()).Name())
		handler.success(ctx)
	} else {
		http.NotFound(w, r)
	}
}

func (router *New) getRoute(method string, url string) (bool, Route, map[string]interface{}) {
	var params = make(map[string]interface{})
	var findedRoute Route

	for route, routeParams := range routes {
		log.Log(method + url)
		log.Log(route)

		if route == method+url {
			return true, routeParams, params
		}

		parts := strings.Replace(url, routeParams.prefix, "", -1)

		if len(parts) != 0 {
			routesParts := strings.Split(parts, "/")

			for i, param := range routeParams.params {
				if len(routesParts[i+1]) != 0 {
					if method == routeParams.method {
						params[param] = routesParts[i+1]
						findedRoute = routeParams
					}
				} else {
					return false, Route{}, params
				}
			}
		}
	}

	if len(params) != 0 {
		return true, findedRoute, params
	}

	return false, Route{}, params
}

func (router *New) parseBody(w http.ResponseWriter, r *http.Request, conn connection) (bool, map[string]interface{}) {
	var body map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&body)

	if err == io.EOF {
		return true, body
	} else if err != nil {
		requestError := errors.IRequestError{http.StatusBadRequest, "json is not valid", "JSON_NOT_VALID"}

		conn.reject(requestError)
		return false, body
	} else {
		return true, body
	}
}
