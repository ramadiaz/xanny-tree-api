package services

func (s *compServices) AddLike() (*int64, error) {
	return s.repo.AddLike()
}

func (s *compServices) GetLike() (*int64, error) {
	return s.repo.GetLike()
}