package services

import (
	"xanny-tree-api/dto"
	"xanny-tree-api/repositories"
)

type CompService interface{
	RegisterUrl(data dto.Tree) error
}

type compServices struct {
	repo repositories.CompRepository
}

func NewService(r repositories.CompRepository) *compServices {
	return &compServices{
		repo: r,
	}
}
