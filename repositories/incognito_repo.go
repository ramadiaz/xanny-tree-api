package repositories

func (r *compRepository) UploadIncognitoMessage(message string) error {
	_, err := r.DB.Exec("INSERT INTO incognito_messages (message) VALUES($1)", message)
	if err != nil {
		return err
	}

	return nil
}
