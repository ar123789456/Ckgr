package usecase

import (
	"cgr/link"
	"cgr/models"
	"cgr/tool/logger"
	"context"
)

type Usecase struct {
	repository link.Repository
	logger     *logger.Logger
}

func NewUsecase(repository link.Repository, logger *logger.Logger) *Usecase {
	return &Usecase{
		repository: repository,
		logger:     logger,
	}
}

func (uc *Usecase) Create(con context.Context, link models.Link) error {
	return uc.repository.Create(con, link)
}
func (uc *Usecase) Delete(con context.Context, id int) error {
	return uc.repository.Delete(con, id)
}
func (uc *Usecase) GetAll(con context.Context) ([]models.Link, error) {
	return uc.repository.GetAll(con)
}
