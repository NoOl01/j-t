package service

func (s *service) PlaceAnOrder(token string) error {
	claims, err := DecodeToken(token)
	if err != nil {
		return err
	}

	id, err := GetIdFromToken(claims)
	if err != nil {
		return err
	}

	return s.repo.PlaceAnOrder(id)
}
