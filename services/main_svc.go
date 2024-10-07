package services

import (
	"xanny-tree-api/dto"
	"xanny-tree-api/repositories"
)

type CompService interface{
	RegisterUrl(data dto.Tree) error
	GetUrl(url string) (*string, error) 

	AddLike() (*int64, error) 
	GetLike() (*int64, error) 

	UploadIncognitoMessage(message string) error
}

type compServices struct {
	repo repositories.CompRepository
}

func NewService(r repositories.CompRepository) *compServices {
	return &compServices{
		repo: r,
	}
}
