package services

import "xanny-tree-api/dto"

func (s *compServices) RegisterUrl(data dto.Tree) error {
	return s.repo.RegisterUrl(data)
}