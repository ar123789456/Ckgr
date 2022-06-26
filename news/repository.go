package news

import (
	"cgr/models"
	"context"
)

type Repository interface {
	Post(context.Context, models.News) error
	Get(context.Context, int) (models.News, error)
	Delete(context.Context, int) error
	Update(context.Context, models.News) error
	GetAllForClient(context.Context) ([]models.News, error)
	GetAllForAdmin(context.Context) ([]models.News, error)
}
