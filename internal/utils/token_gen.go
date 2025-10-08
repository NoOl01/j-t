package utils

import (
	"crypto/rand"
	"github.com/google/uuid"
	"johny-tuna/internal/errs"
	"math/big"
	"time"
)

type UserInfo struct {
	Email    string
	Login    string
	Password string
}

var tokens = make(map[string]UserInfo)
var tokenTimers = make(map[string]*time.Timer)

var otpToken = make(map[string]int64)
var otpTokenTimer = make(map[string]*time.Timer)

func StoreToken(info UserInfo) string {
	token := uuid.New().String()
	timer := time.AfterFunc(30*time.Minute, func() {
		deleteToken(token)
	})

	tokens[token] = info
	tokenTimers[token] = timer

	return token
}

func StoreOtpToken(email string) int64 {
	maxValue := big.NewInt(10000)
	token, err := rand.Int(rand.Reader, maxValue)
	if err != nil {
		return 0
	}

	timer := time.AfterFunc(30*time.Minute, func() {
		deleteOtpToken(email)
	})

	otpToken[email] = token.Int64()
	otpTokenTimer[email] = timer

	return token.Int64()
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

func VerifyOtpToken(token int64, email string) error {
	if val, exists := otpToken[email]; exists {
		if val != token {
			return errs.WrongToken
		}
		if timer, ok := otpTokenTimer[email]; ok {
			timer.Stop()
			deleteOtpToken(email)
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

func deleteOtpToken(email string) {
	delete(otpToken, email)
	if timer, ok := otpTokenTimer[email]; ok {
		timer.Stop()
		delete(otpTokenTimer, email)
	}
}
