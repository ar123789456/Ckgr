package user

import (
	"cgr/models"
	"context"
)

type Repository interface {
	Create(context.Context, models.User) error
	Delete(context.Context, int) error
	Update(context.Context, models.User) error
	UpdateSesion(context.Context, int, string) error
	GetAll(context.Context) ([]models.User, error)
	Get(context.Context, string) (models.User, error)
	GetByToken(context.Context, string) (bool, error)
}
