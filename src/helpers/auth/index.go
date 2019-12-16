package auth

import (
	"constants/requestError"
	"core/Router"
	"core/crypto"
	"core/logger"
	"encoding/hex"
	"encoding/json"
	"models/User"
)

var log = logger.Logger{Context: "Auth Helper"}

func getUserFromToken(token string) (bool, User.Token) {
	var Json User.Token
	log.Debug("received token: ", token)
	if len(token) < 1 {
		return false, User.Token{}
	}

	decodedToken, _ := hex.DecodeString(token)
	decodedJson := crypto.DecodeToken(decodedToken)

	if err := json.Unmarshal(decodedJson, &Json); err != nil {
		log.Error("Error from decode token", err)

		return false, User.Token{}
	}

	return true, Json
}

func IsAuthenticated(ctx *Router.Context) (bool, interface{}, interface{}) {
	isValid, userData := getUserFromToken(ctx.Req.Header.Get("Authorization"))
	user := User.FindById(userData.ID, []string{"_id"})

	if !isValid || user.IsNotExist() {
		return false, nil, requestError.UNAUTHORIZED
	}

	userWithIP := User.FindWithIP(ctx.RequestIP, userData.ID)

	if userWithIP.IsNotExist() {
		return false, nil, requestError.WRONG_IP
	}

	ctx.User = userData

	return true, nil, nil
}
