package services

func(s *compServices) UploadIncognitoMessage(message string) error {
	return s.repo.UploadIncognitoMessage(message)
}