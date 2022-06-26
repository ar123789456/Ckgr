package usecase

import (
	"cgr/models"
	"cgr/news"
	"cgr/tool/logger"
	"context"
)

type Usecase struct {
	repository news.Repository
	logger     *logger.Logger
}

func NewUsecase(repository news.Repository, logger *logger.Logger) *Usecase {
	return &Usecase{
		repository: repository,
		logger:     logger,
	}
}

func (uc *Usecase) Post(c context.Context, n models.News) error {
	uc.logger.InfoLogger.Println("usecase Post")
	return uc.repository.Post(c, n)
}
func (uc *Usecase) Get(c context.Context, id int) (models.News, error) {
	uc.logger.InfoLogger.Println("usecase Get")
	return uc.repository.Get(c, id)
}
func (uc *Usecase) Delete(c context.Context, id int) error {
	uc.logger.InfoLogger.Println("usecase Delete")
	return uc.repository.Delete(c, id)
}
func (uc *Usecase) Update(c context.Context, n models.News) error {
	uc.logger.InfoLogger.Println("usecase Update")
	return uc.repository.Update(c, n)
}
func (uc *Usecase) GetAllForClient(c context.Context) ([]models.News, error) {
	uc.logger.InfoLogger.Println("usecase Client Get")
	return uc.repository.GetAllForClient(c)
}
func (uc *Usecase) GetAllForAdmin(c context.Context) ([]models.News, error) {
	uc.logger.InfoLogger.Println("usecase Admin Get")
	return uc.repository.GetAllForAdmin(c)
}
