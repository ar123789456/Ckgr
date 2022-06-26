package usecase

import (
	"cgr/models"
	"cgr/project"
	"cgr/tool/logger"
	"context"
)

type Usecase struct {
	repository project.Repository
	logger     *logger.Logger
}

func NewUsecase(repository project.Repository, logger *logger.Logger) *Usecase {
	return &Usecase{
		repository: repository,
		logger:     logger,
	}
}

func (uc *Usecase) Post(c context.Context, n models.Project) error {
	uc.logger.InfoLogger.Println("usecase Post")
	return uc.repository.Post(c, n)
}
func (uc *Usecase) Get(c context.Context, id int) (models.Project, error) {
	uc.logger.InfoLogger.Println("usecase Get")
	return uc.repository.Get(c, id)
}
func (uc *Usecase) Delete(c context.Context, id int) error {
	uc.logger.InfoLogger.Println("usecase Delete")
	return uc.repository.Delete(c, id)
}
func (uc *Usecase) Update(c context.Context, n models.Project) error {
	uc.logger.InfoLogger.Println("usecase Update")
	return uc.repository.Update(c, n)
}
func (uc *Usecase) GetAllforClient(c context.Context) ([]models.Project, error) {
	uc.logger.InfoLogger.Println("usecase Client Get")
	return uc.repository.GetAllforClient(c)
}
func (uc *Usecase) GetAllforAdmin(c context.Context) ([]models.Project, error) {
	uc.logger.InfoLogger.Println("usecase Admin Get")
	return uc.repository.GetAllForAdmin(c)
}
