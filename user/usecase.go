package user

import (
	"cgr/models"
	"context"
)

type UseCase interface {
	LogIn(context.Context, models.User) (string, error)
	Create(context.Context, models.User) error
	Delete(context.Context, int) error
	Update(context.Context, models.User) error
	GetByToken(context.Context, string) (bool, error)
}
