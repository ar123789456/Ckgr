package link

import (
	"cgr/models"
	"context"
)

type Repository interface {
	Create(context.Context, models.Link) error
	Delete(context.Context, int) error
	GetAll(context.Context) ([]models.Link, error)
}
