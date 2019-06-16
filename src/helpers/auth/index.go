package auth

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"helpers/errors"
	"lib/Router"
	"lib/crypto"
	"models/User"
	"net/http"
)

func getUserFromToken(token string) (bool, User.Token) {
	var Json User.Token

	decodedToken, _ := hex.DecodeString(token)
	fmt.Println("test are:", decodedToken)

	decodedJson := crypto.DecodeToken(decodedToken)

	fmt.Println("test are:", decodedJson)

	if err := json.Unmarshal(decodedJson, &Json); err != nil {
		fmt.Println(err)

		return false, User.Token{}
	}

	return true, User.Token{}
}

func IsAuthenticated(ctx *Router.Context) (bool, interface{}, interface{}) {
	fmt.Println("token", ctx.Req.Header.Get("Authorization"))

	isValid, userData := getUserFromToken(ctx.Req.Header.Get("Authorization"))

	if !isValid {
		return false, nil, errors.RequestError{
			StatusCode: http.StatusUnauthorized,
			Message:    "Unauthorized",
			Token:      "UNAUTHORIZED",
		}
	}

	ctx.User = userData

	return true, nil, nil
}
