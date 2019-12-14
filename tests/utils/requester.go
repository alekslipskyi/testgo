package utils

import (
	"../../src"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
)

type Requester struct {
	prefix  string
	headers map[string]interface{}
	isDebug bool
}

func (req *Requester) Init(prefix string, headers map[string]interface{}) {
	srv := httptest.NewServer(router.Handler())
	req.prefix = srv.URL + prefix
	req.headers = headers
}

func (req *Requester) ChangeParams(providedParams map[string]interface{}) {
	newParams := map[string]interface{}{}

	for key, val := range req.headers {
		if providedParams[key] != nil {
			newParams[key] = providedParams[key]
		} else {
			newParams[key] = val
		}
	}

	req.headers = newParams
}

func (req *Requester) GET(params ...interface{}) (*http.Response, interface{}) {
	if len(params) != 0 {
		return req.exec(params[0], "GET", nil)
	} else {
		return req.exec("", "GET", nil)
	}
}

func (req *Requester) DELETE(params ...interface{}) (*http.Response, interface{}) {
	if len(params) != 0 {
		return req.exec(params[0], "DELETE", nil)
	} else {
		return req.exec("", "DELETE", nil)
	}
}

func (req *Requester) POST(params ...interface{}) (*http.Response, interface{}) {
	if len(params) != 0 && reflect.TypeOf(params[0]) == reflect.TypeOf("") {
		return req.exec(params[0], "POST", req.getPayload(params[1]))
	} else {
		return req.exec("", "POST", req.getPayload(params[0]))
	}
}

func (req *Requester) PUT(params ...interface{}) (*http.Response, interface{}) {
	if len(params) != 0 && reflect.TypeOf(params[0]) == reflect.TypeOf("") {
		if len(params) > 1 {
			return req.exec(params[0], "PUT", req.getPayload(params[1]))
		} else {
			return req.exec(params[0], "PUT", nil)
		}
	} else {
		return req.exec("", "PUT", req.getPayload(params[0]))
	}
}

func (req *Requester) Debug() {
	req.isDebug = true
	os.Setenv("LOGGER_LEVEL", "DEBUG")
}

func (req *Requester) SetAuth(token string) {
	req.headers["auth"] = token
}

func (req *Requester) SetHeader(key string, val string) {
	req.headers[key] = val
}

func (req *Requester) UnsetHeader(key string) {
	req.headers[key] = nil
}

func (req *Requester) UnsetAuth() {
	req.headers["auth"] = nil
}

func (req *Requester) getPayload(params ...interface{}) io.Reader {
	if len(params) > 0 {
		if req.isDebug {
			fmt.Printf("type of payload %s", reflect.TypeOf(params[0]))
		}

		if reflect.TypeOf(params[0]) == reflect.TypeOf([]interface{}{}) || reflect.TypeOf(params[0]) == reflect.TypeOf(map[string]interface{}{}) {
			json, _ := json.Marshal(params[0])
			return bytes.NewBuffer(json)
		}
	}

	return nil
}

func (req *Requester) exec(url interface{}, method string, payload io.Reader) (*http.Response, interface{}) {
	client := &http.Client{}
	var parsedURL string

	switch url.(type) {
	case string:
		parsedURL = url.(string)
	case float64:
		parsedURL = fmt.Sprintf("%d", int(url.(float64)))
	default:
		parsedURL = fmt.Sprintf("%d", url)
	}

	if req.isDebug {
		fmt.Printf("url is [%s] \n", method)
		fmt.Printf("method is [%s] \n", req.prefix+parsedURL)
		fmt.Printf("payload is [%s] \n", payload)
	}

	request, err := http.NewRequest(method, req.prefix+parsedURL, payload)

	if err != nil {
		fmt.Printf("error from creating request: %s \n", err)
	}

	if req.headers["auth"] != nil {
		request.Header.Set("Authorization", req.headers["auth"].(string))
		request.Header.Set("X-Real-IP", "127.0.0.1")
	}

	for key, val := range req.headers {
		if key != "auth" && val != nil {
			request.Header.Set(key, val.(string))
		}
	}

	res, err := client.Do(request)

	if err != nil {
		fmt.Printf("error from sending request: %s", err)
	}

	var responseBody interface{}
	_ = json.NewDecoder(res.Body).Decode(&responseBody)

	if req.isDebug {
		fmt.Printf("response body are ---> %s \n", responseBody)
		fmt.Printf("response status are ---> %d", res.StatusCode)

		os.Setenv("LOGGER_LEVEL", "NONE")
		req.isDebug = false
	}

	return res, responseBody
}
