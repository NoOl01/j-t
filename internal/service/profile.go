package service

import "strings"

func (s *service) GetProfileIdFromToken(token string) (int64, error) {
	token = strings.ReplaceAll(token, "Bearer ", "")

	claims, err := DecodeToken(token)
	if err != nil {
		return 0, err
	}

	id, err := GetIdFromToken(claims)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *service) GetProfileInfo(id int64) (interface{}, error) {
	return s.repo.GetProfileInfo(id)
}

func (s *service) EditProfileEmail(id int64, newValue string) error {
	return s.repo.EditProfileEmail(id, newValue)
}

func (s *service) EditProfileLogin(id int64, newValue string) error {
	return s.repo.EditProfileLogin(id, newValue)
}
