package news

import (
	"cgr/models"
	"context"
)

type UseCase interface {
	Post(context.Context, models.News) error
	Get(context.Context, string) (models.News, error)
	Delete(context.Context, string) error
	Update(context.Context, models.News) error
	GetAllForClient(context.Context) ([]models.News, error)
	GetAllForAdmin(context.Context) ([]models.News, error)
}
