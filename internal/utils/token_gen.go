package utils

import (
	"github.com/google/uuid"
	"johny-tuna/internal/errs"
	"time"
)

type UserInfo struct {
	Email    string
	Login    string
	Password string
}

var tokens = make(map[string]UserInfo)
var tokenTimers = make(map[string]*time.Timer)

func StoreToken(info UserInfo) string {
	token := uuid.New().String()
	timer := time.AfterFunc(30*time.Minute, func() {
		deleteToken(token)
	})

	tokens[token] = info
	tokenTimers[token] = timer

	return token
}

func VerifyToken(token string, info *UserInfo) error {
	if val, exists := tokens[token]; exists {
		*info = val
		if timer, ok := tokenTimers[token]; ok {
			timer.Stop()
			deleteToken(token)
		}

		return nil
	}
	return errs.WrongToken
}

func deleteToken(token string) {
	delete(tokens, token)
	if timer, ok := tokenTimers[token]; ok {
		timer.Stop()
		delete(tokenTimers, token)
	}
}
