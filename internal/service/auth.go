package service

import (
	"johny-tuna/internal/utils"
	"regexp"
)

const (
	LoginByUsername = iota
	LoginByEmail
)

func (s *service) Login(loginOrEmail, password string) (string, error) {
	isEmail := checkIsEmail(loginOrEmail)

	var loginType int
	if isEmail {
		loginType = LoginByEmail
	} else {
		loginType = LoginByUsername
	}

	user, err := s.repo.Login(loginOrEmail, loginType)
	if err != nil {
		return "", err
	}

	if err := CheckPass(password, user); err != nil {
		return "", err
	}

	return GenerateToken(user)
}

func (s *service) Register(login, email, password string) error {
	hash, err := EncryptPass(password)
	if err != nil {
		return err
	}

	info := utils.UserInfo{
		Email:    email,
		Login:    login,
		Password: hash,
	}

	token := utils.StoreToken(info)

	if err := utils.SendMessage(email, token); err != nil {
		return err
	}

	return nil
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func checkIsEmail(email string) bool {
	return emailRegex.MatchString(email)
}

func (s *service) VerificationRegister(token string) (string, error) {
	var info utils.UserInfo

	if err := utils.VerifyToken(token, &info); err != nil {
		return "", err
	}

	user, err := s.repo.Register(info.Login, info.Email, info.Password)
	if err != nil {
		return "", err
	}

	return GenerateToken(user)
}
