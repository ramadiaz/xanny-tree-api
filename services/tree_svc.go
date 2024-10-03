package services

import "xanny-tree-api/dto"

func (s *compServices) RegisterUrl(data dto.Tree) error {
	return s.repo.RegisterUrl(data)
}

func (s *compServices) GetUrl(url string) (*string, error)  {
	return s.repo.GetUrl(url)
}