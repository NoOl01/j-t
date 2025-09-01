package service

import (
	"fmt"
	"johny-tuna/internal/handler/dto"
	"johny-tuna/internal/utils"
)

func (s *service) Appeal(body dto.AppealBody) error {
	theme := getAppealTheme(body.Theme)

	message := fmt.Sprintf("Здравствуйте %s.\r\n"+
		"Вы подали обращение: %s. \r\n"+
		"Ваше сообщение: %s. \r\n", body.Name, theme, body.Message)
	if err := utils.SendMessage(body.Email, message); err != nil {
		return err
	}

	if err := utils.SendMessage(body.Email, message); err != nil {
		return err
	}

	return nil
}

const (
	Feedback = iota
	Question
	Problem
	Booking
)

func getAppealTheme(theme int) string {
	switch theme {
	case Feedback:
		return "Отзыв или предложение"
	case Question:
		return "Вопрос"
	case Problem:
		return "Проблема"
	case Booking:
		return "Бронирование"
	default:
		return ""
	}
}
