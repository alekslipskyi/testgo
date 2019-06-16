package Router

import (
	"net/http"
	"os"
	"regexp"
)

type middleware func(ctx *Context) (bool, interface{}, interface{})
type Handler func(ctx Context)

type Route struct {
	success    Handler
	prefix     string
	middleware []middleware
	params     map[int]string
}

var routes = make(map[string]Route)

type Instance struct {
	Prefix    string
	ApiPrefix string
}

func (instance *Instance) GET(url string, handler Handler, middleware ...middleware) {
	instance.buildHandler(http.MethodGet, url, handler, middleware...)
}

func (instance *Instance) POST(url string, handler Handler, middleware ...middleware) {
	instance.buildHandler(http.MethodPost, url, handler, middleware...)
}

func (instance *Instance) PUT(url string, handler Handler, middleware ...middleware) {
	instance.buildHandler(http.MethodPut, url, handler, middleware...)
}

func (instance *Instance) DELETE(url string, handler Handler, middleware ...middleware) {
	instance.buildHandler(http.MethodDelete, url, handler, middleware...)
}

func (instance *Instance) PATCH(url string, handler Handler, middleware ...middleware) {
	instance.buildHandler(http.MethodPatch, url, handler, middleware...)
}

func (instance *Instance) buildHandler(method string, url string, handler Handler, middleware ...middleware) {
	prefix := instance.getPrefix()
	params := instance.findParams(url)

	route := Route{handler, prefix, middleware, params}
	routes[method+instance.getRoute(url, prefix)] = route
}

func (instance *Instance) findParams(url string) map[int]string {
	params := make(map[int]string)

	reParams := regexp.MustCompile(`{\w+}`)
	withoutBrackets := regexp.MustCompile(`[{}]`)

	matchedParams := reParams.FindStringSubmatch(url)

	if len(matchedParams) != 0 {
		for i, param := range matchedParams {
			params[i] = withoutBrackets.ReplaceAllString(param, "")
		}
	}

	return params
}

func (instance *Instance) getPrefix() string {
	apiPrefix := instance.ApiPrefix
	if len(apiPrefix) == 0 {
		apiPrefix = "/api/" + os.Getenv("API_VERSION")
	}

	return apiPrefix + instance.Prefix
}

func (instance *Instance) getRoute(url string, prefix string) string {
	if url == "/" {
		return prefix
	}

	return prefix + url
}
