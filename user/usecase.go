package user

import (
	"cgr/models"
	"context"
)

type UseCase interface {
	Create(context.Context, *models.User) error
	Delete(context.Context, int) error
	Update(context.Context, *models.User) error
}
