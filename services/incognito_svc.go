package services

import "xanny-tree-api/dto"

func (s *compServices) UploadIncognitoMessage(message string) error {
	err := s.repo.UploadIncognitoMessage(message)
	if err != nil {
		return err
	}

	email := dto.Email{
		Email:   "ramadiaz221@gmail.com",
		Subject: "New Incognito Message",
		Body:    message,
	}

	go func(email dto.Email) {
		_ = s.SendEmail(email)
	}(email)

	return nil
}
