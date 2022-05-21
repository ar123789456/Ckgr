package link

import (
	"cgr/models"
	"context"
)

type UseCase interface {
	Create(context.Context, models.Link) error
	Delete(context.Context, int)
	GetAll(context.Context) ([]models.Link, error)
}
