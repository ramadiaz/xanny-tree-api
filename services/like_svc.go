package services

func (s *compServices) AddLike() (*int64, error) {
	return s.repo.AddLike()
}